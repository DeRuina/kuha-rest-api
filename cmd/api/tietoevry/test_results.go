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

type TietoevryTestResultHandler struct {
	store tietoevry.TestResults
	cache *cache.Storage
}

func NewTietoevryTestResultHandler(store tietoevry.TestResults, cache *cache.Storage) *TietoevryTestResultHandler {
	return &TietoevryTestResultHandler{store: store, cache: cache}
}

type TietoevryTestResultInput struct {
	ID                          string  `json:"id" validate:"required,uuid4"`
	UserID                      string  `json:"user_id" validate:"required,uuid4"`
	TypeID                      string  `json:"type_id" validate:"required,uuid4"`
	TypeType                    *string `json:"type_type"`
	TypeResultType              string  `json:"type_result_type" validate:"required"`
	TypeName                    *string `json:"type_name"`
	Timestamp                   string  `json:"timestamp" validate:"required"`
	Name                        *string `json:"name"`
	Comment                     *string `json:"comment"`
	Data                        string  `json:"data" validate:"required"`
	CreatedAt                   string  `json:"created_at" validate:"required"`
	UpdatedAt                   string  `json:"updated_at" validate:"required"`
	TestEventID                 *string `json:"test_event_id"`
	TestEventName               *string `json:"test_event_name"`
	TestEventDate               *string `json:"test_event_date"`
	TestEventTemplateTestID     *string `json:"test_event_template_test_id"`
	TestEventTemplateTestName   *string `json:"test_event_template_test_name"`
	TestEventTemplateTestLimits *string `json:"test_event_template_test_limits"`
}

type TietoevryTestResultsBulkInput struct {
	TestResults []TietoevryTestResultInput `json:"test_results" validate:"required,dive"`
}

// InsertTestResultsBulk godoc
//
//	@Summary		Insert test results (bulk)
//	@Description	Insert multiple test results for users (idempotent)
//	@Tags			Tietoevry - TestResults
//	@Accept			json
//	@Produce		json
//	@Param			test_results	body	swagger.TietoevryTestResultsBulkInput	true	"Test result data"
//	@Success		201				"Test results processed successfully"
//	@Failure		400				{object}	swagger.ValidationErrorResponse
//	@Failure		403				{object}	swagger.ForbiddenResponse
//	@Failure		409				{object}	swagger.ConflictResponse
//	@Failure		500				{object}	swagger.InternalServerErrorResponse
//	@Security		BearerAuth
//	@Router			/tietoevry/test-results [post]
func (h *TietoevryTestResultHandler) InsertTestResultsBulk(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var input TietoevryTestResultsBulkInput
	if err := utils.ReadJSON(w, r, &input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	if err := utils.GetValidator().Struct(input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	// PreValidation
	userIDs := make([]uuid.UUID, len(input.TestResults))
	for i, t := range input.TestResults {
		userID, err := utils.ParseUUID(t.UserID)
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

	testResults := make([]tietoevrysqlc.InsertTestResultParams, len(input.TestResults))
	for i, testResult := range input.TestResults {
		// Parse and convert values
		id, err := utils.ParseUUID(testResult.ID)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		userID, err := utils.ParseUUID(testResult.UserID)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		typeID, err := utils.ParseUUID(testResult.TypeID)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		timestamp, err := utils.ParseTimestamp(testResult.Timestamp)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		createdAt, err := utils.ParseTimestamp(testResult.CreatedAt)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		updatedAt, err := utils.ParseTimestamp(testResult.UpdatedAt)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		testEventID, err := utils.ParseUUIDPtr(testResult.TestEventID)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		testEventDate, err := utils.ParseDatePtr(testResult.TestEventDate)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		testEventTemplateTestID, err := utils.ParseUUIDPtr(testResult.TestEventTemplateTestID)
		if err != nil {
			utils.BadRequestResponse(w, r, err)
			return
		}

		dataJSON := utils.ParseRequiredJSON(testResult.Data)
		templateLimitsJSON := utils.ParseRawJSON(testResult.TestEventTemplateTestLimits)

		testResults[i] = tietoevrysqlc.InsertTestResultParams{
			ID:                          id,
			UserID:                      userID,
			TypeID:                      typeID,
			TypeType:                    utils.NullStringPtr(testResult.TypeType),
			TypeResultType:              testResult.TypeResultType,
			TypeName:                    utils.NullStringPtr(testResult.TypeName),
			Timestamp:                   timestamp,
			Name:                        utils.NullStringPtr(testResult.Name),
			Comment:                     utils.NullStringPtr(testResult.Comment),
			Data:                        dataJSON,
			CreatedAt:                   createdAt,
			UpdatedAt:                   updatedAt,
			TestEventID:                 testEventID,
			TestEventName:               utils.NullStringPtr(testResult.TestEventName),
			TestEventDate:               utils.NullTimePtr(testEventDate),
			TestEventTemplateTestID:     testEventTemplateTestID,
			TestEventTemplateTestName:   utils.NullStringPtr(testResult.TestEventTemplateTestName),
			TestEventTemplateTestLimits: templateLimitsJSON,
		}
	}

	if err := h.store.InsertTestResultsBulk(r.Context(), testResults); err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
