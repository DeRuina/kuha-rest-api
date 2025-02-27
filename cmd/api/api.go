package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/DeRuina/KUHA-REST-API/docs" // This is required to generate swagger docs
	"github.com/DeRuina/KUHA-REST-API/internal/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	fisapi "github.com/DeRuina/KUHA-REST-API/cmd/api/fis"
	utvapi "github.com/DeRuina/KUHA-REST-API/cmd/api/utv"
)

type api struct {
	config config
	store  store.Storage
}

type config struct {
	addr   string
	db     dbConfig
	env    string
	apiURL string
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
		// Healthcheck
		r.Get("/health", app.healthCheckHandler)

		// Swagger docs
		docsURL := fmt.Sprintf("%s/swagger/doc.json", app.config.addr)
		r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/v1/docs/", http.StatusMovedPermanently)
		})
		r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(docsURL)))

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
			polarHandler := utvapi.NewPolarDataHandler(app.store.UTV.Polar())
			suuntoHandler := utvapi.NewSuuntoDataHandler(app.store.UTV.Suunto())

			// Oura routes
			r.Route("/oura", func(r chi.Router) {
				r.Get("/dates", ouraHandler.GetDates)
				r.Get("/types", ouraHandler.GetTypes)
				r.Get("/data", ouraHandler.GetData)
			})

			// Polar routes
			r.Route("/polar", func(r chi.Router) {
				r.Get("/dates", polarHandler.GetDates)
				r.Get("/types", polarHandler.GetTypes)
				r.Get("/data", polarHandler.GetData)
			})

			// Suunto routes
			r.Route("/suunto", func(r chi.Router) {
				r.Get("/dates", suuntoHandler.GetDates)
				r.Get("/types", suuntoHandler.GetTypes)
				r.Get("/data", suuntoHandler.GetData)
			})

		})
	})

	return r
}

func (app *api) run(mux http.Handler) error {
	// Docs
	docs.SwaggerInfo.Version = version
	docs.SwaggerInfo.Host = app.config.apiURL
	docs.SwaggerInfo.BasePath = "/v1"

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
