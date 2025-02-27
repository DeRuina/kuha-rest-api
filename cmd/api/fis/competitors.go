package fisapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DeRuina/KUHA-REST-API/internal/store/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type SectorParams struct {
	Sector string `form:"sector" validate:"required,oneof=JP NK CC"`
}

type CompetitorsHandler struct {
	store fis.Competitors
}

func NewCompetitorsHandler(store fis.Competitors) *CompetitorsHandler {
	return &CompetitorsHandler{store: store}
}

// GetAthletesBySector godoc
//
//	@Summary		Get athletes by sector
//	@Description	Returns a list of athletes for a given sector
//	@Tags			FIS
//	@Accept			json
//	@Produce		json
//	@Param			sector	query		string						true	"Sector Code (JP, NK, CC)"
//	@Success		200		{object}	swagger.AthleteListResponse	"List of athletes"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Router			/fis/athlete [get]
func (h *CompetitorsHandler) GetAthletesBySector(w http.ResponseWriter, r *http.Request) {
	err := utils.ValidateParams(r, []string{"sector"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := SectorParams{
		Sector: r.URL.Query().Get("sector"),
	}

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	competitors, err := h.store.GetAthletesBySector(context.Background(), params.Sector)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	if len(competitors) == 0 {
		utils.NotFoundResponse(w, r, fmt.Errorf("no competitors found for sector %s", params.Sector))
		return
	}

	response := map[string]interface{}{
		"athletes": competitors,
	}

	utils.WriteJSON(w, http.StatusOK, response)
}

// GetNationsBySector godoc
//
//	@Summary		Get nations by sector
//	@Description	Returns a list of nations for a given sector
//	@Tags			FIS
//	@Accept			json
//	@Produce		json
//	@Param			sector	query		string					true	"Sector Code (JP, NK, CC)"
//	@Success		200		{object}	swagger.NationsResponse	"List of nations"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Router			/fis/nation [get]
func (h *CompetitorsHandler) GetNationsBySector(w http.ResponseWriter, r *http.Request) {
	err := utils.ValidateParams(r, []string{"sector"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := SectorParams{
		Sector: r.URL.Query().Get("sector"),
	}

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	nations, err := h.store.GetNationsBySector(context.Background(), params.Sector)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	if len(nations) == 0 {
		utils.NotFoundResponse(w, r, fmt.Errorf("no nations found for sector %s", params.Sector))
		return
	}

	response := map[string]interface{}{
		"nations": nations,
	}

	utils.WriteJSON(w, http.StatusOK, response)
}
