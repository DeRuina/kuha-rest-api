package utvapi

import (
	"context"
	"net/http"

	"github.com/DeRuina/KUHA-REST-API/internal/store/utv"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

// Validation structs
type PolarGetDatesParams struct {
	UserID     string `form:"user_id" validate:"required,uuid4"`
	AfterDate  string `form:"after_date" validate:"omitempty,datetime=2006-01-02"`
	BeforeDate string `form:"before_date" validate:"omitempty,datetime=2006-01-02"`
}

type PolarGetTypesParams struct {
	UserID string `form:"user_id" validate:"required,uuid4"`
	Date   string `form:"date" validate:"required,datetime=2006-01-02"`
}

type PolarGetDataParams struct {
	UserID string `form:"user_id" validate:"required,uuid4"`
	Date   string `form:"date" validate:"required,datetime=2006-01-02"`
	Key    string `form:"key" validate:"omitempty,alphanum"`
}

type PolarDataHandler struct {
	store utv.PolarData
}

// NewPolarDataHandler initializes PolarData handler
func NewPolarDataHandler(store utv.PolarData) *PolarDataHandler {
	return &PolarDataHandler{store: store}
}

// GetDatesPolar godoc
//
//	@Summary		Get available dates
//	@Description	Returns available dates for the specified user (optionally filtered by date range)
//	@Tags			UTV - Polar
//	@Accept			json
//	@Produce		json
//	@Param			user_id		query		string						true	"User ID (UUID)"
//	@Param			after_date	query		string						false	"Filter dates after this date (YYYY-MM-DD)"
//	@Param			before_date	query		string						false	"Filter dates before this date (YYYY-MM-DD)"
//	@Success		200			{object}	swagger.OuraDatesResponse	"List of available dates"
//	@Success		204			"No Content: No available dates found"
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		422			{object}	swagger.OuraInvalidDateRange
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
//	@Security		ApiKeyAuth
//	@Router			/utv/polar/dates [get]
func (h *PolarDataHandler) GetDates(w http.ResponseWriter, r *http.Request) {
	err := utils.ValidateParams(r, []string{"user_id", "after_date", "before_date"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := PolarGetDatesParams{
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

	dates, err := h.store.GetDates(context.Background(), params.UserID, &params.AfterDate, &params.BeforeDate)
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

	utils.WriteJSON(w, http.StatusOK, response)
}

// GetTypesPolar godoc
//
//	@Summary		Get available types
//	@Description	Returns available types for the specified user on the specified date
//	@Tags			UTV - Polar
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		string						true	"User ID (UUID)"
//	@Param			date	query		string						true	"Date (YYYY-MM-DD)"
//	@Success		200		{object}	swagger.PolarTypesResponse	"List of available types"
//	@Success		204		"No Content: No available types found"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Security		ApiKeyAuth
//	@Router			/utv/polar/types [get]
func (h *PolarDataHandler) GetTypes(w http.ResponseWriter, r *http.Request) {
	err := utils.ValidateParams(r, []string{"user_id", "date"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := PolarGetTypesParams{
		UserID: r.URL.Query().Get("user_id"),
		Date:   r.URL.Query().Get("date"),
	}

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	types, err := h.store.GetTypes(context.Background(), params.UserID, params.Date)
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

	utils.WriteJSON(w, http.StatusOK, response)
}

// GetDataPolar godoc
//
//	@Summary		Get available data
//	@Description	Returns data for the specified user on the specified date (optionally filtered by key)
//	@Tags			UTV - Polar
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		string						true	"User ID (UUID)"
//	@Param			date	query		string						true	"Date (YYYY-MM-DD)"
//	@Param			key		query		string						false	"Type"
//	@Success		200		{object}	swagger.PolarDataResponse	"Data"
//	@Success		204		"No Content: No data found"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Security		ApiKeyAuth
//	@Router			/utv/polar/data [get]
func (h *PolarDataHandler) GetData(w http.ResponseWriter, r *http.Request) {
	err := utils.ValidateParams(r, []string{"user_id", "date", "key"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := PolarGetDataParams{
		UserID: r.URL.Query().Get("user_id"),
		Date:   r.URL.Query().Get("date"),
		Key:    r.URL.Query().Get("key"),
	}

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	data, err := h.store.GetData(context.Background(), params.UserID, params.Date, utils.NilIfEmpty(&params.Key))
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

	utils.WriteJSON(w, http.StatusOK, response)
}
