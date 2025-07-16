package utvapi

import (
	"fmt"
	"net/http"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authz"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/DeRuina/KUHA-REST-API/internal/store/utv"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type CoachtechDataHandler struct {
	store utv.CoachtechData
	cache *cache.Storage
}

func NewCoachtechDataHandler(store utv.CoachtechData, cache *cache.Storage) *CoachtechDataHandler {
	return &CoachtechDataHandler{store: store, cache: cache}
}

type CoachtechStatusParams struct {
	UserID string `form:"user_id" validate:"required,uuid4"`
}

// GetCoachtechStatus godoc
//
//	@Summary		Check Coachtech data availability
//	@Description	Returns whether Coachtech data exists for a given user
//	@Tags			UTV - Coachtech
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		string	true	"User ID (UUID)"
//	@Success		200		{object}	swagger.CoachtechStatusResponse
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Security		BearerAuth
//	@Router			/utv/coachtech/status [get]
func (h *CoachtechDataHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	err := utils.ValidateParams(r, []string{"user_id"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := CoachtechStatusParams{
		UserID: r.URL.Query().Get("user_id"),
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

	status, err := h.store.GetStatus(r.Context(), userID)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]bool{"data": status})
}

type CoachtechDataParams struct {
	UserID     string `form:"user_id" validate:"required,uuid4"`
	AfterDate  string `form:"after_date" validate:"omitempty,datetime=2006-01-02"`
	BeforeDate string `form:"before_date" validate:"omitempty,datetime=2006-01-02"`
}

// GetCoachtechData godoc
//
//	@Summary		Get Coachtech data
//	@Description	Returns Coachtech data for a user
//	@Tags			UTV - Coachtech
//	@Accept			json
//	@Produce		json
//	@Param			user_id		query	string	true	"User ID (UUID)"
//	@Param			after_date	query	string	false	"Filter data after this date (YYYY-MM-DD)"
//	@Param			before_date	query	string	false	"Filter data before this date (YYYY-MM-DD)"
//	@Success		200
//	@Success		204	"No Content"
//	@Failure		400	{object}	swagger.ValidationErrorResponse
//	@Failure		403	{object}	swagger.ForbiddenResponse
//	@Failure		500	{object}	swagger.InternalServerErrorResponse
//	@Security		BearerAuth
//	@Router			/utv/coachtech/data [get]
func (h *CoachtechDataHandler) GetData(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	err := utils.ValidateParams(r, []string{"user_id", "after_date", "before_date"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := CoachtechDataParams{
		UserID:     r.URL.Query().Get("user_id"),
		AfterDate:  r.URL.Query().Get("after_date"),
		BeforeDate: r.URL.Query().Get("before_date"),
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

	// Validate date range
	if after != nil && before != nil && after.After(*before) {
		utils.UnprocessableEntityResponse(w, r, utils.ErrInvalidDateRange)
		return
	}

	data, err := h.store.GetData(r.Context(), userID, after, before)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	if len(data) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	utils.WriteJSON(w, http.StatusOK, data)
}
