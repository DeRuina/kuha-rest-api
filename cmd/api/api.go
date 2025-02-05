package main

import (
	"log"
	"net/http"
	"time"

	"github.com/DeRuina/KUHA-REST-API/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	fisapi "github.com/DeRuina/KUHA-REST-API/cmd/api/fis"
	utvapi "github.com/DeRuina/KUHA-REST-API/cmd/api/utv"
)

type api struct {
	config config
	store  store.Storage
}

type config struct {
	addr string
	db   dbConfig
	env  string
}

type dbConfig struct {
	fisAddr      string
	utvAddr      string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *api) mount() http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack (taken from chi )
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)

		// FIS routes
		r.Route("/fis", func(r chi.Router) {
			//register handlers
			competitorsHandler := fisapi.NewCompetitorsHandler(app.store.FIS.Competitors())

			r.Get("/athlete", competitorsHandler.GetAthletesBySector)
			r.Get("/nation", competitorsHandler.GetNationsBySector)
		})

		// UTV routes
		r.Route("/utv", func(r chi.Router) {
			//register handlers
			ouraHandler := utvapi.NewOuraDataHandler(app.store.UTV.Oura())

			// Oura routes
			r.Route("/oura", func(r chi.Router) {
				r.Get("/dates", ouraHandler.GetDates)
				r.Get("/types", ouraHandler.GetTypes)
				r.Get("/data", ouraHandler.GetDataPoint)
				r.Get("/unique-types", ouraHandler.GetUniqueTypes)
			})
		})
	})

	return r
}

func (app *api) run(mux http.Handler) error {

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
