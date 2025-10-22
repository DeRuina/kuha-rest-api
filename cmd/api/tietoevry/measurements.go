package tietoevryapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/DeRuina/KUHA-REST-API/docs/swagger"
	"github.com/DeRuina/KUHA-REST-API/internal/auth/authz"
	tietoevrysqlc "github.com/DeRuina/KUHA-REST-API/internal/db/tietoevry"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/DeRuina/KUHA-REST-API/internal/store/tietoevry"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
	"github.com/google/uuid"
)

// Handler struct
type TietoevryMeasurementHandler struct {
	store tietoevry.Measurements
	cache *cache.Storage
}

func NewTietoevryMeasurementHandler(store tietoevry.Measurements, cache *cache.Storage) *TietoevryMeasurementHandler {
	return &TietoevryMeasurementHandler{store: store, cache: cache}
}

// Input struct for validation
type TietoevryMeasurementInput struct {
	ID             string   `json:"id" validate:"required,uuid4"`
	CreatedAt      string   `json:"created_at" validate:"required"`
	UpdatedAt      string   `json:"updated_at" validate:"required"`
	UserID         string   `json:"user_id" validate:"required,uuid4"`
	Date           string   `json:"date" validate:"required"`
	Name           string   `json:"name" validate:"required"`
	NameType       string   `json:"name_type" validate:"required"`
	Source         string   `json:"source" validate:"required"`
	Value          string   `json:"value" validate:"required"`
	ValueNumeric   *float64 `json:"value_numeric"`
	Comment        *string  `json:"comment"`
	RawID          *string  `json:"raw_id"`
	RawData        *string  `json:"raw_data"`
	AdditionalInfo *string  `json:"additional_info"`
}

type TietoevryMeasurementsBulkInput struct {
	Measurements []TietoevryMeasurementInput `json:"measurements" validate:"required,dive"`
}

// InsertMeasurements godoc
//
//	@Summary		Insert measurements (bulk)
//	@Description	Insert multiple measurements for user (idempotent)
//	@Tags			Tietoevry - Measurements
//	@Accept			json
//	@Produce		json
//	@Param			measurements	body	swagger.TietoevryMeasurementsBulkInput	true	"Measurement data"
//	@Success		201				"Measurements processed successfully"
//	@Failure		400				{object}	swagger.ValidationErrorResponse
//	@Failure		401				{object}	swagger.UnauthorizedResponse
//	@Failure		403				{object}	swagger.ForbiddenResponse
//	@Failure		409				{object}	swagger.ConflictResponse
//	@Failure		500				{object}	swagger.InternalServerErrorResponse
//	@Failure		503				{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/tietoevry/measurements [post]
func (h *TietoevryMeasurementHandler) InsertMeasurementsBulk(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var input TietoevryMeasurementsBulkInput
	if err := utils.ReadJSON(w, r, &input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	if err := utils.GetValidator().Struct(input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	// PreValidation
	userIDs := make([]uuid.UUID, len(input.Measurements))
	for i, m := range input.Measurements {
		userID, err := utils.ParseUUID(m.UserID)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}
		userIDs[i] = userID
	}

	if err := h.store.ValidateUsersExist(r.Context(), userIDs); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := make([]tietoevrysqlc.InsertMeasurementParams, len(input.Measurements))
	for i, m := range input.Measurements {
		id, err := utils.ParseUUID(m.ID)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}
		userID, err := utils.ParseUUID(m.UserID)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}
		createdAt, err := utils.ParseTimestamp(m.CreatedAt)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}
		updatedAt, err := utils.ParseTimestamp(m.UpdatedAt)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}
		date, err := utils.ParseDate(m.Date)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}
		rawData := utils.ParseRawJSON(m.RawData)
		additionalInfo := utils.ParseRawJSON(m.AdditionalInfo)

		params[i] = tietoevrysqlc.InsertMeasurementParams{
			ID:             id,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
			UserID:         userID,
			Date:           date,
			Name:           m.Name,
			NameType:       m.NameType,
			Source:         m.Source,
			Value:          m.Value,
			ValueNumeric:   utils.NullFloat64Ptr(m.ValueNumeric),
			Comment:        utils.NullStringPtr(m.Comment),
			RawID:          utils.NullStringPtr(m.RawID),
			RawData:        rawData,
			AdditionalInfo: additionalInfo,
		}
	}

	if err := h.store.InsertMeasurementsBulk(r.Context(), params); err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type TietoevryMeasurementParams struct {
	UserID string `form:"user_id" validate:"required,uuid4"`
}

// GetMeasurements godoc
//
//	@Summary		Get measurements by user ID
//	@Description	Get all measurements for a specific user
//	@Tags			Tietoevry - Measurements
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		string	true	"User ID (UUID)"
//	@Success		200		{object}	swagger.TietoevryMeasurementResponse
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/tietoevry/measurements [get]
func (h *TietoevryMeasurementHandler) GetMeasurements(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{"user_id"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := TietoevryMeasurementParams{
		UserID: r.URL.Query().Get("user_id"),
	}

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	cacheKey := fmt.Sprintf("tietoevry:measurements:%s", params.UserID)
	if h.cache != nil {
		if cached, err := h.cache.Get(r.Context(), cacheKey); err == nil && cached != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(cached))
			return
		}
	}

	userID, err := utils.ParseUUID(params.UserID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	measurements, err := h.store.GetMeasurementsByUser(r.Context(), userID)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	if len(measurements) == 0 {
		utils.WriteJSON(w, http.StatusOK, map[string]any{
			"measurements": []swagger.TietoevryMeasurementInput{},
		})
		return
	}

	var output []swagger.TietoevryMeasurementInput
	for _, measurement := range measurements {
		out := swagger.TietoevryMeasurementInput{
			ID:             measurement.ID.String(),
			CreatedAt:      measurement.CreatedAt.Format(time.RFC3339),
			UpdatedAt:      measurement.UpdatedAt.Format(time.RFC3339),
			UserID:         measurement.UserID.String(),
			Date:           measurement.Date.Format("2006-01-02"),
			Name:           measurement.Name,
			NameType:       measurement.NameType,
			Source:         measurement.Source,
			Value:          measurement.Value,
			ValueNumeric:   utils.Float64PtrOrNil(measurement.ValueNumeric),
			Comment:        utils.StringPtrOrNil(measurement.Comment),
			RawID:          utils.StringPtrOrNil(measurement.RawID),
			RawData:        utils.RawMessagePtrOrNil(measurement.RawData),
			AdditionalInfo: utils.RawMessagePtrOrNil(measurement.AdditionalInfo),
		}
		output = append(output, out)
	}

	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, output, 3*time.Minute)

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"measurements": output,
	})
}
