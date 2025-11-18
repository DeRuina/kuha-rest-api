package fisapi

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authz"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/DeRuina/KUHA-REST-API/internal/store/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type RaceCCHandler struct {
	store fis.Racecc
	cache *cache.Storage
}

func NewRaceCCHandler(store fis.Racecc, cache *cache.Storage) *RaceCCHandler {
	return &RaceCCHandler{store: store, cache: cache}
}

// GetSeasonCodesCC godoc
//
//	@Summary	Get Cross-Country season codes
//	@Tags		FIS - Season Discipline & Category Codes
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	swagger.FISSeasonsCCResponse
//	@Failure	400	{object}	swagger.ValidationErrorResponse
//	@Failure	401	{object}	swagger.UnauthorizedResponse
//	@Failure	403	{object}	swagger.ForbiddenResponse
//	@Failure	500	{object}	swagger.InternalServerErrorResponse
//	@Failure	503	{object}	swagger.ServiceUnavailableResponse
//	@Security	BearerAuth
//	@Router		/fis/seasoncodeCC [get]
func (h *RaceCCHandler) GetSeasonCodesCC(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	cacheKey := fmt.Sprintf("%s:seasons", fisRaceCCCodesPrefix)
	if h.cache != nil {
		if raw, err := h.cache.Get(r.Context(), cacheKey); err == nil && raw != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(raw))
			return
		}
	}

	rows, err := h.store.GetCrossCountrySeasons(r.Context())
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	body := map[string]any{"seasons": rows}
	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, body, FISCacheTTL)
	utils.WriteJSON(w, http.StatusOK, body)
}

// GetDisciplineCodesCC godoc
//
//	@Summary	Get Cross-Country discipline codes
//	@Tags		FIS - Season Discipline & Category Codes
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	swagger.FISDisciplinesCCResponse
//	@Failure	400	{object}	swagger.ValidationErrorResponse
//	@Failure	401	{object}	swagger.UnauthorizedResponse
//	@Failure	403	{object}	swagger.ForbiddenResponse
//	@Failure	500	{object}	swagger.InternalServerErrorResponse
//	@Failure	503	{object}	swagger.ServiceUnavailableResponse
//	@Security	BearerAuth
//	@Router		/fis/disciplinecodeCC [get]
func (h *RaceCCHandler) GetDisciplineCodesCC(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	cacheKey := fmt.Sprintf("%s:disciplines", fisRaceCCCodesPrefix)
	if h.cache != nil {
		if raw, err := h.cache.Get(r.Context(), cacheKey); err == nil && raw != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(raw))
			return
		}
	}

	rows, err := h.store.GetCrossCountryDisciplines(r.Context())
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	body := map[string]any{"disciplines": rows}
	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, body, FISCacheTTL)
	utils.WriteJSON(w, http.StatusOK, body)
}

// GetCategoryCodesCC godoc
//
//	@Summary	Get Cross-Country category codes
//	@Tags		FIS - Season Discipline & Category Codes
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	swagger.FISCategoriesCCResponse
//	@Failure	400	{object}	swagger.ValidationErrorResponse
//	@Failure	401	{object}	swagger.UnauthorizedResponse
//	@Failure	403	{object}	swagger.ForbiddenResponse
//	@Failure	500	{object}	swagger.InternalServerErrorResponse
//	@Failure	503	{object}	swagger.ServiceUnavailableResponse
//	@Security	BearerAuth
//	@Router		/fis/catcodeCC [get]
func (h *RaceCCHandler) GetCategoryCodesCC(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	cacheKey := fmt.Sprintf("%s:categories", fisRaceCCCodesPrefix)
	if h.cache != nil {
		if raw, err := h.cache.Get(r.Context(), cacheKey); err == nil && raw != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(raw))
			return
		}
	}

	rows, err := h.store.GetCrossCountryCategories(r.Context())
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	body := map[string]any{"categories": rows}
	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, body, FISCacheTTL)
	utils.WriteJSON(w, http.StatusOK, body)
}

// GetRacesCC godoc
//
//	@Summary	Get list of Cross-Country races
//	@Tags		FIS - Race Data
//	@Accept		json
//	@Produce	json
//	@Param		seasoncode		query		[]int32		false	"Season code (repeat or comma-separated)"
//	@Param		disciplinecode	query		[]string	false	"Discipline code (repeat or comma-separated)"
//	@Param		catcode			query		[]string	false	"Category code (repeat or comma-separated)"
//	@Success	200				{object}	swagger.FISRacesCCResponse
//	@Failure	400				{object}	swagger.ValidationErrorResponse
//	@Failure	401				{object}	swagger.UnauthorizedResponse
//	@Failure	403				{object}	swagger.ForbiddenResponse
//	@Failure	500				{object}	swagger.InternalServerErrorResponse
//	@Failure	503				{object}	swagger.ServiceUnavailableResponse
//	@Security	BearerAuth
//	@Router		/fis/racecc [get]
func (h *RaceCCHandler) GetRacesCC(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	// accept repeated query params OR comma-separated lists
	parseList := func(key string) []string {
		vals := r.URL.Query()[key]
		if len(vals) == 1 && strings.Contains(vals[0], ",") {
			return strings.Split(vals[0], ",")
		}
		return vals
	}
	seasonsS := parseList("seasoncode")
	discs := parseList("disciplinecode")
	cats := parseList("catcode")

	// trim + convert seasons to []int32
	var seasons []int32
	for _, s := range seasonsS {
		s = strings.TrimSpace(s)
		if s == "" {
			continue
		}
		n, err := strconv.ParseInt(s, 10, 32)
		if err != nil {
			utils.BadRequestResponse(w, r, fmt.Errorf("invalid seasoncode: %s", s))
			return
		}
		seasons = append(seasons, int32(n))
	}

	cacheKey := fmt.Sprintf("%s:sc=%v:dc=%v:cc=%v", fisRaceCCListPrefix, seasons, discs, cats)
	if h.cache != nil {
		if raw, err := h.cache.Get(r.Context(), cacheKey); err == nil && raw != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(raw))
			return
		}
	}

	rows, err := h.store.GetRacesCC(r.Context(), seasons, discs, cats)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	out := make([]FISRaceCCFullResponse, 0, len(rows))
	for _, row := range rows {
		out = append(out, FISRaceCCFullFromSqlc(row))
	}

	body := map[string]any{"races": out}
	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, body, FISCacheTTL)
	utils.WriteJSON(w, http.StatusOK, body)
}

