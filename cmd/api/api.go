package main

import (
	"context"
	"errors"
	"expvar"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/DeRuina/KUHA-REST-API/docs" // This is required to generate swagger docs
	"github.com/DeRuina/KUHA-REST-API/internal/env"
	"github.com/DeRuina/KUHA-REST-API/internal/logger"
	"github.com/DeRuina/KUHA-REST-API/internal/ratelimiter"
	"github.com/DeRuina/KUHA-REST-API/internal/store"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	httpSwagger "github.com/swaggo/http-swagger/v2"

	archapi "github.com/DeRuina/KUHA-REST-API/cmd/api/archinisis"
	authapi "github.com/DeRuina/KUHA-REST-API/cmd/api/auth"
	fisapi "github.com/DeRuina/KUHA-REST-API/cmd/api/fis"
	klabapi "github.com/DeRuina/KUHA-REST-API/cmd/api/klab"
	tietoevryapi "github.com/DeRuina/KUHA-REST-API/cmd/api/tietoevry"
	utvapi "github.com/DeRuina/KUHA-REST-API/cmd/api/utv"
)

type api struct {
	config           config
	store            store.Storage
	cacheStorage     *cache.Storage
	redisRateLimiter *ratelimiter.RedisSlidingLimiter
	localRateLimiter *ratelimiter.FixedWindowRateLimiter
}

type config struct {
	addr        string
	db          dbConfig
	env         string
	apiURL      string
	auth        authConfig
	redisCfg    redisConfig
	rateLimiter ratelimiter.Config
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
	fisAddr        string
	utvAddr        string
	authAddr       string
	tietoevryAddr  string
	kamkAddr       string
	klabAddr       string
	archinisisAddr string
	maxOpenConns   int
	maxIdleConns   int
	maxIdleTime    string
}

