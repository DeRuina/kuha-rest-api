package main

import (
	"net/http"

	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

// healthcheckHandler godoc
//
//	@Summary		Healthcheck
//	@Description	Healthcheck endpoint
//	@Tags			Ops
//	@Produce		json
//	@Success		200	{object}	string	"ok"
//	@Failure		500	{object}	error
//	@Router			/health [get]
func (app *api) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"status": "ok",
		"env":    app.config.env,
	}

	if err := utils.WriteJSON(w, http.StatusOK, data); err != nil {
		utils.InternalServerError(w, r, err)
	}
}
