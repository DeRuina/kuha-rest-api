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

type ResultJPHandler struct {
	store       fis.Resultjp
	competitors fis.Competitors
	cache       *cache.Storage
}

func NewResultJPHandler(store fis.Resultjp, competitors fis.Competitors, cache *cache.Storage) *ResultJPHandler {
	return &ResultJPHandler{store: store, competitors: competitors, cache: cache}
}

// GetLastRowResultJP godoc
//
//	@Summary		Get last Ski Jumping result record
//	@Description	Returns the last row in a_resultjp (by RecID DESC)
//	@Tags			FIS - Result Management – Ski Jumping
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	swagger.FISLastResultJPResponse
//	@Failure		400	{object}	swagger.ValidationErrorResponse
//	@Failure		401	{object}	swagger.UnauthorizedResponse
//	@Failure		403	{object}	swagger.ForbiddenResponse
//	@Failure		404	{object}	swagger.NotFoundResponse
//	@Failure		500	{object}	swagger.InternalServerErrorResponse
//	@Failure		503	{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/lastrow/resultjp [get]
func (h *ResultJPHandler) GetLastRowResultJP(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if h.cache != nil {
		if raw, err := h.cache.Get(r.Context(), fisResultJPLastRowPrefix); err == nil && raw != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(raw))
			return
		}
	}

	row, err := h.store.GetLastRowResultJP(r.Context())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			utils.NotFoundResponse(w, r, err)
			return
		}
		utils.InternalServerError(w, r, err)
		return
	}

	body := map[string]any{"result": FISResultJPFullFromSqlc(row)}
	cache.SetCacheJSON(r.Context(), h.cache, fisResultJPLastRowPrefix, body, FISCacheTTL)
	utils.WriteJSON(w, http.StatusOK, body)
}

// InsertResultJP godoc
//
//	@Summary		Add new Ski Jumping result
//	@Description	Inserts a new a_resultjp row
//	@Tags			FIS - Result Management – Ski Jumping
//	@Accept			json
//	@Produce		json
//	@Param			resultjp	body	swagger.FISInsertResultJPExample	true	"Result payload"
//	@Success		201			"Created"
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		401			{object}	swagger.UnauthorizedResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		409			{object}	swagger.ConflictResponse
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
//	@Failure		503			{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/resultjp [post]
func (h *ResultJPHandler) InsertResultJP(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var in InsertResultJPInput
	if err := utils.ReadJSON(w, r, &in); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	if err := utils.GetValidator().Struct(in); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	clean, err := mapInsertResultJPInput(in)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	if err := h.store.InsertResultJP(r.Context(), clean); err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	invalidateResultJP(r.Context(), h.cache, clean.Recid)
	w.WriteHeader(http.StatusCreated)
}

// UpdateResultJP godoc
//
//	@Summary		Update Ski Jumping result by RecID
//	@Description	Updates an existing a_resultjp row
//	@Tags			FIS - Result Management – Ski Jumping
//	@Accept			json
//	@Produce		json
//	@Param			resultjp	body	swagger.FISUpdateResultJPExample	true	"Result payload"
//	@Success		200			"Updated"
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		401			{object}	swagger.UnauthorizedResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		404			{object}	swagger.NotFoundResponse
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
//	@Failure		503			{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/resultjp [put]
func (h *ResultJPHandler) UpdateResultJP(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var in UpdateResultJPInput
	if err := utils.ReadJSON(w, r, &in); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	if err := utils.GetValidator().Struct(in); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	clean, err := mapUpdateResultJPInput(in)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	if err := h.store.UpdateResultJPByRecID(r.Context(), clean); err != nil {
		if err == sql.ErrNoRows {
			utils.NotFoundResponse(w, r, err)
			return
		}
		utils.HandleDatabaseError(w, r, err)
		return
	}

	invalidateResultJP(r.Context(), h.cache, clean.Recid)
	w.WriteHeader(http.StatusOK)
}

// DeleteResultJP godoc
//
//	@Summary		Delete Ski Jumping result
//	@Description	Deletes a result by RecID
//	@Tags			FIS - Result Management – Ski Jumping
//	@Accept			json
//	@Produce		json
//	@Param			id	query	int32	true	"Result RecID"
//	@Success		200	"Deleted"
//	@Failure		400	{object}	swagger.ValidationErrorResponse
//	@Failure		401	{object}	swagger.UnauthorizedResponse
//	@Failure		403	{object}	swagger.ForbiddenResponse
//	@Failure		404	{object}	swagger.NotFoundResponse
//	@Failure		500	{object}	swagger.InternalServerErrorResponse
//	@Failure		503	{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/fis/resultjp [delete]
func (h *ResultJPHandler) DeleteResultJP(w http.ResponseWriter, r *http.Request) {
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

	if err := h.store.DeleteResultJPByRecID(r.Context(), id); err != nil {
		if err == sql.ErrNoRows {
			utils.NotFoundResponse(w, r, err)
			return
		}
		utils.HandleDatabaseError(w, r, err)
		return
	}

	invalidateResultJP(r.Context(), h.cache, id)
	w.WriteHeader(http.StatusOK)
}

