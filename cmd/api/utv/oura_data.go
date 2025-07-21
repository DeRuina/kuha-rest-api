package utvapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authz"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/DeRuina/KUHA-REST-API/internal/store/utv"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

// Validation structs
type OuraGetDatesParams struct {
	UserID     string `form:"user_id" validate:"required,uuid4"`
	AfterDate  string `form:"after_date" validate:"omitempty,datetime=2006-01-02"`
	BeforeDate string `form:"before_date" validate:"omitempty,datetime=2006-01-02"`
}

type OuraGetTypesParams struct {
	UserID string `form:"user_id" validate:"required,uuid4"`
	Date   string `form:"date" validate:"required,datetime=2006-01-02"`
}

type OuraGetDataParams struct {
	UserID string `form:"user_id" validate:"required,uuid4"`
	Date   string `form:"date" validate:"required,datetime=2006-01-02"`
	Key    string `form:"key" validate:"omitempty,key"`
}

type OuraPostDataInput struct {
	UserID string          `json:"user_id" validate:"required,uuid4"`
	Date   string          `json:"date" validate:"required,datetime=2006-01-02"`
	Data   json.RawMessage `json:"data" validate:"required"`
}

type OuraDeleteDataInput struct {
	UserID string `json:"user_id" validate:"required,uuid4"`
	Date   string `json:"date" validate:"required,datetime=2006-01-02"`
}

type DeleteAllOuraParams struct {
	UserID string `validate:"required,uuid4" json:"user_id"`
}

// store and cache interfaces
type OuraDataHandler struct {
	store utv.OuraData
	cache *cache.Storage
}

// NewOuraDataHandler initializes OuraData handler
func NewOuraDataHandler(store utv.OuraData, cache *cache.Storage) *OuraDataHandler {
	return &OuraDataHandler{store: store, cache: cache}
}

// GetDatesOura godoc
//
//	@Summary		Get available dates
//	@Description	Returns available dates for the specified user (optionally filtered by date range)
//	@Tags			UTV - Oura
//	@Accept			json
//	@Produce		json
//	@Param			user_id		query		string					true	"User ID (UUID)"
//	@Param			after_date	query		string					false	"Filter dates after this date (YYYY-MM-DD)"
//	@Param			before_date	query		string					false	"Filter dates before this date (YYYY-MM-DD)"
//	@Success		200			{object}	swagger.DatesResponse	"List of available dates"
//	@Success		204			"No Content: No available dates found"
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		422			{object}	swagger.InvalidDateRange
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
//	@Security		BearerAuth
//	@Router			/utv/oura/dates [get]
func (h *OuraDataHandler) GetDates(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	err := utils.ValidateParams(r, []string{"user_id", "after_date", "before_date"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := OuraGetDatesParams{
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

	cacheKey := fmt.Sprintf("oura:dates:%s:%s:%s", params.UserID, params.AfterDate, params.BeforeDate)

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

	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, response, 10*time.Minute)

	utils.WriteJSON(w, http.StatusOK, response)
}

// GetTypesOura godoc

// @Summary		Get available types
// @Description	Returns available types for the specified user on the specified date
// @Tags			UTV - Oura
// @Accept			json
// @Produce		json
// @Param			user_id	query		string						true	"User ID (UUID)"
// @Param			date	query		string						true	"Date (YYYY-MM-DD)"
// @Success		200		{object}	swagger.OuraTypesResponse	"List of available types"
// @Success		204		"No Content: No available types found"
// @Failure		400		{object}	swagger.ValidationErrorResponse
// @Failure		403		{object}	swagger.ForbiddenResponse
// @Failure		500		{object}	swagger.InternalServerErrorResponse
// @Security		BearerAuth
// @Router			/utv/oura/types [get]
func (h *OuraDataHandler) GetTypes(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	err := utils.ValidateParams(r, []string{"user_id", "date"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := OuraGetTypesParams{
		UserID: r.URL.Query().Get("user_id"),
		Date:   r.URL.Query().Get("date"),
	}

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	cacheKey := fmt.Sprintf("oura:types:%s:%s", params.UserID, params.Date)

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

	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, response, 10*time.Minute)

	utils.WriteJSON(w, http.StatusOK, response)
}

// GetDataOura godoc

// @Summary		Get available data
// @Description	Returns data for the specified user on the specified date (optionally filtered by key)
// @Tags			UTV - Oura
// @Accept			json
// @Produce		json
// @Param			user_id	query		string						true	"User ID (UUID)"
// @Param			date	query		string						true	"Date (YYYY-MM-DD)"
// @Param			key		query		string						false	"Type"
// @Success		200		{object}	swagger.OuraDataResponse	"Data"
// @Success		204		"No Content: No data found"
// @Failure		400		{object}	swagger.ValidationErrorResponse
// @Failure		403		{object}	swagger.ForbiddenResponse
// @Failure		500		{object}	swagger.InternalServerErrorResponse
// @Security		BearerAuth
// @Router			/utv/oura/data [get]
func (h *OuraDataHandler) GetData(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	err := utils.ValidateParams(r, []string{"user_id", "date", "key"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := OuraGetDataParams{
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
	cacheKey := fmt.Sprintf("oura:data:%s:%s:%s", params.UserID, params.Date, keyPart)

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

	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, response, 10*time.Minute)

	utils.WriteJSON(w, http.StatusOK, response)
}

// PostDataOura godoc
//
//	@Summary		Post Oura data
//	@Description	Posts Oura data for the specified user on the specified date
//	@Tags			UTV - Oura
//	@Accept			json
//	@Produce		json
//	@Param			body	body	swagger.OuraPostDataInput	true	"Oura data input"
//	@Success		201		"Created: Data successfully stored (no content in response body)"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Security		BearerAuth
//	@Router			/utv/oura/data [post]
func (h *OuraDataHandler) InsertData(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var input OuraPostDataInput
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

	w.WriteHeader(http.StatusCreated)
}

// DeleteAllDataOura godoc
//
// @Summary		Delete all Oura data for a user
// @Description	Deletes all Oura data entries for a specific user
// @Tags			UTV - Oura
// @Accept			json
// @Produce		json
// @Param			user_id	query	string	true	"User ID (UUID)"
// @Success		200		"Deleted: Data successfully removed"
// @Success		204		"No Content: No matching data"
// @Failure		400		{object} swagger.ValidationErrorResponse
// @Failure		403		{object} swagger.ForbiddenResponse
// @Failure		500		{object} swagger.InternalServerErrorResponse
// @Security		BearerAuth
// @Router			/utv/oura/data [delete]
func (h *OuraDataHandler) DeleteAllData(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	err := utils.ValidateParams(r, []string{"user_id"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := DeleteAllOuraParams{UserID: r.URL.Query().Get("user_id")}

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

	w.WriteHeader(http.StatusOK)
}
