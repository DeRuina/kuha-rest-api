package main

import (
	"github.com/DeRuina/KUHA-REST-API/internal/auth/authn"
	"github.com/DeRuina/KUHA-REST-API/internal/db"
	"github.com/DeRuina/KUHA-REST-API/internal/env"
	"github.com/DeRuina/KUHA-REST-API/internal/logger"
	"github.com/DeRuina/KUHA-REST-API/internal/store"
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
	}

	// Logger
	logDir := env.GetString("LOG_DIR", "./logs")
	logger.Init(logDir)
	defer logger.Cleanup()

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
		config: cfg,
		store:  *store,
	}

	mux := app.mount()

	logger.Logger.Fatal(app.run(mux))
}
