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

const version = "1.0.0"

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
			fisAddr:      env.GetString("FIS_DB_ADDR", ""),
			utvAddr:      env.GetString("UTV_DB_ADDR", ""),
			authAddr:     env.GetString("AUTH_DB_ADDR", ""),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
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
		if err := rdb.Ping(context.Background()).Err(); err != nil {
			logger.Logger.Warnw("failed to connect to Redis", "error", err)
		} else {
			cacheStorage = cache.NewRedisStorage(rdb)
			redisLimiter = ratelimiter.NewRedisSlidingLimiter(rdb)
			logger.Logger.Info("Redis cache connection established")
			defer rdb.Close()
		}
		defer rdb.Close()
	} else {
		logger.Logger.Info("Redis cache disabled by configuration")
		localLimiter = ratelimiter.NewFixedWindowLimiter(
			cfg.rateLimiter.RequestsPerTimeFrame,
			cfg.rateLimiter.TimeFrame,
		)
	}

	// Database
	databases, err := db.New(cfg.db.fisAddr, cfg.db.utvAddr, cfg.db.authAddr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		logger.Logger.Fatal(err)
	}

	defer databases.FIS.Close()
	defer databases.UTV.Close()
	defer databases.Auth.Close()
	logger.Logger.Info("database connection pool established")

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
	expvar.Publish("goroutines", expvar.Func(func() any {
		return runtime.NumGoroutine()
	}))

	mux := app.mount()

	logger.Logger.Fatal(app.run(mux))
}
