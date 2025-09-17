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

type ArchinisisTokenHandler struct {
	store utv.ArchinisisToken
	cache *cache.Storage
}

func NewArchinisisTokenHandler(store utv.ArchinisisToken, cache *cache.Storage) *ArchinisisTokenHandler {
	return &ArchinisisTokenHandler{store: store, cache: cache}
}

type ArchinisisStatusParams struct {
	UserID string `form:"user_id" validate:"required,uuid4"`
}

// GetStatus godoc
//
//	@Summary		Check Archinisis connection
//	@Description	Returns whether a user has connected their Archinisis account
//	@Tags			UTV - Archinisis
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		string	true	"User ID (UUID)"
//	@Success		200		{object}	swagger.ArchinisisStatusResponse
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/utv/archinisis/status [get]
func (h *ArchinisisTokenHandler) GetStatus(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	err := utils.ValidateParams(r, []string{"user_id"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := ArchinisisStatusParams{
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

	connected, err := h.store.GetStatus(r.Context(), userID)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]bool{
		"connected": connected,
	})
}

type ArchinisisTokenInput struct {
	UserID  string          `json:"user_id" validate:"required,uuid4"`
	Details json.RawMessage `json:"details" validate:"required"`
}

// UpsertToken godoc
//
//	@Summary		Upsert Archinisis token
//	@Description	Upserts a Archinisis token for a user
//	@Tags			UTV - Archinisis
//	@Accept			json
//	@Produce		json
//	@Param			body	body	swagger.ArchinisisTokenInput	true	"Archinisis token input"
//	@Success		201		"Created"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/utv/archinisis/token [post]
func (h *ArchinisisTokenHandler) UpsertToken(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var input ArchinisisTokenInput
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

// GetSportIDs godoc
//
//	@Summary		List Archinisis sport IDs
//	@Description	Returns distinct sport_id values from Archinisis tokens
//	@Tags			UTV - Archinisis
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	swagger.SportIDsResponse
//	@Failure		400	{object}	swagger.ValidationErrorResponse
//	@Failure		401	{object}	swagger.UnauthorizedResponse
//	@Failure		403	{object}	swagger.ForbiddenResponse
//	@Failure		500	{object}	swagger.InternalServerErrorResponse
//	@Failure		503	{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/utv/archinisis/sport_ids [get]
func (h *ArchinisisTokenHandler) GetSportIDs(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}
	ids, err := h.store.GetSportIDs(r.Context())
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string][]string{"sportti_ids": ids})
}
