package main

import (
	"log"

	"github.com/DeRuina/KUHA-REST-API/internal/db"
	"github.com/DeRuina/KUHA-REST-API/internal/env"
	"github.com/DeRuina/KUHA-REST-API/internal/store"
)

const version = "0.0.1"

//	@title			KUHA REST API
//	@description	API for integrating, analyzing, and visualizing sports and health data
//	@termsOfService	https://csc.fi/en/security-privacy-data-policy-and-open-source-policy/privacy/

//	@BasePath	/v1

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description
func main() {

	cfg := config{
		addr:   env.GetString("ADDR", ":8080"),
		apiURL: env.GetString("EXTERNAL_URL", "localhost:8080"),
		db: dbConfig{
			fisAddr:      env.GetString("FIS_DB_ADDR", ""),
			utvAddr:      env.GetString("UTV_DB_ADDR", ""),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
		env: env.GetString("ENV", "development"),
	}

	databases, err := db.New(cfg.db.fisAddr, cfg.db.utvAddr, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Panic(err)
	}

	defer databases.FIS.Close()
	defer databases.UTV.Close()
	log.Println("database connection pool established")

	store := store.NewStorage(databases)

	app := &api{
		config: cfg,
		store:  *store,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
