package main

import (
	"context"
	"net/http"
	"runtime"
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

	statusCode := http.StatusOK
	status := "ok"

	// --- Redis Check ---
	if app.cacheStorage != nil {
		if err := app.cacheStorage.Ping(ctx); err != nil {
			data["redis"] = "unreachable"
			status = "fail"
			statusCode = http.StatusInternalServerError
		} else {
			data["redis"] = "ok"
		}
	} else {
		data["redis"] = "disabled"
	}

	// --- DB Checks ---
	if err := app.store.FIS.Ping(ctx); err != nil {
		data["db_fis"] = "unreachable"
		status = "fail"
		statusCode = http.StatusInternalServerError
	} else {
		data["db_fis"] = "ok"
	}

	if err := app.store.UTV.Ping(ctx); err != nil {
		data["db_utv"] = "unreachable"
		status = "fail"
		statusCode = http.StatusInternalServerError
	} else {
		data["db_utv"] = "ok"
	}

	if err := app.store.Auth.Ping(ctx); err != nil {
		data["db_auth"] = "unreachable"
		status = "fail"
		statusCode = http.StatusInternalServerError
	} else {
		data["db_auth"] = "ok"
	}

	// Optional: Uptime and goroutines
	data["uptime_seconds"] = int64(time.Since(startTime).Seconds())
	data["goroutines"] = runtime.NumGoroutine()

	data["api"] = status

	if err := utils.WriteJSON(w, statusCode, data); err != nil {
		utils.InternalServerError(w, r, err)
	}
}
