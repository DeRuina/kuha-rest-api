package utvapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authz"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/DeRuina/KUHA-REST-API/internal/store/utv"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
	"github.com/google/uuid"
)

// Validation structs
type LatestDataInput struct {
	UserID string `form:"user_id" validate:"required,uuid4"`
	Type   string `form:"type" validate:"required,key"`
	Device string `form:"device" validate:"omitempty,oneof=garmin oura polar suunto"`
	Limit  int32  `form:"limit" validate:"omitempty,min=1,max=100"`
}

type AllByTypeInput struct {
	UserID     string `form:"user_id" validate:"required,uuid4"`
	Type       string `form:"type" validate:"required,key"`
	AfterDate  string `form:"after_date" validate:"omitempty,datetime=2006-01-02"`
	BeforeDate string `form:"before_date" validate:"omitempty,datetime=2006-01-02"`
	Limit      int32  `form:"limit" validate:"omitempty,min=1,max=10"`
	Offset     int32  `form:"offset" validate:"omitempty,min=0"`
}

// store and cache interfaces
type GeneralDataHandler struct {
	oura            utv.OuraData
	polar           utv.PolarData
	suunto          utv.SuuntoData
	garmin          utv.GarminData
	ouraToken       utv.OuraToken
	polarToken      utv.PolarToken
	suuntoToken     utv.SuuntoToken
	garminToken     utv.GarminToken
	klabToken       utv.KlabToken
	archinisisToken utv.ArchinisisToken
	cache           *cache.Storage
}

// response structs
type LatestDataResponse struct {
	Device string          `json:"device"`
	Date   string          `json:"date"`
	Data   json.RawMessage `json:"data"`
}

// NewGeneralDataHandler initializes the handler
func NewGeneralDataHandler(
	oura utv.OuraData,
	polar utv.PolarData,
	suunto utv.SuuntoData,
	garmin utv.GarminData,
	ouraToken utv.OuraToken,
	polarToken utv.PolarToken,
	suuntoToken utv.SuuntoToken,
	garminToken utv.GarminToken,
	klabToken utv.KlabToken,
	archinisisToken utv.ArchinisisToken,
	cache *cache.Storage,
) *GeneralDataHandler {
	return &GeneralDataHandler{
		oura:            oura,
		polar:           polar,
		suunto:          suunto,
		garmin:          garmin,
		ouraToken:       ouraToken,
		polarToken:      polarToken,
		suuntoToken:     suuntoToken,
		garminToken:     garminToken,
		klabToken:       klabToken,
		archinisisToken: archinisisToken,
		cache:           cache,
	}
}

// GetLatestData godoc
//
//	@Summary		Get latest data by type
//	@Description	Returns latest entries of a specific type for a user, optionally filtered by device and limited in number (defaults to 1).
//	@Tags			UTV - General
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query	string						true	"User ID (UUID)"
//	@Param			type	query	string						true	"Data type (e.g., 'sleep', 'activity')"
//	@Param			device	query	string						false	"Device type (one of: 'garmin', 'oura', 'polar', 'suunto')"
//	@Param			limit	query	int							false	"Limit the number of results (default: 1, max: 100)"
//	@Success		200		{array}	swagger.LatestDataResponse	"Latest Data"
//	@Success		204		"No Content: No data found"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/utv/latest [get]
func (h *GeneralDataHandler) GetLatestData(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	err := utils.ValidateParams(r, []string{"user_id", "type", "device", "limit"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := LatestDataInput{
		UserID: r.URL.Query().Get("user_id"),
		Type:   r.URL.Query().Get("type"),
		Device: r.URL.Query().Get("device"),
	}
	if val := r.URL.Query().Get("limit"); val != "" {
		parsed, err := strconv.Atoi(val)
		if err != nil {
			utils.BadRequestResponse(w, r, fmt.Errorf("limit must be a number"))
			return
		}
		params.Limit = int32(parsed)
	}
	if params.Limit == 0 {
		params.Limit = 1 // default
	}

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	userID, err := utils.ParseUUID(params.UserID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	cacheKey := fmt.Sprintf("utv:latest:%s:%s:%s:%d", params.UserID, params.Type, params.Device, params.Limit)

	if h.cache != nil {
		if cached, err := h.cache.Get(r.Context(), cacheKey); err == nil && cached != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(cached))
			return
		}
	}

	var results []LatestDataResponse

	// Helper to fetch from one device
	fetch := func(name string, store interface {
		GetLatestByType(ctx context.Context, userID uuid.UUID, typ string, limit int32) ([]utv.LatestDataEntry, error)
	}) {
		data, err := store.GetLatestByType(r.Context(), userID, params.Type, params.Limit)
		if err != nil {
			return // silently ignore errors per device
		}
		for _, row := range data {
			results = append(results, LatestDataResponse{
				Device: name,
				Date:   row.Date.Format("2006-01-02"),
				Data:   row.Data,
			})
		}
	}

	// Conditional fetch
	switch params.Device {
	case "garmin":
		fetch("garmin", h.garmin)
	case "oura":
		fetch("oura", h.oura)
	case "polar":
		fetch("polar", h.polar)
	case "suunto":
		fetch("suunto", h.suunto)
	default:
		fetch("garmin", h.garmin)
		fetch("oura", h.oura)
		fetch("polar", h.polar)
		fetch("suunto", h.suunto)
	}

	if len(results) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, results, 3*time.Minute)

	utils.WriteJSON(w, http.StatusOK, results)
}

