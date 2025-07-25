package tietoevryapi

import (
	"fmt"
	"net/http"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authz"
	tietoevrysqlc "github.com/DeRuina/KUHA-REST-API/internal/db/tietoevry"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/DeRuina/KUHA-REST-API/internal/store/tietoevry"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
	"github.com/google/uuid"
)

type TietoevryActivityZoneHandler struct {
	store tietoevry.ActivityZones
	cache *cache.Storage
}

func NewTietoevryActivityZoneHandler(store tietoevry.ActivityZones, cache *cache.Storage) *TietoevryActivityZoneHandler {
	return &TietoevryActivityZoneHandler{store: store, cache: cache}
}

type TietoevryActivityZoneInput struct {
	UserID         string   `json:"user_id" validate:"required,uuid4"`
	Date           string   `json:"date" validate:"required"`
	CreatedAt      string   `json:"created_at" validate:"required"`
	UpdatedAt      string   `json:"updated_at" validate:"required"`
	SecondsInZone0 *float64 `json:"seconds_in_zone_0"`
	SecondsInZone1 *float64 `json:"seconds_in_zone_1"`
	SecondsInZone2 *float64 `json:"seconds_in_zone_2"`
	SecondsInZone3 *float64 `json:"seconds_in_zone_3"`
	SecondsInZone4 *float64 `json:"seconds_in_zone_4"`
	SecondsInZone5 *float64 `json:"seconds_in_zone_5"`
	Source         string   `json:"source" validate:"required"`
	RawData        *string  `json:"raw_data"`
}

type TietoevryActivityZonesBulkInput struct {
	ActivityZones []TietoevryActivityZoneInput `json:"activity_zones" validate:"required,dive"`
}

// InsertActivityZonesBulk godoc
//
//	@Summary		Insert activity zones (bulk)
//	@Description	Insert multiple activity zone summaries for user (idempotent)
//	@Tags			Tietoevry - ActivityZones
//	@Accept			json
//	@Produce		json
//	@Param			activity_zones	body	swagger.TietoevryActivityZonesBulkInput	true	"Activity zone summaries"
//	@Success		201				"Activity zones processed successfully (idempotent operation)"
//	@Failure		400				{object}	swagger.ValidationErrorResponse
//	@Failure		403				{object}	swagger.ForbiddenResponse
//	@Failure		409				{object}	swagger.ConflictResponse
//	@Failure		500				{object}	swagger.InternalServerErrorResponse
//	@Security		BearerAuth
//	@Router			/tietoevry/activity-zones [post]
func (h *TietoevryActivityZoneHandler) InsertActivityZonesBulk(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var input TietoevryActivityZonesBulkInput
	if err := utils.ReadJSON(w, r, &input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	if err := utils.GetValidator().Struct(input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	// PreValidation
	userIDs := make([]uuid.UUID, len(input.ActivityZones))
	for i, a := range input.ActivityZones {
		uid, err := utils.ParseUUID(a.UserID)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}
		userIDs[i] = uid
	}

	if err := h.store.ValidateUsersExist(r.Context(), userIDs); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	activityZones := make([]tietoevrysqlc.InsertActivityZoneParams, len(input.ActivityZones))
	for i, activityZone := range input.ActivityZones {
		// Parse and convert values
		userID, err := utils.ParseUUID(activityZone.UserID)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		date, err := utils.ParseDate(activityZone.Date)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		createdAt, err := utils.ParseTimestamp(activityZone.CreatedAt)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		updatedAt, err := utils.ParseTimestamp(activityZone.UpdatedAt)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		rawData := utils.ParseRawJSON(activityZone.RawData)

		activityZones[i] = tietoevrysqlc.InsertActivityZoneParams{
			UserID:         userID,
			Date:           date,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
			SecondsInZone0: utils.NullFloat64Ptr(activityZone.SecondsInZone0),
			SecondsInZone1: utils.NullFloat64Ptr(activityZone.SecondsInZone1),
			SecondsInZone2: utils.NullFloat64Ptr(activityZone.SecondsInZone2),
			SecondsInZone3: utils.NullFloat64Ptr(activityZone.SecondsInZone3),
			SecondsInZone4: utils.NullFloat64Ptr(activityZone.SecondsInZone4),
			SecondsInZone5: utils.NullFloat64Ptr(activityZone.SecondsInZone5),
			Source:         activityZone.Source,
			RawData:        rawData,
		}
	}

	if err := h.store.InsertActivityZonesBulk(r.Context(), activityZones); err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
