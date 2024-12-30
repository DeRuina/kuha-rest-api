package main

import (
	"log"
	"net/http"
	"time"
)

type api struct {
	config config
}

type config struct {
	addr string
}

func (app *api) run() error {

	mux := http.NewServeMux()

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("Server started listening on %s", app.config.addr)
	return srv.ListenAndServe()
}