func (app *api) mount() http.Handler {
	r := chi.NewRouter()

	// Middlewares
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(middleware.Recoverer)
	r.Use(middleware.RealIP)
	r.Use(ExtractClientIDMiddleware())
	r.Use(app.RateLimiterMiddleware)
	r.Use(middleware.RequestID)
	r.Use(logger.LoggerMiddleware)

	origins := strings.Split(env.GetString("CORS_ALLOWED_ORIGIN", ""), ",")
	for i := range origins {
		origins[i] = strings.TrimSpace(origins[i])
	}

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   origins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Route("/v1", func(r chi.Router) {
		// Auth routes
		if app.store.Auth != nil {
			r.Route("/auth", func(r chi.Router) {
				authHandler := authapi.NewAuthHandler(app.store.Auth)
				r.Post("/token", authHandler.IssueTokens)
				r.Post("/refresh", authHandler.RefreshToken)
			})
		} else {
			logger.Logger.Warn("Auth routes disabled: database not connected")
			r.Route("/auth", func(r chi.Router) {
				r.Handle("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					utils.ServiceUnavailableDBResponse(w, r, "Auth")
				}))
			})
		}

		// Healthcheck
		r.Get("/health", app.healthCheckHandler)

		// Metrics
		r.With(app.BasicAuthMiddleware()).Get("/metrics", expvar.Handler().ServeHTTP)

		// Swagger docs
		docsURL := fmt.Sprintf("%s/swagger/doc.json", app.config.addr)
		r.Get("/docs", func(w http.ResponseWriter, r *http.Request) {
			http.Redirect(w, r, "/v1/docs/", http.StatusMovedPermanently)
		})
		r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL(docsURL)))

		r.Group(func(r chi.Router) {
			r.Use(JWTMiddleware())

			// Tietoevry routes
			if app.store.Tietoevry != nil {
				r.Route("/tietoevry", func(r chi.Router) {
					// Register handlers
					userHandler := tietoevryapi.NewTietoevryUserHandler(app.store.Tietoevry.Users(), app.cacheStorage)
					exerciseHandler := tietoevryapi.NewTietoevryExerciseHandler(app.store.Tietoevry.Exercises(), app.cacheStorage)
					symptomHandler := tietoevryapi.NewTietoevrySymptomHandler(app.store.Tietoevry.Symptoms(), app.cacheStorage)
					measurementHandler := tietoevryapi.NewTietoevryMeasurementHandler(app.store.Tietoevry.Measurements(), app.cacheStorage)
					testResultHandler := tietoevryapi.NewTietoevryTestResultHandler(app.store.Tietoevry.TestResults(), app.cacheStorage)
					questionnaireHandler := tietoevryapi.NewTietoevryQuestionnaireHandler(app.store.Tietoevry.Questionnaires(), app.cacheStorage)
					activityZoneHandler := tietoevryapi.NewTietoevryActivityZoneHandler(app.store.Tietoevry.ActivityZones(), app.cacheStorage)

					// user routes
					r.Post("/users", userHandler.UpsertUser)
					r.Delete("/users", userHandler.DeleteUser)
					r.Get("/users", userHandler.GetUser)
					r.Get("/deleted-users", userHandler.GetDeletedUsers)

					// exercise routes
					r.With(GzipDecompressionMiddleware()).Post("/exercises", exerciseHandler.InsertExercisesBulk)
					r.Get("/exercises", exerciseHandler.GetExercises)

					// symptom routes
					r.With(GzipDecompressionMiddleware()).Post("/symptoms", symptomHandler.InsertSymptomsBulk)
					r.Get("/symptoms", symptomHandler.GetSymptoms)

					// measurement routes
					r.With(GzipDecompressionMiddleware()).Post("/measurements", measurementHandler.InsertMeasurementsBulk)
					r.Get("/measurements", measurementHandler.GetMeasurements)

					// test result routes
					r.With(GzipDecompressionMiddleware()).Post("/test-results", testResultHandler.InsertTestResultsBulk)
					r.Get("/test-results", testResultHandler.GetTestResults)

					// questionnaire routes
					r.With(GzipDecompressionMiddleware()).Post("/questionnaires", questionnaireHandler.InsertQuestionnaireAnswersBulk)
					r.Get("/questionnaires", questionnaireHandler.GetQuestionnaires)

					// activity zone routes
					r.With(GzipDecompressionMiddleware()).Post("/activity-zones", activityZoneHandler.InsertActivityZonesBulk)
					r.Get("/activity-zones", activityZoneHandler.GetActivityZones)
				})
			} else {
				logger.Logger.Warn("tietoevry routes disabled: database not connected")
				r.Route("/tietoevry", func(r chi.Router) {
					r.Handle("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						utils.ServiceUnavailableDBResponse(w, r, "Tietoevry")
					}))
				})
			}

			// Archinisis routes
			if app.store.ARCHINISIS != nil {
				r.Route("/archinisis", func(r chi.Router) {
					// Register handlers
					userDataHandler := archapi.NewUserDataHandler(app.store.ARCHINISIS.Users(), app.cacheStorage)
					dataHandler := archapi.NewDataHandler(app.store.ARCHINISIS.Data(), app.cacheStorage)

					// user routes
					r.Get("/sport-ids", userDataHandler.GetSporttiIDs)

					// data routes
					r.Get("/race-report/sessions", dataHandler.GetRaceReportSessions)
					r.Get("/race-report", dataHandler.GetRaceReportHTML)
				})
			} else {
				logger.Logger.Warn("archinisis routes disabled: database not connected")
				r.Route("/archinisis", func(r chi.Router) {
					r.Handle("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						utils.ServiceUnavailableDBResponse(w, r, "archinisis")
					}))
				})
			}

			// KLAB routes
			if app.store.KLAB != nil {
				r.Route("/klab", func(r chi.Router) {
					// Register handlers
					userDataHandler := klabapi.NewUserDataHandler(app.store.KLAB.Users(), app.cacheStorage)
					klabDataHandler := klabapi.NewKlabDataHandler(app.store.KLAB.Data(), app.cacheStorage)

					// user routes
					r.Get("/sport-ids", userDataHandler.GetSporttiIDs)
					r.Get("/user", userDataHandler.GetUser)

					// data routes
					r.Post("/data", klabDataHandler.InsertKlabDataBulk)
					r.Get("/data", klabDataHandler.GetKlabData)
				})
			} else {
				logger.Logger.Warn("klab routes disabled: database not connected")
				r.Route("/klab", func(r chi.Router) {
					r.Handle("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						utils.ServiceUnavailableDBResponse(w, r, "klab")
					}))
				})
			}

			// FIS routes
			if app.store.FIS != nil {
				r.Route("/fis", func(r chi.Router) {
					// Register handlers
					competitorsHandler := fisapi.NewCompetitorsHandler(app.store.FIS.Competitors(), app.cacheStorage)

					// competitor routes
					r.Get("/athlete", competitorsHandler.GetAthletesBySector)
					r.Get("/nation", competitorsHandler.GetNationsBySector)
				})
			} else {
				logger.Logger.Warn("fis routes disabled: database not reachable")
				r.Route("/fis", func(r chi.Router) {
					r.Handle("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						utils.ServiceUnavailableDBResponse(w, r, "fis")
					}))
				})
			}
			// UTV routes
			if app.store.UTV != nil {
				r.Route("/utv", func(r chi.Router) {
					// Register handlers
					generalHandler := utvapi.NewGeneralDataHandler(
						app.store.UTV.Oura(),
						app.store.UTV.Polar(),
						app.store.UTV.Suunto(),
						app.store.UTV.Garmin(),
						app.store.UTV.OuraToken(),
						app.store.UTV.PolarToken(),
						app.store.UTV.SuuntoToken(),
						app.store.UTV.GarminToken(),
						app.store.UTV.KlabToken(),
						app.cacheStorage,
					)
					ouraHandler := utvapi.NewOuraDataHandler(app.store.UTV.Oura(), app.cacheStorage)
					polarHandler := utvapi.NewPolarDataHandler(app.store.UTV.Polar(), app.cacheStorage)
					suuntoHandler := utvapi.NewSuuntoDataHandler(app.store.UTV.Suunto(), app.cacheStorage)
					garminHandler := utvapi.NewGarminDataHandler(app.store.UTV.Garmin(), app.cacheStorage)
					polarTokenHandler := utvapi.NewPolarTokenHandler(app.store.UTV.PolarToken(), app.cacheStorage)
					ouraTokenHandler := utvapi.NewOuraTokenHandler(app.store.UTV.OuraToken(), app.cacheStorage)
					suuntoTokenHandler := utvapi.NewSuuntoTokenHandler(app.store.UTV.SuuntoToken(), app.cacheStorage)
					garminTokenHandler := utvapi.NewGarminTokenHandler(app.store.UTV.GarminToken(), app.cacheStorage)
					klabTokenHandler := utvapi.NewKlabTokenHandler(app.store.UTV.KlabToken(), app.cacheStorage)
					userDataHandler := utvapi.NewUserDataHandler(app.store.UTV.UserData(), app.cacheStorage)
					coachtechHandler := utvapi.NewCoachtechDataHandler(app.store.UTV.Coachtech(), app.cacheStorage)

					// General routes
					r.Get("/latest", generalHandler.GetLatestData)
					r.Get("/all", generalHandler.GetAllByType)
					r.Delete("/disconnect", generalHandler.Disconnect)
					r.Get("/tokens4update", generalHandler.GetTokensForUpdate)
					r.Get("/data4update", generalHandler.GetDataForUpdate)

					// User data routes
					r.Get("/user", userDataHandler.GetUserData)
					r.Post("/user", userDataHandler.UpsertUserData)
					r.Delete("/user", userDataHandler.DeleteUserData)
					r.Get("/user-id-by-sport-id", userDataHandler.GetUserIDBySportID)
					r.Get("/user-linked-devices", userDataHandler.GetLinkedDevices)

					// Klab routes
					r.Route("/klab", func(r chi.Router) {
						r.Get("/status", klabTokenHandler.GetStatus)
						r.Post("/token", klabTokenHandler.UpsertToken)
					})

					// Coachtech routes
					r.Route("/coachtech", func(r chi.Router) {
						r.Get("/status", coachtechHandler.GetStatus)
						r.Get("/data", coachtechHandler.GetData)
						r.Post("/insert", coachtechHandler.Insert)
					})

					// Oura routes
					r.Route("/oura", func(r chi.Router) {
						r.Get("/dates", ouraHandler.GetDates)
						r.Get("/types", ouraHandler.GetTypes)
						r.Get("/data", ouraHandler.GetData)
						r.Post("/data", ouraHandler.InsertData)
						r.Delete("/data", ouraHandler.DeleteAllData)
						r.Get("/status", ouraTokenHandler.GetStatus)
						r.Post("/token", ouraTokenHandler.UpsertToken)
						r.Get("/token-by-id", ouraTokenHandler.GetTokenByOuraID)
					})

					// Polar routes
					r.Route("/polar", func(r chi.Router) {
						r.Get("/dates", polarHandler.GetDates)
						r.Get("/types", polarHandler.GetTypes)
						r.Get("/data", polarHandler.GetData)
						r.Post("/data", polarHandler.InsertData)
						r.Delete("/data", polarHandler.DeleteAllData)
						r.Get("/status", polarTokenHandler.GetStatus)
						r.Post("/token", polarTokenHandler.UpsertToken)
						r.Get("/token-by-id", polarTokenHandler.GetTokenByPolarID)
					})

					// Suunto routes
					r.Route("/suunto", func(r chi.Router) {
						r.Get("/dates", suuntoHandler.GetDates)
						r.Get("/types", suuntoHandler.GetTypes)
						r.Get("/data", suuntoHandler.GetData)
						r.Post("/data", suuntoHandler.InsertData)
						r.Delete("/data", suuntoHandler.DeleteAllData)
						r.Get("/status", suuntoTokenHandler.GetStatus)
						r.Post("/token", suuntoTokenHandler.UpsertToken)
						r.Get("/token-by-username", suuntoTokenHandler.GetTokenByUsername)
					})

					// Garmin routes
					r.Route("/garmin", func(r chi.Router) {
						r.Get("/dates", garminHandler.GetDates)
						r.Get("/types", garminHandler.GetTypes)
						r.Get("/data", garminHandler.GetData)
						r.Post("/data", garminHandler.InsertData)
						r.Delete("/data", garminHandler.DeleteAllData)
						r.Get("/status", garminTokenHandler.GetStatus)
						r.Post("/token", garminTokenHandler.UpsertToken)
						r.Get("/token-exists", garminTokenHandler.TokenExists)
						r.Get("/user-id-by-token", garminTokenHandler.GetUserIDByToken)
					})
				})
			} else {
				logger.Logger.Warn("utv routes disabled: database not connected")
				r.Route("/utv", func(r chi.Router) {
					r.Handle("/*", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
						utils.ServiceUnavailableDBResponse(w, r, "utv")
					}))
				})
			}
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
		WriteTimeout: time.Second * 60,
		ReadTimeout:  time.Second * 30,
		IdleTimeout:  time.Minute,
	}

	shutdown := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)

		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		logger.Logger.Infow("signal caught", "signal", s.String())

		shutdown <- srv.Shutdown(ctx)
	}()

	logger.Logger.Infow("server has started", "addr", app.config.addr, "env", app.config.env)

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdown
	if err != nil {
		return err
	}

	logger.Logger.Infow("server has stopped", "addr", app.config.addr, "env", app.config.env)

	return nil
}
