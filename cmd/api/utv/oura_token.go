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

type OuraTokenHandler struct {
	store utv.OuraToken
	cache *cache.Storage
}

func NewOuraTokenHandler(store utv.OuraToken, cache *cache.Storage) *OuraTokenHandler {
	return &OuraTokenHandler{store: store, cache: cache}
}

type OuraStatusParams struct {
	UserID string `form:"user_id" validate:"required,uuid4"`
}

// GetStatus godoc
//
//	@Summary		Check Oura connection & data status
//	@Description	Returns whether a user has connected their Oura account and whether any Oura data exists
//	@Tags			UTV - Oura
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		string	true	"User ID (UUID)"
//	@Success		200		{object}	swagger.OuraStatusResponse
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Security		BearerAuth
//	@Router			/utv/oura/status [get]
func (h *OuraTokenHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	err := utils.ValidateParams(r, []string{"user_id"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := OuraStatusParams{
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

type OuraTokenInput struct {
	UserID  string          `json:"user_id" validate:"required,uuid4"`
	Details json.RawMessage `json:"details" validate:"required"`
}

// UpsertToken godoc
//
//	@Summary		Save or update Oura token
//	@Description	Saves or updates the Oura token for a user
//	@Tags			UTV - Oura
//	@Accept			json
//	@Produce		json
//	@Param			body	body	swagger.OuraTokenInput	true	"Oura token input"
//	@Success		201		"Created"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Security		BearerAuth
//	@Router			/utv/oura/token [post]
func (h *OuraTokenHandler) UpsertToken(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var input OuraTokenInput
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

type GetTokenByOuraIDParams struct {
	OuraID string `form:"oura-id" validate:"required"`
}

// GetTokenByOuraID godoc
//
//	@Summary		Get Oura token by Oura ID
//	@Description	Retrieves the user ID and token data associated with a specific Oura ID
//	@Tags			UTV - Oura
//	@Accept			json
//	@Produce		json
//	@Param			oura-id	query		string	true	"Oura ID"
//	@Success		200		{object}	swagger.OuraTokenByIDResponse
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Security		BearerAuth
//	@Router			/utv/oura/token-by-id [get]
func (h *OuraTokenHandler) GetTokenByOuraID(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	err := utils.ValidateParams(r, []string{"oura-id"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := GetTokenByOuraIDParams{
		OuraID: r.URL.Query().Get("oura-id"),
	}
	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	userID, data, err := h.store.GetTokenByOuraID(r.Context(), params.OuraID)
	if err != nil {
		utils.WriteJSON(w, http.StatusOK, map[string]any{})
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"user_id": userID,
		"data":    data,
	})
}
