package archapi

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authz"
	archsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/archinisis"
	"github.com/DeRuina/KUHA-REST-API/internal/store/archinisis"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type DataHandler struct {
	store archinisis.Data
	cache *cache.Storage
}

func NewDataHandler(store archinisis.Data, cache *cache.Storage) *DataHandler {
	return &DataHandler{store: store, cache: cache}
}

type RaceReportSessionsQuery struct {
	SporttiID string `validate:"required,numeric"`
}

type RaceReportHTMLQuery struct {
	SporttiID string `validate:"required,numeric"`
	SessionID string `validate:"required,numeric"`
}

// GetRaceReportSessions godoc
//
//	@Summary		List race-report session IDs for a Sportti ID
//	@Description	Returns all session_id values that have race reports for the given sportti_id
//	@Tags			ARCHINISIS - Data
//	@Accept			json
//	@Produce		json
//	@Param			sportti_id	query		string	true	"Sportti ID"
//	@Success		200			{object}	swagger.RaceReportSessionsResponse
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		401			{object}	swagger.UnauthorizedResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
//	@Failure		503			{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/archinisis/race-report/sessions [get]
func (h *DataHandler) GetRaceReportSessions(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{"sportti_id"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	q := RaceReportSessionsQuery{
		SporttiID: r.URL.Query().Get("sportti_id"),
	}
	if err := utils.GetValidator().Struct(q); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	sid, err := utils.ParseSporttiID(q.SporttiID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	cacheKey := fmt.Sprintf("arch:race-report:sessions:%s", sid)

	if h.cache != nil {
		if cached, err := h.cache.Get(r.Context(), cacheKey); err == nil && cached != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(cached))
			return
		}
	}

	sessionIDs, err := h.store.GetRaceReportSessions(r.Context(), sid)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, map[string]any{"race_report": sessionIDs}, ARCHCacheTTL)

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"race_report": sessionIDs,
	})
}

// GetRaceReportHTML godoc
//
//	@Summary		Get a specific race report (HTML)
//	@Description	Returns the raw HTML race report for a (sportti_id, session_id). Content-Type is text/html.
//	@Tags			ARCHINISIS - Data
//	@Accept			json
//	@Produce		html
//	@Param			sportti_id	query		string	true	"Sportti ID"
//	@Param			session_id	query		string	true	"Session ID"
//	@Success		200			{string}	string	"<!DOCTYPE html><html><head><title>Race Report</title></head><body><h1>HTML RACE REPORT</h1><p>full report returned in html DOCTYPE</p></body></html>"
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		401			{object}	swagger.UnauthorizedResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		404			{object}	swagger.NotFoundResponse
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
//	@Failure		503			{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/archinisis/race-report [get]
func (h *DataHandler) GetRaceReportHTML(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{"sportti_id", "session_id"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	q := RaceReportHTMLQuery{
		SporttiID: r.URL.Query().Get("sportti_id"),
		SessionID: r.URL.Query().Get("session_id"),
	}
	if err := utils.GetValidator().Struct(q); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	sid, err := utils.ParseSporttiID(q.SporttiID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	sessionID, err := utils.ParsePositiveInt32(q.SessionID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	cacheKey := fmt.Sprintf("arch:race-report:html:%s:%d", sid, sessionID)

	if h.cache != nil {
		if cached, err := h.cache.Get(r.Context(), cacheKey); err == nil && cached != "" {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(cached))
			return
		}
	}

	html, err := h.store.GetRaceReport(r.Context(), sid, sessionID)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.NotFoundResponse(w, r, err)
			return
		}
		utils.InternalServerError(w, r, err)
		return
	}

	if h.cache != nil {
		_ = h.cache.Set(r.Context(), cacheKey, html, ARCHCacheTTL)
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(html))
}

type RaceReportUpsertInput struct {
	SporttiID  string `json:"sportti_id" validate:"required,numeric"`
	SessionID  int32  `json:"session_id" validate:"required,gt=0"`
	RaceReport string `json:"race_report" validate:"required"`
}