// GetAllByType godoc
//
//	@Summary		Get all data by type
//	@Description	Returns all entries of a specific data type for a user, across all wearable devices, optionally filtered by date range, paginated by limit (default 3) and offset.
//	@Tags			UTV - General
//	@Accept			json
//	@Produce		json
//	@Param			user_id		query	string						true	"User ID (UUID)"
//	@Param			type		query	string						true	"Data type (e.g., 'sleep', 'activity')"
//	@Param			after_date	query	string						false	"Filter data after this date (YYYY-MM-DD)"
//	@Param			before_date	query	string						false	"Filter data before this date (YYYY-MM-DD)"
//	@Param			limit		query	int							false	"Limit the number of results returned (default: 3, max: 10)"
//	@Param			offset		query	int							false	"Offset for pagination (default: 0)"
//
//	@Success		200			{array}	swagger.LatestDataResponse	"List of data entries across devices"
//	@Success		204			"No Content: No data available"
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		401			{object}	swagger.UnauthorizedResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		422			{object}	swagger.InvalidDateRange
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
//	@Failure		503			{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/utv/all [get]
func (h *GeneralDataHandler) GetAllByType(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	err := utils.ValidateParams(r, []string{"user_id", "type", "after_date", "before_date", "limit", "offset"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := AllByTypeInput{
		UserID:     r.URL.Query().Get("user_id"),
		Type:       r.URL.Query().Get("type"),
		AfterDate:  r.URL.Query().Get("after_date"),
		BeforeDate: r.URL.Query().Get("before_date"),
	}

	if val := r.URL.Query().Get("limit"); val != "" {
		parsed, err := strconv.Atoi(val)
		if err != nil {
			utils.BadRequestResponse(w, r, fmt.Errorf("limit must be a number"))
			return
		}
		params.Limit = int32(parsed)
	}
	if params.Limit == 0 {
		params.Limit = 3 // default
	}

	// Parse optional offset
	if val := r.URL.Query().Get("offset"); val != "" {
		parsed, err := strconv.Atoi(val)
		if err != nil {
			utils.BadRequestResponse(w, r, fmt.Errorf("offset must be a number"))
			return
		}
		params.Offset = int32(parsed)
	}

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	// Parse values
	userID, err := utils.ParseUUID(params.UserID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	after, err := utils.ParseDatePtr(utils.NilIfEmpty(&params.AfterDate))
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	before, err := utils.ParseDatePtr(utils.NilIfEmpty(&params.BeforeDate))
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	if after != nil && before != nil && after.After(*before) {
		utils.UnprocessableEntityResponse(w, r, utils.ErrInvalidDateRange)
		return
	}

	cacheKey := fmt.Sprintf("utv:all:%s:%s:after:%s:before:%s,limit:%d,offset:%d", userID, params.Type, after, before, params.Limit, params.Offset)

	if h.cache != nil {
		if cached, err := h.cache.Get(r.Context(), cacheKey); err == nil && cached != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(cached))
			return
		}
	}

	var results []LatestDataResponse

	// Helper to query one device with pagination
	fetch := func(name string, store interface {
		GetAllByType(ctx context.Context, userID uuid.UUID, typ string, after, before *time.Time, limit, offset int32) ([]utv.LatestDataEntry, error)
	}) {
		data, err := store.GetAllByType(r.Context(), userID, params.Type, after, before, params.Limit, params.Offset)
		if err != nil {
			return
		}
		for _, row := range data {
			results = append(results, LatestDataResponse{
				Device: name,
				Date:   row.Date.Format("2006-01-02"),
				Data:   row.Data,
			})
		}
	}

	// Query all 4 devices
	fetch("garmin", h.garmin)
	fetch("oura", h.oura)
	fetch("polar", h.polar)
	fetch("suunto", h.suunto)

	if len(results) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, results, 3*time.Minute)

	utils.WriteJSON(w, http.StatusOK, results)
}

