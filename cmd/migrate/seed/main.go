package main

import (
	"log"

	"github.com/DeRuina/KUHA-REST-API/internal/db"
	"github.com/DeRuina/KUHA-REST-API/internal/env"
	"github.com/DeRuina/KUHA-REST-API/internal/seed"
	"github.com/DeRuina/KUHA-REST-API/internal/store/auth"
)

func main() {
	addr := env.GetString("AUTH_DB_ADDR", "")
	conn, err := db.NewSingleDB(addr, 3, 3, "15m")
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer conn.Close()

	authStore := auth.NewAuthStorage(conn)

	seed.Seed(authStore, conn)
}