// PostRaceReport godoc
//
//	@Summary		Upsert a race report (HTML)
//	@Description	Inserts or updates a race report for (sportti_id, session_id).
//	@Tags			ARCHINISIS - Data
//	@Accept			json
//	@Produce		json
//	@Param			data	body	swagger.ArchRaceReportUpsertRequest	true	"race report"
//	@Success		201		"Data processed successfully"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/archinisis/race-report [post]
func (h *DataHandler) PostRaceReport(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var in RaceReportUpsertInput
	if err := utils.ReadJSON(w, r, &in); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	if err := utils.GetValidator().Struct(in); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	sid, err := utils.ParseSporttiID(in.SporttiID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := archsqlc.UpsertRaceReportParams{
		SporttiID:  utils.NullString(sid),
		SessionID:  utils.NullInt32(in.SessionID),
		RaceReport: utils.NullString(in.RaceReport),
	}

	if err := h.store.UpsertRaceReport(r.Context(), params); err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	invalidateArchRaceReport(r.Context(), h.cache, sid, &in.SessionID)

	w.WriteHeader(http.StatusCreated)
}

// PostArchData godoc
//
//	@Summary		Upsert Archinisis athlete + measurements
//	@Description	Inserts/updates athlete profile and related measurements in one request.
//	@Tags			ARCHINISIS - Data
//	@Accept			json
//	@Produce		json
//	@Param			data	body	swagger.ArchDataUpsertRequest	true	"athlete + measurements"
//	@Success		201		"Data processed successfully"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/archinisis/data [post]
func (h *DataHandler) PostArchData(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var in ArchDataUpsertInput
	if err := utils.ReadJSON(w, r, &in); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	if err := utils.GetValidator().Struct(in); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	sid, err := utils.ParseSporttiID(in.NationalID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	ath, err := mapAthleteToParams(in, sid)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	measParams := make([]archsqlc.UpsertMeasurementParams, 0, len(in.Measurements))
	for _, m := range in.Measurements {
		mp, err := mapMeasurementToParams(m, sid)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}
		measParams = append(measParams, mp)
	}

	payload := archinisis.ArchDataPayload{
		Athlete:      ath,
		Measurements: measParams,
	}

	if err := h.store.UpsertData(r.Context(), payload); err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	invalidateArchData(r.Context(), h.cache, sid)

	w.WriteHeader(http.StatusCreated)
}

type archUserParams struct {
	ID string `validate:"required,numeric"`
}

// GetArchData godoc
//
//	@Summary		Get Archinisis data by Sportti ID
//	@Description	Returns athlete profile and measurements for the given sportti id.
//	@Tags			ARCHINISIS - Data
//	@Accept			json
//	@Produce		json
//	@Param			id	query		string	true	"Sportti ID (national_id)"
//	@Success		200	{object}	swagger.ArchDataResponse
//	@Failure		400	{object}	swagger.ValidationErrorResponse
//	@Failure		401	{object}	swagger.UnauthorizedResponse
//	@Failure		403	{object}	swagger.ForbiddenResponse
//	@Failure		404	{object}	swagger.NotFoundResponse
//	@Failure		500	{object}	swagger.InternalServerErrorResponse
//	@Failure		503	{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/archinisis/data [get]
func (h *DataHandler) GetArchData(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{"id"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := archUserParams{
		ID: r.URL.Query().Get("id"),
	}

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	sid, err := utils.ParseSporttiID(params.ID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	cacheKey := fmt.Sprintf("arch:data:%s", sid)
	if h.cache != nil {
		if cached, err := h.cache.Get(r.Context(), cacheKey); err == nil && cached != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(cached))
			return
		}
	}

	res, err := h.store.GetDataBySporttiID(r.Context(), sid)
	if err == sql.ErrNoRows {
		utils.NotFoundResponse(w, r, err)
		return
	}
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, res, ARCHCacheTTL)

	utils.WriteJSON(w, http.StatusOK, res)
}
