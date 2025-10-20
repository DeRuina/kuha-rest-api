package kamkapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authz"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/DeRuina/KUHA-REST-API/internal/store/kamk"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

// Handler
type QueriesHandler struct {
	store kamk.Queries
	cache *cache.Storage
}

func NewQueriesHandler(store kamk.Queries, cache *cache.Storage) *QueriesHandler {
	return &QueriesHandler{store: store, cache: cache}
}

// Validation structs
type KamkAddQuestionnaireInput struct {
	UserID    string  `json:"user_id" validate:"required,numeric"`
	QueryType *int32  `json:"query_type" validate:"omitempty"`
	Answers   *string `json:"answers" validate:"omitempty"`
	Comment   *string `json:"comment" validate:"omitempty"`
	Meta      *string `json:"meta" validate:"omitempty"`
}

type KamkGetQuestionnairesParams struct {
	UserID string `form:"user_id" validate:"required,numeric"`
}

type KamkIsQuizDoneParams struct {
	UserID   string `form:"user_id" validate:"required,numeric"`
	QuizType int32  `form:"quiz_type" validate:"min=0"`
}

type KamkUpdateQuestionnaireQuery struct {
	UserID    string `form:"user_id" validate:"required,numeric"`
	Timestamp string `form:"timestamp" validate:"required"`
}

type KamkUpdateQuestionnaireBody struct {
	Answers string  `json:"answers" validate:"required"`
	Comment *string `json:"comment" validate:"omitempty"`
}

// AddQuestionnaire godoc
//
//	@Summary		Create questionnaire entry
//	@Description	Inserts a questionnaire row for a competitor (timestamp=NOW())
//	@Tags			KAMK - Queries
//	@Accept			json
//	@Produce		json
//	@Param			body	body		swagger.KamkAddQuestionnaireRequest	true	"Questionnaire payload"
//	@Success		201		"Created: 	Query stored (no content)"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/kamk/questionnaire [post]
func (h *QueriesHandler) AddQuestionnaire(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var input KamkAddQuestionnaireInput
	if err := utils.ReadJSON(w, r, &input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	if err := utils.GetValidator().Struct(input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	err := h.store.AddQuestionnaire(r.Context(), input.UserID, kamk.QuestionnaireInput{
		QueryType: input.QueryType,
		Answers:   input.Answers,
		Comment:   input.Comment,
		Meta:      input.Meta,
	})
	if err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetQuestionnaires godoc
//
//	@Summary		List questionnaires
//	@Description	Returns questionnaires for a competitor ordered by timestamp DESC
//	@Tags			KAMK - Queries
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		string	true	"sportti_id"
//	@Success		200		{object}	swagger.KamkQuestionnairesListResponse
//	@Success		204		"No Content: no rows"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/kamk/questionnaire [get]
func (h *QueriesHandler) GetQuestionnaires(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{"user_id"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := KamkGetQuestionnairesParams{
		UserID: r.URL.Query().Get("user_id"),
	}
	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	cacheKey := fmt.Sprintf("kamk:queries:list:%s", params.UserID)
	if h.cache != nil {
		if cached, err := h.cache.Get(r.Context(), cacheKey); err == nil && cached != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(cached))
			return
		}
	}

	items, err := h.store.GetQuestionnaires(r.Context(), params.UserID)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}
	if len(items) == 0 {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	resp := map[string]any{"questionnaires": items}
	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, resp, 3*time.Minute)
	utils.WriteJSON(w, http.StatusOK, resp)
}

// IsQuizDoneToday godoc
//
//	@Summary		Check if certain quiz is due for the user. Used for daily quizzes.
//	@Description	Return the quiz if it was done today
//	@Tags			KAMK - Queries
//	@Accept			json
//	@Produce		json
//	@Param			user_id		query		string	true	"sportti_id"
//	@Param			quiz_type	query		int		true	"Quiz type"
//	@Success		200			{object}	swagger.KamkQuestionnairesListResponse
//	@Success		204			"No Content: no rows today"
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		401			{object}	swagger.UnauthorizedResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
//	@Failure		503			{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/kamk/is-quiz-done [get]
func (h *QueriesHandler) IsQuizDoneToday(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{"user_id", "quiz_type"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := KamkIsQuizDoneParams{
		UserID: r.URL.Query().Get("user_id"),
	}
	qtStr := r.URL.Query().Get("quiz_type")
	qt, err := utils.ParseNonNegativeInt32(qtStr)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	params.QuizType = qt

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	rows, err := h.store.IsQuizDoneToday(r.Context(), params.UserID, params.QuizType)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	if len(rows) == 0 {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{"questionnaires": rows})
}

// UpdateQuestionnaireByTimestamp godoc
//
//	@Summary		Update questionnaire by timestamp
//	@Description	Updates answers/comment for the exact timestamp row for a competitor
//	@Tags			KAMK - Queries
//	@Accept			json
//	@Produce		json
//	@Param			user_id		query	string								true	"sportti_id"
//	@Param			timestamp	query	string								true	"Timestamp (RFC3339)"
//	@Param			body		body	swagger.KamkUpdateQuestionnaireBody	true	"Update payload"
//	@Success		200			"OK: updated"
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		401			{object}	swagger.UnauthorizedResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		404			{object}	swagger.NotFoundResponse
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
//	@Failure		503			{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/kamk/update-quiz [post]
func (h *QueriesHandler) UpdateQuestionnaireByTimestamp(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{"user_id", "timestamp"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	qp := KamkUpdateQuestionnaireQuery{
		UserID:    r.URL.Query().Get("user_id"),
		Timestamp: r.URL.Query().Get("timestamp"),
	}
	if err := utils.GetValidator().Struct(qp); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	ts, err := utils.ParseRFC3339MinuteOrSecond(qp.Timestamp)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	var body KamkUpdateQuestionnaireBody
	if err := utils.ReadJSON(w, r, &body); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	if err := utils.GetValidator().Struct(body); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	if err := h.store.UpdateQuestionnaireByTimestamp(r.Context(), qp.UserID, ts, body.Answers, body.Comment); err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	w.WriteHeader(http.StatusOK)
}
