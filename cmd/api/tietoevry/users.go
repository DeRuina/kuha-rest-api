package tietoevryapi

import (
	"fmt"
	"net/http"

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
		utils.InternalServerError(w, r, err)
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
//	@Success		200
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

	if err := h.store.DeleteUser(r.Context(), userID); err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
