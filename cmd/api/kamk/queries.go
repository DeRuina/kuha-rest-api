package kamkapi

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	UserID    int32   `json:"user_id" validate:"required,gt=0"`
	QueryType *int32  `json:"query_type" validate:"omitempty"`
	Answers   *string `json:"answers" validate:"omitempty"`
	Comment   *string `json:"comment" validate:"omitempty"`
	Meta      *string `json:"meta" validate:"omitempty"`
}

type KamkGetQuestionnairesParams struct {
	UserID int32 `json:"user_id" validate:"required,gt=0"`
}

type KamkIsQuizDoneParams struct {
	UserID   int32 `json:"user_id" validate:"required,gt=0"`
	QuizType int32 `form:"quiz_type" validate:"gte=0"`
}

type KamkUpdateQuestionnaireQuery struct {
	UserID int32 `form:"user_id" validate:"required,gt=0"`
	ID     int64 `form:"id" validate:"required,gt=0"`
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
//	@Success		201		{object}	swagger.KamkCreateQuestionnaireResponse
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

	id, err := h.store.AddQuestionnaire(r.Context(), input.UserID, kamk.QuestionnaireInput{
		QueryType: input.QueryType,
		Answers:   input.Answers,
		Comment:   input.Comment,
		Meta:      input.Meta,
	})
	if err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	invalidateKamkQueries(r.Context(), h.cache, input.UserID)
	utils.WriteJSON(w, http.StatusCreated, map[string]int64{"id": id})
}

// GetQuestionnaires godoc
//
//	@Summary		List questionnaires
//	@Description	Returns questionnaires for a competitor ordered by timestamp DESC
//	@Tags			KAMK - Queries
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		integer	true	"sportti_id"
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

	uidStr := r.URL.Query().Get("user_id")
	uid, err := utils.ParsePositiveInt32(uidStr)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := KamkGetQuestionnairesParams{
		UserID: uid,
	}
	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	cacheKey := fmt.Sprintf("kamk:queries:list:%d", uid)
	if h.cache != nil {
		if cached, err := h.cache.Get(r.Context(), cacheKey); err == nil && cached != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(cached))
			return
		}
	}

	items, err := h.store.GetQuestionnaires(r.Context(), uid)
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
	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, resp, KAMKCacheTTL)
	utils.WriteJSON(w, http.StatusOK, resp)
}

// IsQuizDoneToday godoc
//
//	@Summary		Check if certain quiz is due for the user. Used for daily quizzes.
//	@Description	Return the quiz if it was done today
//	@Tags			KAMK - Queries
//	@Accept			json
//	@Produce		json
//	@Param			user_id		query		integer	true	"sportti_id"
//	@Param			quiz_type	query		integer	true	"Quiz type"
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

	uidStr := r.URL.Query().Get("user_id")
	uid, err := utils.ParsePositiveInt32(uidStr)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := KamkIsQuizDoneParams{
		UserID: uid,
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

	rows, err := h.store.IsQuizDoneToday(r.Context(), uid, params.QuizType)
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

// UpdateQuestionnaireByID godoc
//
//	@Summary		Update questionnaire by id
//	@Description	Updates the questionnaire row identified by id (and user_id)
//	@Tags			KAMK - Queries
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query	integer								true	"sportti_id"
//	@Param			id		query	integer								true	"Questionnaire ID"
//	@Param			body	body	swagger.KamkUpdateQuestionnaireBody	true	"Update payload"
//	@Success		200		"OK: updated"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		404		{object}	swagger.NotFoundResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/kamk/update-quiz [post]
func (h *QueriesHandler) UpdateQuestionnaireByTimestamp(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{"user_id", "id"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	uidStr := r.URL.Query().Get("user_id")
	uid, err := utils.ParsePositiveInt32(uidStr)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	qid, err := utils.ParsePositiveInt64(r.URL.Query().Get("id"))
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	qp := KamkUpdateQuestionnaireQuery{
		UserID: uid,
		ID:     qid,
	}
	if err := utils.GetValidator().Struct(qp); err != nil {
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

	n, err := h.store.UpdateQuestionnaireByID(r.Context(), uid, qid, body.Answers, body.Comment)
	if err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	if n == 0 {
		utils.NotFoundResponse(w, r, fmt.Errorf("no questionnaire found for that id"))
		return
	}

	invalidateKamkQueries(r.Context(), h.cache, uid)

	w.WriteHeader(http.StatusOK)
}

// DeleteQuestionnaire godoc
//
//	@Summary		Delete a questionnaire by id
//	@Description	Deletes a questionnaire row (competitor_id=user_id AND id=id)
//	@Tags			KAMK - Queries
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query	integer	true	"sportti_id"
//	@Param			id		query	integer	true	"Questionnaire ID"
//	@Success		200		"OK: deleted"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		404		{object}	swagger.NotFoundResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/kamk/delete-quiz [delete]
func (h *QueriesHandler) DeleteQuestionnaire(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}
	if err := utils.ValidateParams(r, []string{"user_id", "id"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	uid, err := utils.ParsePositiveInt32(r.URL.Query().Get("user_id"))
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	qid, err := utils.ParsePositiveInt64(r.URL.Query().Get("id"))
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	q := KamkUpdateQuestionnaireQuery{
		ID:     qid,
		UserID: uid,
	}

	if err := utils.GetValidator().Struct(q); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	n, err := h.store.DeleteQuestionnaireByID(r.Context(), q.UserID, q.ID)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}
	if n == 0 {
		utils.NotFoundResponse(w, r, fmt.Errorf("no questionnaire found for that id"))
		return
	}

	invalidateKamkQueries(r.Context(), h.cache, q.UserID)
	w.WriteHeader(http.StatusOK)
}
