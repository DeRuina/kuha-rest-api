package fisapi

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authz"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/DeRuina/KUHA-REST-API/internal/store/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type CompetitorHandler struct {
	store fis.Competitors
	cache *cache.Storage
}

func NewCompetitorHandler(store fis.Competitors, cache *cache.Storage) *CompetitorHandler {
	return &CompetitorHandler{store: store, cache: cache}
}

// GetAthletesBySector godoc
//
//	@Summary	Get all athletes for a given sector
//	@Tags		FIS - Athlete
//	@Accept		json
//	@Produce	json
//	@Param		sectorcode	query		string	true	"Sector code (JP, NK, CC)"
//	@Success	200			{object}	swagger.FISAthletesResponse
//	@Failure	400			{object}	swagger.ValidationErrorResponse
//	@Failure	401			{object}	swagger.UnauthorizedResponse
//	@Failure	403			{object}	swagger.ForbiddenResponse
//	@Failure	500			{object}	swagger.InternalServerErrorResponse
//	@Failure	503			{object}	swagger.ServiceUnavailableResponse
//	@Security	BearerAuth
//	@Router		/fis/athlete [get]
func (h *CompetitorHandler) GetAthletesBySector(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}
	if err := utils.ValidateParams(r, []string{"sectorcode"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := SectorParam{
		Sector: r.URL.Query().Get("sectorcode"),
	}

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	if h.cache != nil {
		key := fmt.Sprintf("%s:%s", fisAthletesPrefix, params.Sector)
		if raw, err := h.cache.Get(r.Context(), key); err == nil && raw != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(raw))
			return
		}
	}

	rows, err := h.store.GetAthletesBySector(r.Context(), params.Sector)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	body := map[string]any{"athletes": rows}
	cache.SetCacheJSON(r.Context(), h.cache, fmt.Sprintf("%s:%s", fisAthletesPrefix, params.Sector), body, FISCacheTTL)
	utils.WriteJSON(w, http.StatusOK, body)
}

// GetNationsBySector godoc
//
//	@Summary	Get distinct nation codes for a given sector
//	@Tags		FIS - Athlete
//	@Accept		json
//	@Produce	json
//	@Param		sectorcode	query		string	true	"Sector code (JP, NK, CC)"
//	@Success	200			{object}	swagger.FISNationsBySectorResponse
//	@Failure	400			{object}	swagger.ValidationErrorResponse
//	@Failure	401			{object}	swagger.UnauthorizedResponse
//	@Failure	403			{object}	swagger.ForbiddenResponse
//	@Failure	500			{object}	swagger.InternalServerErrorResponse
//	@Failure	503			{object}	swagger.ServiceUnavailableResponse
//	@Security	BearerAuth
//	@Router		/fis/nation [get]
func (h *CompetitorHandler) GetNationsBySector(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}
	if err := utils.ValidateParams(r, []string{"sectorcode"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := SectorParam{
		Sector: r.URL.Query().Get("sectorcode"),
	}
	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	if h.cache != nil {
		key := fmt.Sprintf("%s:%s", fisNationsPrefix, params.Sector)
		if raw, err := h.cache.Get(r.Context(), key); err == nil && raw != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(raw))
			return
		}
	}

	nations, err := h.store.GetNationsBySector(r.Context(), params.Sector)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	body := NationsBySectorResponse{
		Nations: nations,
	}

	cache.SetCacheJSON(r.Context(), h.cache, fmt.Sprintf("%s:%s", fisNationsPrefix, params.Sector), body, FISCacheTTL)
	utils.WriteJSON(w, http.StatusOK, body)
}

