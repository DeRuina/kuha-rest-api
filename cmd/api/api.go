package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/DeRuina/KUHA-REST-API/docs" // This is required to generate swagger docs
	"github.com/DeRuina/KUHA-REST-API/internal/logger"
	"github.com/DeRuina/KUHA-REST-API/internal/store"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	authapi "github.com/DeRuina/KUHA-REST-API/cmd/api/auth"
	fisapi "github.com/DeRuina/KUHA-REST-API/cmd/api/fis"
	utvapi "github.com/DeRuina/KUHA-REST-API/cmd/api/utv"
)

type api struct {
	config       config
	store        store.Storage
	cacheStorage *cache.Storage
}

type config struct {
	addr     string
	db       dbConfig
	env      string
	apiURL   string
	auth     authConfig
	redisCfg redisConfig
}

type redisConfig struct {
	addr    string
	pw      string
	db      int
	enabled bool
}

type authConfig struct {
	basic basicConfig
	jwt   jwtConfig
}

type basicConfig struct {
	user string
	pass string
}

type jwtConfig struct {
	secret   []byte
	issuer   string
	audience string
}

type dbConfig struct {
	fisAddr      string
	utvAddr      string
	authAddr     string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

func (app *api) mount() http.Handler {
	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// Middlewares
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(ExtractClientIDMiddleware())
	r.Use(logger.LoggerMiddleware)

	r.Route("/v1", func(r chi.Router) {
		// Auth routes
		r.Route("/auth", func(r chi.Router) {
			authHandler := authapi.NewAuthHandler(app.store.Auth)

			r.Post("/token", authHandler.IssueTokens)
			r.Post("/refresh", authHandler.RefreshToken)
		})

		// Healthcheck
		r.Get("/health", app.healthCheckHandler)

		// Swagger docs
		docsURL := fmt.Sprintf("%s/swagger/doc.json", app.config.addr)
		r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/v1/docs/", http.StatusMovedPermanently)
		})
		r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(docsURL)))

		r.Group(func(r chi.Router) {
			r.Use(JWTMiddleware())

			// FIS routes
			r.Route("/fis", func(r chi.Router) {
				//register handlers
				competitorsHandler := fisapi.NewCompetitorsHandler(app.store.FIS.Competitors(), app.cacheStorage)

				r.Get("/athlete", competitorsHandler.GetAthletesBySector)
				r.Get("/nation", competitorsHandler.GetNationsBySector)
			})

			// UTV routes
			r.Route("/utv", func(r chi.Router) {
				//register handlers
				ouraHandler := utvapi.NewOuraDataHandler(app.store.UTV.Oura(), app.cacheStorage)
				polarHandler := utvapi.NewPolarDataHandler(app.store.UTV.Polar(), app.cacheStorage)
				suuntoHandler := utvapi.NewSuuntoDataHandler(app.store.UTV.Suunto(), app.cacheStorage)

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

	logger.Logger.Infow("server started", "addr", app.config.addr, "env", app.config.env)
	return srv.ListenAndServe()
}