type DisconnectParams struct {
	UserID string `form:"user_id" validate:"required,uuid4"`
	Source string `form:"source" validate:"required,oneof=polar oura suunto garmin klab archinisis"`
}

// Disconnect godoc
//
//	@Summary		Disconnect a wearable device
//	@Description	Disconnects a user's wearable device by deleting the associated token
//	@Tags			UTV - General
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query	string	true	"User ID (UUID)"
//	@Param			source	query	string	true	"Source device to disconnect (one of: 'polar', 'oura', 'suunto', 'garmin', 'klab', 'archinisis')"
//	@Success		200		"Successfully disconnected"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/utv/disconnect [delete]
func (h *GeneralDataHandler) Disconnect(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	err := utils.ValidateParams(r, []string{"user_id", "source"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := DisconnectParams{
		UserID: r.URL.Query().Get("user_id"),
		Source: r.URL.Query().Get("source"),
	}

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	userID, err := utils.ParseUUID(params.UserID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	var delErr error
	switch params.Source {
	case "polar":
		delErr = h.polarToken.DeleteToken(r.Context(), userID)
	case "oura":
		delErr = h.ouraToken.DeleteToken(r.Context(), userID)
	case "suunto":
		delErr = h.suuntoToken.DeleteToken(r.Context(), userID)
	case "garmin":
		delErr = h.garminToken.DeleteToken(r.Context(), userID)
	case "klab":
		delErr = h.klabToken.DeleteToken(r.Context(), userID)
	case "archinisis":
		delErr = h.archinisisToken.DeleteToken(r.Context(), userID)
	}

	if delErr != nil {
		utils.InternalServerError(w, r, delErr)
		return
	}

	w.WriteHeader(http.StatusOK)
}

type TokensForUpdateParams struct {
	Source string `form:"source" validate:"required,oneof=polar oura suunto garmin"`
	Hours  int    `form:"hours" validate:"required,min=1,max=8760"` // up to 1 year
}

type UserDataResponse struct {
	UserID string      `json:"user_id"`
	Data   interface{} `json:"data"`
}

// GetTokensForUpdate godoc
//
//	@Summary		Get tokens for update
//	@Description	Retrieves tokens that need to be updated based on the source and time cutoff
//	@Tags			UTV - General
//	@Accept			json
//	@Produce		json
//	@Param			source	query	string					true	"Source device (one of: 'polar', 'oura', 'suunto', 'garmin')"
//	@Param			hours	query	int						true	"Number of hours to look back (1-8760)"
//	@Success		200		{array}	swagger.PolarTokenInput	"List of tokens needing update"
//	@Success		204		"No Content: No tokens found"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/utv/tokens4update [get]
func (h *GeneralDataHandler) GetTokensForUpdate(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{"source", "hours"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := TokensForUpdateParams{
		Source: r.URL.Query().Get("source"),
	}

	if hoursStr := r.URL.Query().Get("hours"); hoursStr != "" {
		hours, err := strconv.Atoi(hoursStr)
		if err != nil {
			utils.BadRequestResponse(w, r, fmt.Errorf("invalid hours format"))
			return
		}
		params.Hours = hours
	}

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	cutoff := time.Now().Add(-time.Duration(params.Hours) * time.Hour)

	switch params.Source {
	case "polar":
		tokens, err := h.polarToken.GetTokensForUpdate(r.Context(), cutoff)
		if err != nil {
			utils.InternalServerError(w, r, err)
			return
		}
		if len(tokens) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		resp := make([]UserDataResponse, 0, len(tokens))
		for _, t := range tokens {
			resp = append(resp, UserDataResponse{UserID: t.UserID.String(), Data: t.Data})
		}
		utils.WriteJSON(w, http.StatusOK, resp)

	case "oura":
		tokens, err := h.ouraToken.GetTokensForUpdate(r.Context(), cutoff)
		if err != nil {
			utils.InternalServerError(w, r, err)
			return
		}
		if len(tokens) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		resp := make([]UserDataResponse, 0, len(tokens))
		for _, t := range tokens {
			resp = append(resp, UserDataResponse{UserID: t.UserID.String(), Data: t.Data})
		}
		utils.WriteJSON(w, http.StatusOK, resp)

	case "suunto":
		tokens, err := h.suuntoToken.GetTokensForUpdate(r.Context(), cutoff)
		if err != nil {
			utils.InternalServerError(w, r, err)
			return
		}
		if len(tokens) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		resp := make([]UserDataResponse, 0, len(tokens))
		for _, t := range tokens {
			resp = append(resp, UserDataResponse{UserID: t.UserID.String(), Data: t.Data})
		}
		utils.WriteJSON(w, http.StatusOK, resp)

	case "garmin":
		tokens, err := h.garminToken.GetTokensForUpdate(r.Context(), cutoff)
		if err != nil {
			utils.InternalServerError(w, r, err)
			return
		}
		if len(tokens) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		resp := make([]UserDataResponse, 0, len(tokens))
		for _, t := range tokens {
			resp = append(resp, UserDataResponse{UserID: t.UserID.String(), Data: t.Data})
		}
		utils.WriteJSON(w, http.StatusOK, resp)

	default:
		utils.BadRequestResponse(w, r, fmt.Errorf("invalid source: must be one of polar, oura, suunto, garmin"))
	}
}

type DataForUpdateParams struct {
	Source string `form:"source" validate:"required,oneof=polar oura suunto garmin"`
	Hours  int    `form:"hours" validate:"required,min=1,max=8760"`
}

// GetDataForUpdate godoc
//
//	@Summary		Get data for update
//	@Description	Returns token records where 'data_last_fetched' is older than the cutoff
//	@Tags			UTV - General
//	@Accept			json
//	@Produce		json
//	@Param			source	query	string					true	"Source (one of: 'polar', 'oura', 'suunto', 'garmin')"
//	@Param			hours	query	int						true	"Number of hours to look back (1-8760)"
//	@Success		200		{array}	swagger.PolarTokenInput	"List of tokens needing data update"
//	@Success		204		"No Content: No data found"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/utv/data4update [get]
func (h *GeneralDataHandler) GetDataForUpdate(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{"source", "hours"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := DataForUpdateParams{
		Source: r.URL.Query().Get("source"),
	}

	if hoursStr := r.URL.Query().Get("hours"); hoursStr != "" {
		hours, err := strconv.Atoi(hoursStr)
		if err != nil {
			utils.BadRequestResponse(w, r, fmt.Errorf("invalid hours format"))
			return
		}
		params.Hours = hours
	}

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	cutoff := time.Now().Add(-time.Duration(params.Hours) * time.Hour)

	switch params.Source {
	case "polar":
		data, err := h.polarToken.GetDataForUpdate(r.Context(), cutoff)
		if err != nil {
			utils.InternalServerError(w, r, err)
			return
		}
		if len(data) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		resp := make([]UserDataResponse, 0, len(data))
		for _, t := range data {
			resp = append(resp, UserDataResponse{UserID: t.UserID.String(), Data: t.Data})
		}
		utils.WriteJSON(w, http.StatusOK, resp)

	case "oura":
		data, err := h.ouraToken.GetDataForUpdate(r.Context(), cutoff)
		if err != nil {
			utils.InternalServerError(w, r, err)
			return
		}
		if len(data) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		resp := make([]UserDataResponse, 0, len(data))
		for _, t := range data {
			resp = append(resp, UserDataResponse{UserID: t.UserID.String(), Data: t.Data})
		}
		utils.WriteJSON(w, http.StatusOK, resp)

	case "suunto":
		data, err := h.suuntoToken.GetDataForUpdate(r.Context(), cutoff)
		if err != nil {
			utils.InternalServerError(w, r, err)
			return
		}
		if len(data) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		resp := make([]UserDataResponse, 0, len(data))
		for _, t := range data {
			resp = append(resp, UserDataResponse{UserID: t.UserID.String(), Data: t.Data})
		}
		utils.WriteJSON(w, http.StatusOK, resp)

	case "garmin":
		data, err := h.garminToken.GetDataForUpdate(r.Context(), cutoff)
		if err != nil {
			utils.InternalServerError(w, r, err)
			return
		}
		if len(data) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		resp := make([]UserDataResponse, 0, len(data))
		for _, t := range data {
			resp = append(resp, UserDataResponse{UserID: t.UserID.String(), Data: t.Data})
		}
		utils.WriteJSON(w, http.StatusOK, resp)

	default:
		utils.BadRequestResponse(w, r, fmt.Errorf("invalid source: must be one of polar, oura, suunto, garmin"))
	}
}
