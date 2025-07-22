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
	Date           string  `json:"date" validate:"required"` // ISO 8601 date
	Symptom        string  `json:"symptom" validate:"required"`
	Severity       int32   `json:"severity" validate:"required"`
	Comment        *string `json:"comment"`
	Source         string  `json:"source" validate:"required"`
	CreatedAt      string  `json:"created_at" validate:"required"` // ISO 8601 datetime
	UpdatedAt      string  `json:"updated_at" validate:"required"` // ISO 8601 datetime
	RawID          *string `json:"raw_id"`
	OriginalID     *string `json:"original_id"`
	Recovered      *bool   `json:"recovered"`
	PainIndex      *int32  `json:"pain_index"`
	Side           *string `json:"side"`
	Category       *string `json:"category"`
	AdditionalData *string `json:"additional_data"` // JSON as string
}

type TietoevrySymptomsBulkInput struct {
	Symptoms []TietoevrySymptomInput `json:"symptoms" validate:"required,dive"`
}

// InsertSymptoms godoc
//
//	@Summary		Insert symptoms (bulk)
//	@Description	Insert multiple symptoms with idempotent behavior
//	@Tags			Tietoevry - Symptoms
//	@Accept			json
//	@Produce		json
//	@Param			symptoms	body	swagger.TietoevrySymptomsBulkInput	true	"Symptom data"
//	@Success		201			"Symptoms processed successfully (idempotent operation)"
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
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
