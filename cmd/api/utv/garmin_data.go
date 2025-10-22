package utvapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authz"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/DeRuina/KUHA-REST-API/internal/store/utv"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

// Validation structs
type GarminGetDatesParams struct {
	UserID     string `form:"user_id" validate:"required,uuid4"`
	AfterDate  string `form:"after_date" validate:"omitempty,datetime=2006-01-02"`
	BeforeDate string `form:"before_date" validate:"omitempty,datetime=2006-01-02"`
}

type GarminGetTypesParams struct {
	UserID string `form:"user_id" validate:"required,uuid4"`
	Date   string `form:"date" validate:"required,datetime=2006-01-02"`
}

type GarminGetDataParams struct {
	UserID string `form:"user_id" validate:"required,uuid4"`
	Date   string `form:"date" validate:"required,datetime=2006-01-02"`
	Key    string `form:"key" validate:"omitempty,key"`
}

type GarminPostDataInput struct {
	UserID string          `json:"user_id" validate:"required,uuid4"`
	Date   string          `json:"date" validate:"required,datetime=2006-01-02"`
	Data   json.RawMessage `json:"data" validate:"required"`
}

type DeleteAllGarminParams struct {
	UserID string `validate:"required,uuid4" json:"user_id"`
}

// store and cache interfaces
type GarminDataHandler struct {
	store utv.GarminData
	cache *cache.Storage
}

// NewGarminDataHandler initializes GarminData handler
func NewGarminDataHandler(store utv.GarminData, cache *cache.Storage) *GarminDataHandler {
	return &GarminDataHandler{store: store, cache: cache}
}

// GetDatesGarmin godoc
//
//	@Summary		Get available dates
//	@Description	Returns available dates for the specified user (optionally filtered by date range)
//	@Tags			UTV - Garmin
//	@Accept			json
//	@Produce		json
//	@Param			user_id		query		string					true	"User ID (UUID)"
//	@Param			after_date	query		string					false	"Filter dates after this date (YYYY-MM-DD)"
//	@Param			before_date	query		string					false	"Filter dates before this date (YYYY-MM-DD)"
//	@Success		200			{object}	swagger.DatesResponse	"List of available dates"
//	@Success		204			"No Content: No available dates found"
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		401			{object}	swagger.UnauthorizedResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		422			{object}	swagger.InvalidDateRange
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
//	@Failure		503			{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/utv/garmin/dates [get]
func (h *GarminDataHandler) GetDates(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	err := utils.ValidateParams(r, []string{"user_id", "after_date", "before_date"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := GarminGetDatesParams{
		UserID:     r.URL.Query().Get("user_id"),
		AfterDate:  r.URL.Query().Get("after_date"),
		BeforeDate: r.URL.Query().Get("before_date"),
	}

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	if params.AfterDate != "" && params.BeforeDate != "" && params.AfterDate > params.BeforeDate {
		utils.UnprocessableEntityResponse(w, r, utils.ErrInvalidDateRange)
		return
	}

	cacheKey := fmt.Sprintf("utv:garmin:dates:%s:%s:%s", params.UserID, params.AfterDate, params.BeforeDate)

	if h.cache != nil {
		if cached, err := h.cache.Get(r.Context(), cacheKey); err == nil && cached != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(cached))
			return
		}
	}

	dates, err := h.store.GetDates(r.Context(), params.UserID, &params.AfterDate, &params.BeforeDate)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	if len(dates) == 0 {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	response := map[string]interface{}{
		"dates": dates,
	}

	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, response, UTVCacheTTL)

	utils.WriteJSON(w, http.StatusOK, response)
}

// GetTypesGarmin godoc

// @Summary		Get available types
// @Description	Returns available types for the specified user on the specified date
// @Tags			UTV - Garmin
// @Accept			json
// @Produce		json
// @Param			user_id	query		string						true	"User ID (UUID)"
// @Param			date	query		string						true	"Date (YYYY-MM-DD)"
// @Success		200		{object}	swagger.GarminTypesResponse	"List of available types"
// @Success		204		"No Content: No available types found"
// @Failure		400		{object}	swagger.ValidationErrorResponse
// @Failure		401		{object}	swagger.UnauthorizedResponse
// @Failure		403		{object}	swagger.ForbiddenResponse
// @Failure		500		{object}	swagger.InternalServerErrorResponse
// @Failure		503		{object}	swagger.ServiceUnavailableResponse
// @Security		BearerAuth
// @Router			/utv/garmin/types [get]
func (h *GarminDataHandler) GetTypes(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	err := utils.ValidateParams(r, []string{"user_id", "date"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := GarminGetTypesParams{
		UserID: r.URL.Query().Get("user_id"),
		Date:   r.URL.Query().Get("date"),
	}

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	cacheKey := fmt.Sprintf("utv:garmin:types:%s:%s", params.UserID, params.Date)

	if h.cache != nil {
		if cached, err := h.cache.Get(r.Context(), cacheKey); err == nil && cached != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(cached))
			return
		}
	}

	types, err := h.store.GetTypes(r.Context(), params.UserID, params.Date)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	if len(types) == 0 {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	response := map[string]interface{}{
		"types": types,
	}

	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, response, UTVCacheTTL)

	utils.WriteJSON(w, http.StatusOK, response)
}

