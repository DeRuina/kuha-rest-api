package main

import (
	"context"
	"net/http"
	"time"

	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

// Store app start time for uptime calculation
var startTime = time.Now()

// healthCheckHandler godoc
//
//	@Summary		Healthcheck
//	@Description	Healthcheck endpoint
//	@Tags			Health
//	@Produce		json
//	@Success		200	{object}	swagger.HealthStatusResponse	"Health status"
//	@Failure		500	{object}	swagger.InternalServerErrorResponse
//	@Router			/health [get]
func (app *api) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), 1*time.Second)
	defer cancel()

	data := map[string]any{
		"env":     app.config.env,
		"version": version,
	}

	status := "ok"
	statusCode := http.StatusOK

	// Redis check
	if app.config.redisCfg.enabled {
		if app.cacheStorage == nil || app.cacheStorage.Ping(ctx) != nil {
			data["redis"] = "down"
		} else {
			data["redis"] = "ok"
		}
	} else {
		data["redis"] = "down"
	}

	// Helper function for DB checks
	checkDB := func(name string, db interface{ Ping(context.Context) error }) {
		if db == nil {
			data[name] = "down"
			return
		}
		if err := db.Ping(ctx); err != nil {
			data[name] = "down"
			return
		}
		data[name] = "ok"
	}

	// Check each DB safely
	checkDB("db_fis", app.store.FIS)
	checkDB("db_utv", app.store.UTV)
	checkDB("db_auth", app.store.Auth)
	checkDB("db_tietoevry", app.store.Tietoevry)
	checkDB("db_kamk", app.store.KAMK)
	checkDB("db_klab", app.store.KLAB)
	checkDB("db_archinisis", app.store.ARCHINISIS)

	// Add uptime and goroutine info
	data["uptime_seconds"] = int64(time.Since(startTime).Seconds())
	data["api"] = status

	statusCode = http.StatusOK

	if err := utils.WriteJSON(w, statusCode, data); err != nil {
		utils.InternalServerError(w, r, err)
	}
}
