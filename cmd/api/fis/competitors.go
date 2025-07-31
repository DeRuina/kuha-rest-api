package fisapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authn"
	"github.com/DeRuina/KUHA-REST-API/internal/auth/authz"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/DeRuina/KUHA-REST-API/internal/store/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type SectorParams struct {
	Sector string `form:"sector" validate:"required,oneof=JP NK CC"`
}

type CompetitorsHandler struct {
	store fis.Competitors
	cache *cache.Storage
}

func NewCompetitorsHandler(store fis.Competitors, cache *cache.Storage) *CompetitorsHandler {
	return &CompetitorsHandler{store: store, cache: cache}
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
//	@Failure		401 	{object}	swagger.UnauthorizedResponse
//	@Failure		403 	{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503 	{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/athlete [get]
func (h *CompetitorsHandler) GetAthletesBySector(w http.ResponseWriter, r *http.Request) {
	fmt.Println("roles:", authn.GetClientRoles(r.Context()))
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

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

	cacheKey := fmt.Sprintf("fis:athletes:%s", params.Sector)

	if h.cache != nil {
		if cached, err := h.cache.Get(r.Context(), cacheKey); err == nil && cached != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(cached))
			return
		}
	}

	competitors, err := h.store.GetAthletesBySector(r.Context(), params.Sector)
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

	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, response, 3*time.Minute)

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
//	@Failure		401 	{object}	swagger.UnauthorizedResponse
//	@Failure		403 	{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503 	{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/nation [get]
func (h *CompetitorsHandler) GetNationsBySector(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

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

	cacheKey := fmt.Sprintf("fis:nations:%s", params.Sector)

	if h.cache != nil {
		if cached, err := h.cache.Get(r.Context(), cacheKey); err == nil && cached != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(cached))
			return
		}
	}

	nations, err := h.store.GetNationsBySector(r.Context(), params.Sector)
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

	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, response, 3*time.Minute)

	utils.WriteJSON(w, http.StatusOK, response)
}