// GetLastRowCompetitor godoc
//
//	@Summary		Get last competitor record
//	@Description	Returns the last row in the competitor table
//	@Tags			FIS - Competitor Management
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	swagger.FISLastCompetitorResponse
//	@Failure		401	{object}	swagger.UnauthorizedResponse
//	@Failure		403	{object}	swagger.ForbiddenResponse
//	@Failure		404	{object}	swagger.NotFoundResponse
//	@Failure		500	{object}	swagger.InternalServerErrorResponse
//	@Failure		503	{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/lastrow/competitor [get]
func (h *CompetitorHandler) GetLastRowCompetitor(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if h.cache != nil {
		if raw, err := h.cache.Get(r.Context(), fisLastRowPrefix); err == nil && raw != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(raw))
			return
		}
	}

	row, err := h.store.GetLastRowCompetitor(r.Context())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.NotFoundResponse(w, r, err)
			return
		}
		utils.InternalServerError(w, r, err)
		return
	}

	resp := FISCompetitorFullFromSqlc(row)
	body := map[string]any{"competitor": resp}

	cache.SetCacheJSON(r.Context(), h.cache, fisLastRowPrefix, body, FISCacheTTL)
	utils.WriteJSON(w, http.StatusOK, body)
}

// InsertCompetitor godoc
//
//	@Summary		Add new competitor
//	@Description	Inserts a new competitor
//	@Tags			FIS - Competitor Management
//	@Accept			json
//	@Produce		json
//	@Param			competitor	body	swagger.FISInsertCompetitorExample	true	"Competitor payload"
//	@Success		201			"Created"
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		401			{object}	swagger.UnauthorizedResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		409			{object}	swagger.ConflictResponse
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
//	@Failure		503			{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/competitor [post]
func (h *CompetitorHandler) InsertCompetitor(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var in InsertCompetitorInput
	if err := utils.ReadJSON(w, r, &in); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	if err := utils.GetValidator().Struct(in); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	clean, err := mapInsertInput(in)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	if err := h.store.InsertCompetitor(r.Context(), clean); err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	invalidateCompetitor(r.Context(), h.cache, clean.Competitorid)

	if clean.Sectorcode != nil && *clean.Sectorcode != "" {
		invalidateSector(r.Context(), h.cache, *clean.Sectorcode)
	}

	w.WriteHeader(http.StatusCreated)
}

