package utvapi

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authz"
	"github.com/DeRuina/KUHA-REST-API/internal/logger"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/DeRuina/KUHA-REST-API/internal/store/utv"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
	"github.com/google/uuid"
)

type UserDataHandler struct {
	store utv.UserData
	cache *cache.Storage
}

func NewUserDataHandler(store utv.UserData, cache *cache.Storage) *UserDataHandler {
	return &UserDataHandler{store: store, cache: cache}
}

type UserIDParam struct {
	UserID string `form:"user_id" validate:"required,uuid4"`
}

type SportIDParam struct {
	SportID string `form:"sport_id" validate:"required,numeric"`
}

type UserDataInput struct {
	Data json.RawMessage `json:"data" validate:"required"`
}

// GetUserData godoc
//
//	@Summary		Get user
//	@Description	Returns JSON user data for a given UTV user ID
//	@Tags			UTV - User
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		string	true	"User ID (UUID)"
//	@Success		200		{object}	swagger.UserDataResponse
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		404		{object}	swagger.NotFoundResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/utv/user [get]
func (h *UserDataHandler) GetUserData(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{"user_id"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := UserIDParam{
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

	data, err := h.store.GetUserData(r.Context(), userID)
	if err == sql.ErrNoRows {
		utils.NotFoundResponse(w, r, err)
		return
	}
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"data": data,
	})
}

type UserDataUpsertInput struct {
	Data json.RawMessage `json:"data" validate:"required"`
}

// UpsertUserData godoc
//
//	@Summary		Save or update user
//	@Description	Upserts the user data JSON blob for a specific user
//	@Tags			UTV - User
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query	string					true	"User ID (UUID)"
//	@Param			body	body	swagger.UserDataInput	true	"User data input"
//	@Success		201		"Created"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/utv/user [post]
func (h *UserDataHandler) UpsertUserData(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}
	if err := utils.ValidateParams(r, []string{"user_id"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := UserIDParam{
		UserID: r.URL.Query().Get("user_id"),
	}
	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	userID, err := uuid.Parse(params.UserID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	var input UserDataUpsertInput
	if err := utils.ReadJSON(w, r, &input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	if err := utils.GetValidator().Struct(input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	if err := h.store.UpsertUserData(r.Context(), userID, input.Data); err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// DeleteUserData godoc
//
//	@Summary		Delete a user
//	@Description	Removes a user
//	@Tags			UTV - User
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query	string	true	"User ID (UUID)"
//	@Success		200
//	@Failure		400	{object}	swagger.ValidationErrorResponse
//	@Failure		401	{object}	swagger.UnauthorizedResponse
//	@Failure		403	{object}	swagger.ForbiddenResponse
//	@Failure		500	{object}	swagger.InternalServerErrorResponse
//	@Failure		503	{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/utv/user [delete]
func (h *UserDataHandler) DeleteUserData(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{"user_id"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := UserIDParam{
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

	if err := h.store.DeleteUserData(r.Context(), userID); err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetUserIDBySportID godoc
//
//	@Summary		Get UTV user ID by sport ID
//	@Description	Find user_id from sport_id
//	@Tags			UTV - User
//	@Accept			json
//	@Produce		json
//	@Param			sport_id	query		string	true	"Sport ID"
//	@Success		200			{object}	swagger.UserIDResponse
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		401			{object}	swagger.UnauthorizedResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		404			{object}	swagger.NotFoundResponse
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
//	@Failure		503			{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/utv/user-id-by-sport-id [get]
func (h *UserDataHandler) GetUserIDBySportID(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{"sport_id"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := SportIDParam{
		SportID: r.URL.Query().Get("sport_id"),
	}
	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	logger.Logger.Warnw("Looking up sport_id", "sport_id", params.SportID)

	userID, err := h.store.GetUserIDBySportID(r.Context(), params.SportID)
	if err == sql.ErrNoRows {
		utils.NotFoundResponse(w, r, err)
		return
	}
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{
		"user_id": userID.String(),
	})
}
