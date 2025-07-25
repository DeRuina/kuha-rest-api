package tietoevryapi

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/DeRuina/KUHA-REST-API/docs/swagger"
	"github.com/DeRuina/KUHA-REST-API/internal/auth/authz"
	tietoevrysqlc "github.com/DeRuina/KUHA-REST-API/internal/db/tietoevry"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/DeRuina/KUHA-REST-API/internal/store/tietoevry"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

// handler struct
type TietoevryUserHandler struct {
	store tietoevry.Users
	cache *cache.Storage
}

func NewTietoevryUserHandler(store tietoevry.Users, cache *cache.Storage) *TietoevryUserHandler {
	return &TietoevryUserHandler{store: store, cache: cache}
}

// request model
type TietoevryUserUpsertInput struct {
	ID                        string   `json:"id" validate:"required,uuid4"`
	SporttiID                 int32    `json:"sportti_id" validate:"required"`
	ProfileGender             *string  `json:"profile_gender"`
	ProfileBirthdate          *string  `json:"profile_birthdate"` // ISO 8601 string
	ProfileWeight             *float64 `json:"profile_weight"`
	ProfileHeight             *float64 `json:"profile_height"`
	ProfileRestingHeartRate   *int32   `json:"profile_resting_heart_rate"`
	ProfileMaximumHeartRate   *int32   `json:"profile_maximum_heart_rate"`
	ProfileAerobicThreshold   *int32   `json:"profile_aerobic_threshold"`
	ProfileAnaerobicThreshold *int32   `json:"profile_anaerobic_threshold"`
	ProfileVo2max             *int32   `json:"profile_vo2max"`
}

type TietoevryUserDeleteParams struct {
	ID string `form:"id" validate:"required,uuid4"`
}

// UpsertUser godoc
//
//	@Summary		Upsert user
//	@Description	Upserts a user with the provided data
//	@Tags			Tietoevry - User
//	@Accept			json
//	@Produce		json
//	@Param			user	body	swagger.TietoevryUserUpsertInput	true	"User data"
//	@Success		201		"created"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Security		BearerAuth
//	@Router			/tietoevry/users [post]
func (h *TietoevryUserHandler) UpsertUser(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var input TietoevryUserUpsertInput
	if err := utils.ReadJSON(w, r, &input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	if err := utils.GetValidator().Struct(input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	ID, err := utils.ParseUUID(input.ID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	parsedBirthdate, err := utils.ParseDatePtr(input.ProfileBirthdate)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	birthdate := utils.NullTimeIfEmpty(parsedBirthdate)

	arg := tietoevrysqlc.UpsertUserParams{
		ID:                        ID,
		SporttiID:                 input.SporttiID,
		ProfileGender:             utils.NullStringPtr(input.ProfileGender),
		ProfileBirthdate:          birthdate,
		ProfileWeight:             utils.NullFloat64Ptr(input.ProfileWeight),
		ProfileHeight:             utils.NullFloat64Ptr(input.ProfileHeight),
		ProfileRestingHeartRate:   utils.NullInt32Ptr(input.ProfileRestingHeartRate),
		ProfileMaximumHeartRate:   utils.NullInt32Ptr(input.ProfileMaximumHeartRate),
		ProfileAerobicThreshold:   utils.NullInt32Ptr(input.ProfileAerobicThreshold),
		ProfileAnaerobicThreshold: utils.NullInt32Ptr(input.ProfileAnaerobicThreshold),
		ProfileVo2max:             utils.NullInt32Ptr(input.ProfileVo2max),
	}

	if err := h.store.UpsertUser(r.Context(), arg); err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// DeleteUser godoc
//
//	@Summary		Delete a user
//	@Description	Removes a user by ID
//	@Tags			Tietoevry - User
//	@Accept			json
//	@Produce		json
//	@Param			id	query	string	true	"User ID (UUID)"
//	@Success		200 "User deleted successfully"
//	@Success		204 "No content, user not found"
//	@Failure		400	{object}	swagger.ValidationErrorResponse
//	@Failure		403	{object}	swagger.ForbiddenResponse
//	@Failure		500	{object}	swagger.InternalServerErrorResponse
//	@Security		BearerAuth
//	@Router			/tietoevry/users [delete]
func (h *TietoevryUserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{"id"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := TietoevryUserDeleteParams{
		ID: r.URL.Query().Get("id"),
	}

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	userID, err := utils.ParseUUID(params.ID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	rows, err := h.store.DeleteUserWithLogging(r.Context(), userID)
	if err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	if rows == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusOK)

}

// GetUser godoc
//
//	@Summary		Get user by ID
//	@Description	Retrieve a single user by UUID
//	@Tags			Tietoevry - User
//	@Accept			json
//	@Produce		json
//	@Param			id	query		string	true	"User ID (UUID)"
//	@Success		200	{object}	swagger.TietoevryUserResponse
//	@Failure		400	{object}	swagger.ValidationErrorResponse
//	@Failure		403	{object}	swagger.ForbiddenResponse
//	@Failure		404	{object}	swagger.NotFoundResponse
//	@Failure		500	{object}	swagger.InternalServerErrorResponse
//	@Security		BearerAuth
//	@Router			/tietoevry/users [get]
func (h *TietoevryUserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{"id"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := TietoevryUserDeleteParams{
		ID: r.URL.Query().Get("id"),
	}

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	userID, err := utils.ParseUUID(params.ID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	user, err := h.store.GetUser(r.Context(), userID)
	if err == sql.ErrNoRows {
		utils.NotFoundResponse(w, r, err)
		return
	}
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	resp := swagger.TietoevryUserUpsertInput{
		ID:                        user.ID.String(),
		SporttiID:                 user.SporttiID,
		ProfileGender:             utils.StringPtrOrNil(user.ProfileGender),
		ProfileBirthdate:          utils.FormatDatePtr(user.ProfileBirthdate),
		ProfileWeight:             utils.Float64PtrOrNil(user.ProfileWeight),
		ProfileHeight:             utils.Float64PtrOrNil(user.ProfileHeight),
		ProfileRestingHeartRate:   utils.Int32PtrOrNil(user.ProfileRestingHeartRate),
		ProfileMaximumHeartRate:   utils.Int32PtrOrNil(user.ProfileMaximumHeartRate),
		ProfileAerobicThreshold:   utils.Int32PtrOrNil(user.ProfileAerobicThreshold),
		ProfileAnaerobicThreshold: utils.Int32PtrOrNil(user.ProfileAnaerobicThreshold),
		ProfileVo2max:             utils.Int32PtrOrNil(user.ProfileVo2max),
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"user": resp,
	})

}

// GetDeletedUsers godoc
//
//	@Summary		List deleted users
//	@Description	Returns a list of deleted users with timestamps
//	@Tags			Tietoevry - User
//	@Produce		json
//	@Success		200	{object}	swagger.TietoevryDeletedUsersResponse
//	@Failure		403	{object}	swagger.ForbiddenResponse
//	@Failure		500	{object}	swagger.InternalServerErrorResponse
//	@Security		BearerAuth
//	@Router			/tietoevry/deleted-users [get]
func (h *TietoevryUserHandler) GetDeletedUsers(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	data, err := h.store.GetDeletedUsers(r.Context())
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	var output []swagger.TietoevryDeletedUser
	for _, d := range data {
		output = append(output, swagger.TietoevryDeletedUser{
			UserID:    d.UserID.String(),
			SporttiID: d.SporttiID,
			DeletedAt: d.DeletedAt.Format(time.RFC3339),
		})
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"deleted_users": output,
	})
}
