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
type TietoevrySymptomHandler struct {
	store tietoevry.Symptoms
	cache *cache.Storage
}

func NewTietoevrySymptomHandler(store tietoevry.Symptoms, cache *cache.Storage) *TietoevrySymptomHandler {
	return &TietoevrySymptomHandler{store: store, cache: cache}
}

// Input struct for validation
type TietoevrySymptomInput struct {
	ID             string  `json:"id" validate:"required,uuid4"`
	UserID         string  `json:"user_id" validate:"required,uuid4"`
	Date           string  `json:"date" validate:"required"`
	Symptom        string  `json:"symptom" validate:"required"`
	Severity       int32   `json:"severity" validate:"required"`
	Comment        *string `json:"comment"`
	Source         string  `json:"source" validate:"required"`
	CreatedAt      string  `json:"created_at" validate:"required"`
	UpdatedAt      string  `json:"updated_at" validate:"required"`
	RawID          *string `json:"raw_id"`
	OriginalID     *string `json:"original_id"`
	Recovered      *bool   `json:"recovered"`
	PainIndex      *int32  `json:"pain_index"`
	Side           *string `json:"side"`
	Category       *string `json:"category"`
	AdditionalData *string `json:"additional_data"`
}

type TietoevrySymptomsBulkInput struct {
	Symptoms []TietoevrySymptomInput `json:"symptoms" validate:"required,dive"`
}

// InsertSymptoms godoc
//
//	@Summary		Insert symptoms (bulk)
//	@Description	Insert multiple symptoms for user (idempotent)
//	@Tags			Tietoevry - Symptoms
//	@Accept			json
//	@Produce		json
//	@Param			symptoms	body	swagger.TietoevrySymptomsBulkInput	true	"Symptom data"
//	@Success		201			"Symptoms processed successfully (idempotent operation)"
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		401			{object}	swagger.UnauthorizedResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		409			{object}	swagger.ConflictResponse
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
//	@Failure		503			{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/tietoevry/symptoms [post]
func (h *TietoevrySymptomHandler) InsertSymptomsBulk(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var input TietoevrySymptomsBulkInput
	if err := utils.ReadJSON(w, r, &input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	if err := utils.GetValidator().Struct(input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	// PreValidation
	userIDs := make([]uuid.UUID, len(input.Symptoms))
	for i, symptom := range input.Symptoms {
		userID, err := utils.ParseUUID(symptom.UserID)
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

	// Convert to database parameters
	symptoms := make([]tietoevrysqlc.InsertSymptomParams, len(input.Symptoms))
	for i, symptom := range input.Symptoms {
		// Parse and convert values
		id, err := utils.ParseUUID(symptom.ID)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		userID, err := utils.ParseUUID(symptom.UserID)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		date, err := utils.ParseDate(symptom.Date)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		createdAt, err := utils.ParseTimestamp(symptom.CreatedAt)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		updatedAt, err := utils.ParseTimestamp(symptom.UpdatedAt)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		originalID, err := utils.ParseUUIDPtr(symptom.OriginalID)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		rawData := utils.ParseRawJSON(symptom.AdditionalData)

		symptoms[i] = tietoevrysqlc.InsertSymptomParams{
			ID:             id,
			UserID:         userID,
			Date:           date,
			Symptom:        symptom.Symptom,
			Severity:       symptom.Severity,
			Comment:        utils.NullStringPtr(symptom.Comment),
			Source:         symptom.Source,
			CreatedAt:      createdAt,
			UpdatedAt:      updatedAt,
			RawID:          utils.NullStringPtr(symptom.RawID),
			OriginalID:     originalID,
			Recovered:      utils.NullBoolPtr(symptom.Recovered),
			PainIndex:      utils.NullInt32Ptr(symptom.PainIndex),
			Side:           utils.NullStringPtr(symptom.Side),
			Category:       utils.NullStringPtr(symptom.Category),
			AdditionalData: rawData,
		}
	}

	if err := h.store.InsertSymptomsBulk(r.Context(), symptoms); err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type TietoevrySymptomParams struct {
	UserID string `form:"user_id" validate:"required,uuid4"`
}

// GetSymptoms godoc
//
//	@Summary		Get symptoms by user ID
//	@Description	Get all symptoms for a specific user
//	@Tags			Tietoevry - Symptoms
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		string	true	"User ID (UUID)"
//	@Success		200		{object}	swagger.TietoevrySymptomResponse
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/tietoevry/symptoms [get]
func (h *TietoevrySymptomHandler) GetSymptoms(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{"user_id"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := TietoevrySymptomParams{
		UserID: r.URL.Query().Get("user_id"),
	}

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	cacheKey := fmt.Sprintf("tietoevry:symptoms:%s", params.UserID)
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

	symptoms, err := h.store.GetSymptomsByUser(r.Context(), userID)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	if len(symptoms) == 0 {
		utils.WriteJSON(w, http.StatusOK, map[string]any{
			"symptoms": []swagger.TietoevrySymptomInput{},
		})
		return
	}

	var output []swagger.TietoevrySymptomInput
	for _, symptom := range symptoms {
		out := swagger.TietoevrySymptomInput{
			ID:             symptom.ID.String(),
			UserID:         symptom.UserID.String(),
			Date:           symptom.Date.Format("2006-01-02"),
			Symptom:        symptom.Symptom,
			Severity:       symptom.Severity,
			Comment:        utils.StringPtrOrNil(symptom.Comment),
			Source:         symptom.Source,
			CreatedAt:      symptom.CreatedAt.Format(time.RFC3339),
			UpdatedAt:      symptom.UpdatedAt.Format(time.RFC3339),
			RawID:          utils.StringPtrOrNil(symptom.RawID),
			OriginalID:     utils.UUIDPtrToStringPtr(symptom.OriginalID),
			Recovered:      utils.BoolPtrOrNil(symptom.Recovered),
			PainIndex:      utils.Int32PtrOrNil(symptom.PainIndex),
			Side:           utils.StringPtrOrNil(symptom.Side),
			Category:       utils.StringPtrOrNil(symptom.Category),
			AdditionalData: utils.RawMessagePtrOrNil(symptom.AdditionalData),
		}
		output = append(output, out)
	}

	resp := map[string]any{"symptoms": output}
	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, resp, 3*time.Minute)
	utils.WriteJSON(w, http.StatusOK, resp)
}
