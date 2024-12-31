package main

import (
	"log"

	"github.com/DeRuina/KUHA-REST-API/internal/env"
)

func main() {

	cfg := config{
		addr: env.GetString("URL", ":8080"),
	}

	app := &api{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