// GetDataGarmin godoc

// @Summary		Get available data
// @Description	Returns data for the specified user on the specified date (optionally filtered by key)
// @Tags			UTV - Garmin
// @Accept			json
// @Produce		json
// @Param			user_id	query		string						true	"User ID (UUID)"
// @Param			date	query		string						true	"Date (YYYY-MM-DD)"
// @Param			key		query		string						false	"Type"
// @Success		200		{object}	swagger.GarminDataResponse	"Data"
// @Success		204		"No Content: No data found"
// @Failure		400		{object}	swagger.ValidationErrorResponse
// @Failure		401		{object}	swagger.UnauthorizedResponse
// @Failure		403		{object}	swagger.ForbiddenResponse
// @Failure		500		{object}	swagger.InternalServerErrorResponse
// @Failure		503		{object}	swagger.ServiceUnavailableResponse
// @Security		BearerAuth
// @Router			/utv/garmin/data [get]
func (h *GarminDataHandler) GetData(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	err := utils.ValidateParams(r, []string{"user_id", "date", "key"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := GarminGetDataParams{
		UserID: r.URL.Query().Get("user_id"),
		Date:   r.URL.Query().Get("date"),
		Key:    r.URL.Query().Get("key"),
	}

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	keyPart := "all"
	if params.Key != "" {
		keyPart = params.Key
	}
	cacheKey := fmt.Sprintf("utv:garmin:data:%s:%s:%s", params.UserID, params.Date, keyPart)

	if h.cache != nil {
		if cached, err := h.cache.Get(r.Context(), cacheKey); err == nil && cached != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(cached))
			return
		}
	}

	data, err := h.store.GetData(r.Context(), params.UserID, params.Date, utils.NilIfEmpty(&params.Key))
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	if len(data) == 0 {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	response := map[string]interface{}{
		"data": data,
	}

	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, response, UTVCacheTTL)

	utils.WriteJSON(w, http.StatusOK, response)
}

// PostDataGarmin godoc
//
//	@Summary		Post Garmin data
//	@Description	Posts Garmin data for the specified user on the specified date
//	@Tags			UTV - Garmin
//	@Accept			json
//	@Produce		json
//	@Param			body	body	swagger.GarminPostDataInput	true	"Garmin data input"
//	@Success		201		"Created: Data successfully stored (no content in response body)"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/utv/garmin/data [post]
func (h *GarminDataHandler) InsertData(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var input GarminPostDataInput
	if err := utils.ReadJSON(w, r, &input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	if err := utils.GetValidator().Struct(input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	userID, err := utils.ParseUUID(input.UserID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	date, err := utils.ParseDate(input.Date)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	err = h.store.InsertData(r.Context(), userID, date, input.Data)
	if err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	invalidateUTVSource(r.Context(), h.cache, userID, "garmin")

	w.WriteHeader(http.StatusCreated)
}

// DeleteAllDataGarmin godoc
//
// @Summary		Delete all Garmin data for a user
// @Description	Deletes all Garmin data entries for a specific user
// @Tags			UTV - Garmin
// @Accept			json
// @Produce		json
// @Param			user_id	query	string	true	"User ID (UUID)"
// @Success		200		"Deleted: Data successfully removed"
// @Success		204		"No Content: No matching data"
// @Failure		400		{object} swagger.ValidationErrorResponse
// @Failure		401		{object} swagger.UnauthorizedResponse
// @Failure		403		{object} swagger.ForbiddenResponse
// @Failure		500		{object} swagger.InternalServerErrorResponse
// @Failure		503		{object} swagger.ServiceUnavailableResponse
// @Security		BearerAuth
// @Router			/utv/garmin/data [delete]
func (h *GarminDataHandler) DeleteAllData(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	err := utils.ValidateParams(r, []string{"user_id"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := DeleteAllGarminParams{UserID: r.URL.Query().Get("user_id")}

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	userID, err := utils.ParseUUID(params.UserID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	rows, err := h.store.DeleteAllData(r.Context(), userID)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	if rows == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	invalidateUTVSource(r.Context(), h.cache, userID, "garmin")

	w.WriteHeader(http.StatusOK)
}
