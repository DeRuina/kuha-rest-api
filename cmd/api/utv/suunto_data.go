package utvapi

import (
	"context"
	"net/http"

	"github.com/DeRuina/KUHA-REST-API/internal/store/utv"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

// Validation structs
type SuuntoGetDatesParams struct {
	UserID     string `form:"user_id" validate:"required,uuid4"`
	AfterDate  string `form:"after_date" validate:"omitempty,datetime=2006-01-02"`
	BeforeDate string `form:"before_date" validate:"omitempty,datetime=2006-01-02"`
}

type SuuntoGetTypesParams struct {
	UserID string `form:"user_id" validate:"required,uuid4"`
	Date   string `form:"date" validate:"required,datetime=2006-01-02"`
}

type SuuntoGetDataParams struct {
	UserID string `form:"user_id" validate:"required,uuid4"`
	Date   string `form:"date" validate:"required,datetime=2006-01-02"`
	Key    string `form:"key" validate:"omitempty,alphanum"`
}

type SuuntoDataHandler struct {
	store utv.SuuntoData
}

// NewSuuntoDataHandler initializes SuuntoData handler
func NewSuuntoDataHandler(store utv.SuuntoData) *SuuntoDataHandler {
	return &SuuntoDataHandler{store: store}
}

// GetDatesSuunto godoc
//
//	@Summary		Get available dates (Suunto)
//	@Description	Returns available dates for the specified user (optionally filtered by date range)
//	@Tags			UTV
//	@Accept			json
//	@Produce		json
//	@Param			user_id		query	string	true	"User ID (UUID)"
//	@Param			after_date	query	string	false	"Filter dates after this date (YYYY-MM-DD)"
//	@Param			before_date	query	string	false	"Filter dates before this date (YYYY-MM-DD)"
//	@Success		200			{array}	string	"List of available dates"
//	@Success		204			"No Content: No available dates found"
//	@Failure		400			{object}	error
//	@Failure		422			{object}	error
//	@Failure		500			{object}	error
//	@Security		ApiKeyAuth
//	@Router			/utv/suunto/dates [get]
func (h *SuuntoDataHandler) GetDates(w http.ResponseWriter, r *http.Request) {
	err := utils.ValidateParams(r, []string{"user_id", "after_date", "before_date"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := SuuntoGetDatesParams{
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

// GetTypesSuunto godoc
//
//	@Summary		Get available types (Suunto)
//	@Description	Returns available types for the specified user on the specified date
//	@Tags			UTV
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query	string	true	"User ID (UUID)"
//	@Param			date	query	string	true	"Date (YYYY-MM-DD)"
//	@Success		200		{array}	string	"List of available types"
//	@Success		204		"No Content: No available types found"
//	@Failure		400		{object}	error
//	@Failure		422		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/utv/suunto/types [get]
func (h *SuuntoDataHandler) GetTypes(w http.ResponseWriter, r *http.Request) {
	err := utils.ValidateParams(r, []string{"user_id", "date"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := SuuntoGetTypesParams{
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

// GetDataSuunto godoc
//
//	@Summary		Get available data (Suunto)
//	@Description	Returns data for the specified user on the specified date (optionally filtered by key)
//	@Tags			UTV
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		string					true	"User ID (UUID)"
//	@Param			date	query		string					true	"Date (YYYY-MM-DD)"
//	@Param			key		query		string					false	"Type"
//	@Success		200		{object}	map[string]interface{}	"Data"
//	@Success		204		"No Content: No data found"
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Security		ApiKeyAuth
//	@Router			/utv/suunto/data [get]
func (h *SuuntoDataHandler) GetData(w http.ResponseWriter, r *http.Request) {
	err := utils.ValidateParams(r, []string{"user_id", "date", "key"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := SuuntoGetDataParams{
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
