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

type TietoevryQuestionnaireHandler struct {
	store tietoevry.Questionnaires
	cache *cache.Storage
}

func NewTietoevryQuestionnaireHandler(store tietoevry.Questionnaires, cache *cache.Storage) *TietoevryQuestionnaireHandler {
	return &TietoevryQuestionnaireHandler{store: store, cache: cache}
}

type TietoevryQuestionnaireAnswerInput struct {
	UserID                  string  `json:"user_id" validate:"required,uuid4"`
	QuestionnaireInstanceID string  `json:"questionnaire_instance_id" validate:"required,uuid4"`
	QuestionnaireNameFi     *string `json:"questionnaire_name_fi"`
	QuestionnaireNameEn     *string `json:"questionnaire_name_en"`
	QuestionnaireKey        string  `json:"questionnaire_key" validate:"required"`
	QuestionID              string  `json:"question_id" validate:"required,uuid4"`
	QuestionLabelFi         *string `json:"question_label_fi"`
	QuestionLabelEn         *string `json:"question_label_en"`
	QuestionType            string  `json:"question_type" validate:"required"`
	OptionID                *string `json:"option_id"`
	OptionValue             *int32  `json:"option_value"`
	OptionLabelFi           *string `json:"option_label_fi"`
	OptionLabelEn           *string `json:"option_label_en"`
	FreeText                *string `json:"free_text"`
	CreatedAt               string  `json:"created_at" validate:"required"`
	UpdatedAt               string  `json:"updated_at" validate:"required"`
	Value                   *string `json:"value"`
}

type TietoevryQuestionnaireAnswersBulkInput struct {
	Questionnaires []TietoevryQuestionnaireAnswerInput `json:"questionnaires" validate:"required,dive"`
}

// InsertQuestionnaireAnswers godoc
//
//	@Summary		Insert questionnaire answers (bulk)
//	@Description	Insert multiple questionnaire answers for users (idempotent)
//	@Tags			Tietoevry - Questionnaires
//	@Accept			json
//	@Produce		json
//	@Param			questionnaires	body	swagger.TietoevryQuestionnaireAnswersBulkInput	true	"Questionnaire answers"
//	@Success		201				"Questionnaire answers processed successfully"
//	@Failure		400				{object}	swagger.ValidationErrorResponse
//	@Failure		403				{object}	swagger.ForbiddenResponse
//	@Failure		409				{object}	swagger.ConflictResponse
//	@Failure		500				{object}	swagger.InternalServerErrorResponse
//	@Security		BearerAuth
//	@Router			/tietoevry/questionnaires [post]
func (h *TietoevryQuestionnaireHandler) InsertQuestionnaireAnswersBulk(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var input TietoevryQuestionnaireAnswersBulkInput
	if err := utils.ReadJSON(w, r, &input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	if err := utils.GetValidator().Struct(input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	// PreValidation
	userIDs := make([]uuid.UUID, len(input.Questionnaires))
	for i, a := range input.Questionnaires {
		id, err := utils.ParseUUID(a.UserID)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}
		userIDs[i] = id
	}

	if err := h.store.ValidateUsersExist(r.Context(), userIDs); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	questionnaires := make([]tietoevrysqlc.InsertQuestionnaireAnswerParams, len(input.Questionnaires))
	for i, questionnaire := range input.Questionnaires {
		// Parse and convert values
		userID, err := utils.ParseUUID(questionnaire.UserID)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		questionnaireInstanceID, err := utils.ParseUUID(questionnaire.QuestionnaireInstanceID)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		questionID, err := utils.ParseUUID(questionnaire.QuestionID)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		createdAt, err := utils.ParseTimestamp(questionnaire.CreatedAt)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		updatedAt, err := utils.ParseTimestamp(questionnaire.UpdatedAt)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		optionID, err := utils.ParseUUIDPtr(questionnaire.OptionID)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		valueJSON := utils.ParseRawJSON(questionnaire.Value)

		questionnaires[i] = tietoevrysqlc.InsertQuestionnaireAnswerParams{
			UserID:                  userID,
			QuestionnaireInstanceID: questionnaireInstanceID,
			QuestionnaireNameFi:     utils.NullStringPtr(questionnaire.QuestionnaireNameFi),
			QuestionnaireNameEn:     utils.NullStringPtr(questionnaire.QuestionnaireNameEn),
			QuestionnaireKey:        questionnaire.QuestionnaireKey,
			QuestionID:              questionID,
			QuestionLabelFi:         utils.NullStringPtr(questionnaire.QuestionLabelFi),
			QuestionLabelEn:         utils.NullStringPtr(questionnaire.QuestionLabelEn),
			QuestionType:            questionnaire.QuestionType,
			OptionID:                optionID,
			OptionValue:             utils.NullInt32Ptr(questionnaire.OptionValue),
			OptionLabelFi:           utils.NullStringPtr(questionnaire.OptionLabelFi),
			OptionLabelEn:           utils.NullStringPtr(questionnaire.OptionLabelEn),
			FreeText:                utils.NullStringPtr(questionnaire.FreeText),
			CreatedAt:               createdAt,
			UpdatedAt:               updatedAt,
			Value:                   valueJSON,
		}
	}

	if err := h.store.InsertQuestionnaireAnswersBulk(r.Context(), questionnaires); err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
