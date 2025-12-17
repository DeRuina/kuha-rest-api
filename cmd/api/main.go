package main

import (
	"context"
	"expvar"
	"runtime"
	"time"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authn"
	"github.com/DeRuina/KUHA-REST-API/internal/db"
	"github.com/DeRuina/KUHA-REST-API/internal/env"
	"github.com/DeRuina/KUHA-REST-API/internal/logger"
	"github.com/DeRuina/KUHA-REST-API/internal/ratelimiter"
	"github.com/DeRuina/KUHA-REST-API/internal/store"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
)

const version = "1.3.0"

//	@title			KUHA REST API
//	@description	API for integrating, analyzing, and visualizing sports and health data
//	@termsOfService	https://csc.fi/en/security-privacy-data-policy-and-open-source-policy/privacy/

//	@BasePath	/v1

// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description				Use format: Bearer your_JWT_here
func main() {

	cfg := config{
		addr:   env.GetString("ADDR", ":8080"),
		apiURL: env.GetString("EXTERNAL_URL", "localhost:8080"),
		db: dbConfig{
			fisAddr:        env.GetString("FIS_DB_ADDR", ""),
			utvAddr:        env.GetString("UTV_DB_ADDR", ""),
			authAddr:       env.GetString("AUTH_DB_ADDR", ""),
			tietoevryAddr:  env.GetString("TIETOEVRY_DB_ADDR", ""),
			kamkAddr:       env.GetString("KAMK_DB_ADDR", ""),
			klabAddr:       env.GetString("KLAB_DB_ADDR", ""),
			archinisisAddr: env.GetString("ARCHINISIS_DB_ADDR", ""),
			maxOpenConns:   env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns:   env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:    env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		redisCfg: redisConfig{
			addr:    env.GetString("REDIS_ADDR", "localhost:6379"),
			pw:      env.GetString("REDIS_PW", ""),
			db:      env.GetInt("REDIS_DB", 0),
			enabled: env.GetBool("REDIS_ENABLED", false),
		},
		env: env.GetString("ENV", "development"),
		auth: authConfig{
			basic: basicConfig{
				user: env.GetString("BASIC_AUTH_USER", ""),
				pass: env.GetString("BASIC_AUTH_PASS", ""),
			},
			jwt: jwtConfig{
				secret:   []byte(env.GetString("JWT_SECRET", "")),
				issuer:   env.GetString("JWT_ISSUER", ""),
				audience: env.GetString("JWT_AUDIENCE", ""),
			},
		},
		rateLimiter: ratelimiter.Config{
			RequestsPerTimeFrame: env.GetInt("RATELIMITER_REQUESTS_COUNT", 20),
			TimeFrame:            time.Second * 5,
			Enabled:              env.GetBool("RATE_LIMITER_ENABLED", true),
		},
	}

	// Rate limiter
	var redisLimiter *ratelimiter.RedisSlidingLimiter
	var localLimiter *ratelimiter.FixedWindowRateLimiter

	// Logger
	logDir := env.GetString("LOG_DIR", "./logs")
	logger.Init(logDir)
	defer logger.Cleanup()

	// Cache
	var cacheStorage *cache.Storage
	if cfg.redisCfg.enabled {
		rdb := cache.NewRedisClient(cfg.redisCfg.addr, cfg.redisCfg.pw, cfg.redisCfg.db)
		defer rdb.Close()

		if err := rdb.Ping(context.Background()).Err(); err != nil {
			logger.Logger.Warnw("failed to connect to Redis", "error", err)
		} else {
			cacheStorage = cache.NewRedisStorage(rdb)
			redisLimiter = ratelimiter.NewRedisSlidingLimiter(rdb)
			logger.Logger.Info("Redis cache connection established")
		}
	} else {
		logger.Logger.Info("Redis cache disabled by configuration")
		localLimiter = ratelimiter.NewFixedWindowLimiter(
			cfg.rateLimiter.RequestsPerTimeFrame,
			cfg.rateLimiter.TimeFrame,
		)
	}

	// Database - Connect with graceful failure handling
	databases, dbErrors := db.NewWithGracefulFailure(
		cfg.db.fisAddr,
		cfg.db.utvAddr,
		cfg.db.authAddr,
		cfg.db.tietoevryAddr,
		cfg.db.kamkAddr,
		cfg.db.klabAddr,
		cfg.db.archinisisAddr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)

	// Log connection status for each database
	for name, err := range dbErrors {
		if err == nil {
			logger.Logger.Infow("Database connected successfully", "database", name)
		} else if err.Error() == "not configured" {
			logger.Logger.Infow("database not configured", "database", name)
		} else {
			logger.Logger.Warnw("database connection failed", "database", name, "error", err)
		}
	}

	// Close only successful connections
	defer func() {
		if databases.FIS != nil {
			databases.FIS.Close()
		}
		if databases.UTV != nil {
			databases.UTV.Close()
		}
		if databases.Auth != nil {
			databases.Auth.Close()
		}
		if databases.Tietoevry != nil {
			databases.Tietoevry.Close()
		}
		if databases.KAMK != nil {
			databases.KAMK.Close()
		}
		if databases.KLAB != nil {
			databases.KLAB.Close()
		}
		if databases.ARCHINISIS != nil {
			databases.ARCHINISIS.Close()
		}
	}()

	// Authentication
	authn.LoadJWTConfig(authn.JWTConfig{
		Secret:   cfg.auth.jwt.secret,
		Issuer:   cfg.auth.jwt.issuer,
		Audience: cfg.auth.jwt.audience,
	})

	// Storage
	store := store.NewStorage(databases)

	app := &api{
		config:           cfg,
		store:            *store,
		cacheStorage:     cacheStorage,
		redisRateLimiter: redisLimiter,
		localRateLimiter: localLimiter,
	}

	// metrics
	expvar.NewString("version").Set(version)
	expvar.Publish("database_fis", expvar.Func(func() any {
		if databases.FIS != nil {
			return databases.FIS.Stats()
		}
		return nil
	}))
	expvar.Publish("database_utv", expvar.Func(func() any {
		if databases.UTV != nil {
			return databases.UTV.Stats()
		}
		return nil
	}))
	expvar.Publish("database_auth", expvar.Func(func() any {
		if databases.Auth != nil {
			return databases.Auth.Stats()
		}
		return nil
	}))
	expvar.Publish("database_tietoevry", expvar.Func(func() any {
		if databases.Tietoevry != nil {
			return databases.Tietoevry.Stats()
		}
		return nil
	}))

	expvar.Publish("database_kamk", expvar.Func(func() any {
		if databases.KAMK != nil {
			return databases.KAMK.Stats()
		}
		return nil
	}))

	expvar.Publish("database_klab", expvar.Func(func() any {
		if databases.KLAB != nil {
			return databases.KLAB.Stats()
		}
		return nil
	}))

	expvar.Publish("database_archinisis", expvar.Func(func() any {
		if databases.ARCHINISIS != nil {
			return databases.ARCHINISIS.Stats()
		}
		return nil
	}))

	expvar.Publish("goroutines", expvar.Func(func() any {
		return runtime.NumGoroutine()
	}))

	mux := app.mount()

	logger.Logger.Fatal(app.run(mux))
}