// GetLastRowRaceCC godoc
//
//	@Summary		Get last Cross-Country race record
//	@Description	Returns the last row in a_racecc (by RaceID DESC)
//	@Tags			FIS - Race Management – Cross-Country
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	swagger.FISLastRaceCCResponse
//	@Failure		400	{object}	swagger.ValidationErrorResponse
//	@Failure		401	{object}	swagger.UnauthorizedResponse
//	@Failure		403	{object}	swagger.ForbiddenResponse
//	@Failure		404	{object}	swagger.NotFoundResponse
//	@Failure		500	{object}	swagger.InternalServerErrorResponse
//	@Failure		503	{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/lastrow/racecc [get]
func (h *RaceCCHandler) GetLastRowRaceCC(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if h.cache != nil {
		if raw, err := h.cache.Get(r.Context(), fisRaceCCLastRowPrefix); err == nil && raw != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(raw))
			return
		}
	}

	row, err := h.store.GetLastRowRaceCC(r.Context())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.NotFoundResponse(w, r, err)
			return
		}
		utils.InternalServerError(w, r, err)
		return
	}

	body := map[string]any{"race": FISRaceCCFullFromSqlc(row)}
	cache.SetCacheJSON(r.Context(), h.cache, fisRaceCCLastRowPrefix, body, FISCacheTTL)
	utils.WriteJSON(w, http.StatusOK, body)
}

// InsertRaceCC godoc
//
//	@Summary		Add new Cross-Country race
//	@Description	Inserts a new a_racecc row
//	@Tags			FIS - Race Management – Cross-Country
//	@Accept			json
//	@Produce		json
//	@Param			racecc	body	swagger.FISInsertRaceCCExample	true	"Race payload"
//	@Success		201		"Created"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		409		{object}	swagger.ConflictResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/racecc [post]
func (h *RaceCCHandler) InsertRaceCC(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var in InsertRaceCCInput
	if err := utils.ReadJSON(w, r, &in); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	if err := utils.GetValidator().Struct(in); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	clean, err := mapInsertRaceCCInput(in)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	if err := h.store.InsertRaceCC(r.Context(), clean); err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	invalidateRaceCC(r.Context(), h.cache, clean.Raceid)
	w.WriteHeader(http.StatusCreated)
}

// UpdateRaceCC godoc
//
//	@Summary		Update Cross-Country race by ID
//	@Description	Updates an existing a_racecc row
//	@Tags			FIS - Race Management – Cross-Country
//	@Accept			json
//	@Produce		json
//	@Param			racecc	body	swagger.FISUpdateRaceCCExample	true	"Race payload"
//	@Success		200		"Updated"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		404		{object}	swagger.NotFoundResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/racecc [put]
func (h *RaceCCHandler) UpdateRaceCC(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var in UpdateRaceCCInput
	if err := utils.ReadJSON(w, r, &in); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	if err := utils.GetValidator().Struct(in); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	clean, err := mapUpdateRaceCCInput(in)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	if err := h.store.UpdateRaceCCByID(r.Context(), clean); err != nil {
		if err == sql.ErrNoRows {
			utils.NotFoundResponse(w, r, err)
			return
		}
		utils.HandleDatabaseError(w, r, err)
		return
	}

	invalidateRaceCC(r.Context(), h.cache, clean.Raceid)
	w.WriteHeader(http.StatusOK)
}

// DeleteRaceCC godoc
//
//	@Summary		Delete Cross-Country race
//	@Description	Deletes a race by RaceID
//	@Tags			FIS - Race Management – Cross-Country
//	@Accept			json
//	@Produce		json
//	@Param			id	query	int32	true	"Race ID"
//	@Success		200	"Deleted"
//	@Failure		400	{object}	swagger.ValidationErrorResponse
//	@Failure		401	{object}	swagger.UnauthorizedResponse
//	@Failure		403	{object}	swagger.ForbiddenResponse
//	@Failure		404	{object}	swagger.NotFoundResponse
//	@Failure		500	{object}	swagger.InternalServerErrorResponse
//	@Failure		503	{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/racecc [delete]
func (h *RaceCCHandler) DeleteRaceCC(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}
	if err := utils.ValidateParams(r, []string{"id"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := utils.ParsePositiveInt32(idStr)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	if err := h.store.DeleteRaceCCByID(r.Context(), id); err != nil {
		if err == sql.ErrNoRows {
			utils.NotFoundResponse(w, r, err)
			return
		}
		utils.HandleDatabaseError(w, r, err)
		return
	}

	invalidateRaceCC(r.Context(), h.cache, id)
	w.WriteHeader(http.StatusOK)
}
