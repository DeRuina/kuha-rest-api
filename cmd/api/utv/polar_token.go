package utvapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authz"
	"github.com/DeRuina/KUHA-REST-API/internal/store/utv"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type PolarTokenHandler struct {
	store utv.PolarToken
}

func NewPolarTokenHandler(store utv.PolarToken) *PolarTokenHandler {
	return &PolarTokenHandler{store: store}
}

type GetStatusParams struct {
	UserID string `form:"user_id" validate:"required,uuid4"`
}

// GetStatus godoc
//
//	@Summary		Check Polar connection & data status
//	@Description	Returns whether a user has connected their Polar account and whether any Polar data exists
//	@Tags			UTV - Polar
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		string	true	"User ID (UUID)"
//	@Success		200		{object}	swagger.PolarStatusResponse
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Security		BearerAuth
//	@Router			/utv/polar/status [get]
func (h *PolarTokenHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("unauthorized"))
		return
	}

	err := utils.ValidateParams(r, []string{"user_id"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := GetStatusParams{
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

	connected, dataExists, err := h.store.GetStatus(r.Context(), userID)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]bool{
		"connected": connected,
		"data":      dataExists,
	})
}

type PolarTokenInput struct {
	UserID  string          `json:"user_id" validate:"required,uuid4"`
	Details json.RawMessage `json:"details" validate:"required"`
}

// UpsertToken godoc
//
//	@Summary		Save or update Polar token
//	@Description	Upserts the Polar token details for a specific user
//	@Tags			UTV - Polar
//	@Accept			json
//	@Produce		json
//	@Param			body	body	swagger.PolarTokenInput	true	"Polar token input"
//	@Success		201		"Created"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Security		BearerAuth
//	@Router			/utv/polar/token [post]
func (h *PolarTokenHandler) UpsertToken(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("unauthorized"))
		return
	}

	var input PolarTokenInput
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

	if err := h.store.UpsertToken(r.Context(), userID, input.Details); err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type GetTokenByPolarIDParams struct {
	PolarID string `form:"polar-id" validate:"required"`
}

// GetTokenByPolarID godoc
//
//	@Summary		Get user token by Polar ID
//	@Description	Returns user_id and token data associated with a given Polar x_user_id
//	@Tags			UTV - Polar
//	@Accept			json
//	@Produce		json
//	@Param			polar-id	query		string	true	"Polar x_user_id"
//	@Success		200			{object}	swagger.PolarTokenByIDResponse
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Security		BearerAuth
//	@Router			/utv/polar/token-by-id [get]
func (h *PolarTokenHandler) GetTokenByPolarID(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("unauthorized"))
		return
	}

	err := utils.ValidateParams(r, []string{"polar-id"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := GetTokenByPolarIDParams{
		PolarID: r.URL.Query().Get("polar-id"),
	}
	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	userID, data, err := h.store.GetTokenByPolarID(r.Context(), params.PolarID)
	if err != nil {
		utils.WriteJSON(w, http.StatusOK, map[string]any{})
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"user_id": userID,
		"data":    data,
	})
}
