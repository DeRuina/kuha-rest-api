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

type SuuntoTokenHandler struct {
	store utv.SuuntoToken
	cache *cache.Storage
}

func NewSuuntoTokenHandler(store utv.SuuntoToken, cache *cache.Storage) *SuuntoTokenHandler {
	return &SuuntoTokenHandler{store: store, cache: cache}
}

type SuuntoStatusParams struct {
	UserID string `form:"user_id" validate:"required,uuid4"`
}

// GetStatus godoc
//
//	@Summary		Check Suunto connection & data status
//	@Description	Returns whether a user has connected their Suunto account and whether any Suunto data exists
//	@Tags			UTV - Suunto
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		string	true	"User ID (UUID)"
//	@Success		200		{object}	swagger.SuuntoStatusResponse
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/utv/suunto/status [get]
func (h *SuuntoTokenHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	err := utils.ValidateParams(r, []string{"user_id"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := SuuntoStatusParams{
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

type SuuntoTokenInput struct {
	UserID  string          `json:"user_id" validate:"required,uuid4"`
	Details json.RawMessage `json:"details" validate:"required"`
}

// UpsertToken godoc
//
//	@Summary		Save or update Suunto token
//	@Description	Upserts the Suunto token details for a specific user
//	@Tags			UTV - Suunto
//	@Accept			json
//	@Produce		json
//	@Param			body	body	swagger.SuuntoTokenInput	true	"Suunto token input"
//	@Success		201		"Created"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/utv/suunto/token [post]
func (h *SuuntoTokenHandler) UpsertToken(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var input SuuntoTokenInput
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

type GetTokenByUsernameParams struct {
	Username string `form:"username" validate:"required"`
}

// GetTokenByUsername godoc
//
//	@Summary		Get Suunto token by username
//	@Description	Returns user_id and token data associated with a given Suunto username
//	@Tags			UTV - Suunto
//	@Accept			json
//	@Produce		json
//	@Param			username	query		string	true	"Suunto username"
//	@Success		200			{object}	swagger.SuuntoTokenByUsernameResponse
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		401			{object}	swagger.UnauthorizedResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
//	@Failure		503			{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/utv/suunto/token-by-username [get]
func (h *SuuntoTokenHandler) GetTokenByUsername(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	err := utils.ValidateParams(r, []string{"username"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := GetTokenByUsernameParams{
		Username: r.URL.Query().Get("username"),
	}
	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	userID, data, err := h.store.GetTokenByUsername(r.Context(), params.Username)
	if err != nil {
		utils.WriteJSON(w, http.StatusOK, map[string]any{})
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"user_id": userID,
		"data":    data,
	})
}