// GetRaceResultsJP godoc
//
//	@Summary	Get results for a Ski Jumping race
//	@Tags		FIS - Race Results
//	@Accept		json
//	@Produce	json
//	@Param		raceid	query		int32	true	"Race ID"
//	@Success	200		{object}	swagger.FISRaceResultsJPResponse
//	@Failure	400		{object}	swagger.ValidationErrorResponse
//	@Failure	401		{object}	swagger.UnauthorizedResponse
//	@Failure	403		{object}	swagger.ForbiddenResponse
//	@Failure	500		{object}	swagger.InternalServerErrorResponse
//	@Failure	503		{object}	swagger.ServiceUnavailableResponse
//	@Security	BearerAuth
//	@Router		/fis/resultjp [get]
func (h *ResultJPHandler) GetRaceResultsJP(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}
	if err := utils.ValidateParams(r, []string{"raceid"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	raceStr := r.URL.Query().Get("raceid")
	raceID, err := utils.ParsePositiveInt32(raceStr)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	cacheKey := fmt.Sprintf("%s:race=%d", fisResultJPRacePrefix, raceID)
	if h.cache != nil {
		if raw, err := h.cache.Get(r.Context(), cacheKey); err == nil && raw != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(raw))
			return
		}
	}

	rows, err := h.store.GetRaceResultsJPByRaceID(r.Context(), raceID)
	if err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	if len(rows) == 0 {
		utils.NotFoundResponse(w, r, fmt.Errorf("no results found for raceid %d", raceID))
		return
	}

	out := make([]FISResultJPFullResponse, 0, len(rows))
	for _, row := range rows {
		out = append(out, FISResultJPFullFromSqlc(row))
	}

	body := map[string]any{"results": out}
	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, body, FISCacheTTL)
	utils.WriteJSON(w, http.StatusOK, body)
}

// GetAthleteResultsJP godoc
//
//	@Summary	Get Ski Jumping results for an athlete
//	@Tags		FIS - Athlete
//	@Accept		json
//	@Produce	json
//	@Param		fiscode			query		int32		true	"FIS Code"
//	@Param		seasoncode		query		[]int32		false	"Season code (repeat or comma-separated)"
//	@Param		disciplinecode	query		[]string	false	"Discipline code (repeat or comma-separated)"
//	@Param		catcode			query		[]string	false	"Category code (repeat or comma-separated)"
//	@Success	200				{object}	swagger.FISAthleteResultsJPResponse
//	@Failure	400				{object}	swagger.ValidationErrorResponse
//	@Failure	401				{object}	swagger.UnauthorizedResponse
//	@Failure	403				{object}	swagger.ForbiddenResponse
//	@Failure	404				{object}	swagger.NotFoundResponse
//	@Failure	500				{object}	swagger.InternalServerErrorResponse
//	@Failure	503				{object}	swagger.ServiceUnavailableResponse
//	@Security	BearerAuth
//	@Router		/fis/resultathletejp [get]
func (h *ResultJPHandler) GetAthleteResultsJP(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}
	if err := utils.ValidateParams(r, []string{"fiscode", "seasoncode", "disciplinecode", "catcode"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	fiscodeStr := r.URL.Query().Get("fiscode")
	fiscode, err := utils.ParsePositiveInt32(fiscodeStr)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	// get JP competitor ID from fiscode
	competitorID, err := h.competitors.GetCompetitorIDByFiscodeJP(r.Context(), fiscode)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.NotFoundResponse(w, r, fmt.Errorf("competitor with FIS code %d not found", fiscode))
			return
		}
		utils.InternalServerError(w, r, err)
		return
	}

	// filters: seasoncode, disciplinecode, catcode
	seasonsS := parseListParam(r, "seasoncode")
	discs := parseListParam(r, "disciplinecode")
	cats := parseListParam(r, "catcode")

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

	cacheKey := fmt.Sprintf("%s:fis=%d:sc=%v:dc=%v:cc=%v", fisResultJPAthletePrefix, fiscode, seasons, discs, cats)
	if h.cache != nil {
		if raw, err := h.cache.Get(r.Context(), cacheKey); err == nil && raw != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(raw))
			return
		}
	}

	rows, err := h.store.GetAthleteResultsJP(r.Context(), competitorID, seasons, discs, cats)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	out := make([]FISAthleteResultJPRow, 0, len(rows))
	for _, row := range rows {
		out = append(out, FISAthleteResultJPFromSqlc(row))
	}

	body := map[string]any{"results": out}
	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, body, FISCacheTTL)
	utils.WriteJSON(w, http.StatusOK, body)
}
