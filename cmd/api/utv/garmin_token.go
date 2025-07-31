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

type GarminTokenHandler struct {
	store utv.GarminToken
	cache *cache.Storage
}

func NewGarminTokenHandler(store utv.GarminToken, cache *cache.Storage) *GarminTokenHandler {
	return &GarminTokenHandler{store: store, cache: cache}
}

type GarminStatusParams struct {
	UserID string `form:"user_id" validate:"required,uuid4"`
}

// GetStatus godoc
//
//	@Summary		Check Garmin connection & data status
//	@Description	Returns whether a user has connected their Garmin account and whether any Garmin data exists
//	@Tags			UTV - Garmin
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		string	true	"User ID (UUID)"
//	@Success		200		{object}	swagger.GarminStatusResponse
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/utv/garmin/status [get]
func (h *GarminTokenHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	err := utils.ValidateParams(r, []string{"user_id"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := GarminStatusParams{
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

type GarminTokenInput struct {
	UserID  string          `json:"user_id" validate:"required,uuid4"`
	Details json.RawMessage `json:"details" validate:"required"`
}

// UpsertToken godoc
//
//	@Summary		Save or update Garmin token
//	@Description	Saves or updates the Garmin token for a user
//	@Tags			UTV - Garmin
//	@Accept			json
//	@Produce		json
//	@Param			body	body	swagger.GarminTokenInput	true	"Garmin token input"
//	@Success		201		"Created"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/utv/garmin/token [post]
func (h *GarminTokenHandler) UpsertToken(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var input GarminTokenInput
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
		utils.HandleDatabaseError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type GarminTokenByAccessTokenParams struct {
	Token string `form:"token" validate:"required"`
}

// GetUserIDByToken godoc
//
//	@Summary		Get user_id by Garmin access token
//	@Description	Returns the user ID associated with a given Garmin access token
//	@Tags			UTV - Garmin
//	@Accept			json
//	@Produce		json
//	@Param			token	query		string	true	"Garmin access token"
//	@Success		200		{object}	swagger.GarminUserIDResponse
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/utv/garmin/user-id-by-token [get]
func (h *GarminTokenHandler) GetUserIDByToken(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	err := utils.ValidateParams(r, []string{"token"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := GarminTokenByAccessTokenParams{
		Token: r.URL.Query().Get("token"),
	}
	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	userID, err := h.store.GetUserIDByToken(r.Context(), params.Token)
	if err != nil {
		utils.WriteJSON(w, http.StatusOK, map[string]any{})
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"user_id": userID,
	})
}

// TokenExists godoc
//
//	@Summary		Check if Garmin token exists
//	@Description	Checks whether a given Garmin access token is stored in the system
//	@Tags			UTV - Garmin
//	@Accept			json
//	@Produce		json
//	@Param			token	query		string	true	"Garmin access token"
//	@Success		200		{object}	swagger.GarminTokenExistsResponse
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/utv/garmin/token-exists [get]
func (h *GarminTokenHandler) TokenExists(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	err := utils.ValidateParams(r, []string{"token"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := GarminTokenByAccessTokenParams{
		Token: r.URL.Query().Get("token"),
	}
	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	exists, err := h.store.TokenExists(r.Context(), params.Token)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]bool{
		"exists": exists,
	})
}