// UpdateCompetitor godoc
//
//	@Summary		Update competitor by ID
//	@Description	Updates an existing competitor
//	@Tags			FIS - Competitor Management
//	@Accept			json
//	@Produce		json
//	@Param			competitor	body	swagger.FISUpdateCompetitorExample	true	"Competitor payload"
//	@Success		200			"Updated"
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		401			{object}	swagger.UnauthorizedResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		404			{object}	swagger.NotFoundResponse
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
//	@Failure		503			{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/competitor [put]
func (h *CompetitorHandler) UpdateCompetitor(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var in UpdateCompetitorInput
	if err := utils.ReadJSON(w, r, &in); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	if err := utils.GetValidator().Struct(in); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	clean, err := mapUpdateInput(in)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	if err := h.store.UpdateCompetitorByID(r.Context(), clean); err != nil {
		if err == sql.ErrNoRows {
			utils.NotFoundResponse(w, r, err)
			return
		}
		utils.HandleDatabaseError(w, r, err)
		return
	}

	invalidateCompetitor(r.Context(), h.cache, clean.Competitorid)
	if clean.Sectorcode != nil && *clean.Sectorcode != "" {
		invalidateSector(r.Context(), h.cache, *clean.Sectorcode)
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteCompetitor godoc
//
//	@Summary		Delete competitor
//	@Description	Deletes a competitor by ID
//	@Tags			FIS - Competitor Management
//	@Accept			json
//	@Produce		json
//	@Param			id	query	integer	true	"Competitor ID"
//	@Success		200	"Deleted"
//	@Failure		400	{object}	swagger.ValidationErrorResponse
//	@Failure		401	{object}	swagger.UnauthorizedResponse
//	@Failure		403	{object}	swagger.ForbiddenResponse
//	@Failure		404	{object}	swagger.NotFoundResponse
//	@Failure		500	{object}	swagger.InternalServerErrorResponse
//	@Failure		503	{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/competitor [delete]
func (h *CompetitorHandler) DeleteCompetitor(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}
	if err := utils.ValidateParams(r, []string{"id"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := CompetitorIDParam{CompetitorID: r.URL.Query().Get("id")}
	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	id, err := utils.ParsePositiveInt32(params.CompetitorID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	if err := h.store.DeleteCompetitorByID(r.Context(), id); err != nil {
		if err == sql.ErrNoRows {
			utils.NotFoundResponse(w, r, err)
			return
		}
		utils.HandleDatabaseError(w, r, err)
		return
	}

	invalidateCompetitor(r.Context(), h.cache, id)
	invalidateSector(r.Context(), h.cache, "JP")
	invalidateSector(r.Context(), h.cache, "NK")
	invalidateSector(r.Context(), h.cache, "CC")

	w.WriteHeader(http.StatusOK)
}

// SearchCompetitors godoc
//
//	@Summary		Search competitors
//	@Description	Gets competitors filtered by optional Nationcode, Sectorcode, Gender and age range (in years).
//	@Description	Age filters (agemin/agemax) are converted internally to a birthdate range based on today's date.
//	@Tags			FIS - KAMK
//	@Accept			json
//	@Produce		json
//	@Param			nationcode	query		string	false	"Nation code filter (e.g. FIN)"
//	@Param			sectorcode	query		string	false	"Sector code filter (CC,JP,NK)"
//	@Param			gender		query		string	false	"Gender filter (M/W)"
//	@Param			agemin		query		int		false	"Minimum age in years (inclusive). For example, agemin=18 means competitors who are at least 18."
//	@Param			agemax		query		int		false	"Maximum age in years (inclusive). For example, agemax=30 means competitors who are at most 30."
//	@Success		200			{object}	swagger.FISCompetitorSearchResponse
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		401			{object}	swagger.UnauthorizedResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
//	@Failure		503			{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/competitor/search [get]
func (h *CompetitorHandler) SearchCompetitors(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{
		"nationcode",
		"sectorcode",
		"gender",
		"agemin",
		"agemax",
	}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	q := r.URL.Query()

	var nationPtr *string
	if nv := strings.TrimSpace(q.Get("nationcode")); nv != "" {
		nv = strings.ToUpper(nv)
		nationPtr = &nv
	}

	var sectorPtr *string
	if sv := strings.TrimSpace(q.Get("sectorcode")); sv != "" {
		sv = strings.ToUpper(sv)
		sectorPtr = &sv
	}

	var genderPtr *string
	if gv := strings.TrimSpace(q.Get("gender")); gv != "" {
		gv = strings.ToUpper(gv)
		genderPtr = &gv
	}

	var birthMinPtr, birthMaxPtr *time.Time
	var ageminPtr, agemaxPtr *int32

	if am := strings.TrimSpace(q.Get("agemin")); am != "" {
		v, err := utils.ParsePositiveInt32(am)
		if err != nil {
			utils.BadRequestResponse(w, r, fmt.Errorf("invalid agemin: %s", am))
			return
		}
		ageminPtr = &v
	}
	if ax := strings.TrimSpace(q.Get("agemax")); ax != "" {
		v, err := utils.ParsePositiveInt32(ax)
		if err != nil {
			utils.BadRequestResponse(w, r, fmt.Errorf("invalid agemax: %s", ax))
			return
		}
		agemaxPtr = &v
	}

	if ageminPtr != nil && agemaxPtr != nil && *ageminPtr > *agemaxPtr {
		utils.BadRequestResponse(w, r, fmt.Errorf("agemin cannot be greater than agemax"))
		return
	}

	if ageminPtr != nil || agemaxPtr != nil {
		today := time.Now().UTC()

		if agemaxPtr != nil {
			d := today.AddDate(-int(*agemaxPtr), 0, 0)
			birthMinPtr = &d
		}

		if ageminPtr != nil {
			d := today.AddDate(-int(*ageminPtr), 0, 0)
			birthMaxPtr = &d
		}
	}

	rows, err := h.store.SearchCompetitors(
		r.Context(),
		nationPtr,
		sectorPtr,
		genderPtr,
		birthMinPtr,
		birthMaxPtr,
	)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	competitors := make([]any, 0, len(rows))
	for _, c := range rows {
		competitors = append(competitors, FISCompetitorFullFromSqlc(c))
	}

	body := map[string]any{
		"competitors": competitors,
	}

	utils.WriteJSON(w, http.StatusOK, body)
}
