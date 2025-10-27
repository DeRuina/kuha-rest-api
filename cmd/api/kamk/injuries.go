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
type InjuriesHandler struct {
	store kamk.Injuries
	cache *cache.Storage
}

func NewInjuriesHandler(store kamk.Injuries, cache *cache.Storage) *InjuriesHandler {
	return &InjuriesHandler{store: store, cache: cache}
}

// Validation structs
type KamkAddInjuryInput struct {
	UserID      int32   `json:"user_id" validate:"required,gt=0"`
	InjuryType  int32   `json:"injury_type" validate:"required"`
	Severity    *int32  `json:"severity" validate:"omitempty"`
	PainLevel   *int32  `json:"pain_level" validate:"omitempty"`
	Description *string `json:"description" validate:"omitempty"`
	InjuryID    int32   `json:"injury_id" validate:"required"`
	Meta        *string `json:"meta" validate:"omitempty"`
}

type KamkMarkRecoveredInput struct {
	UserID   int32 `json:"user_id" validate:"required,gt=0"`
	InjuryID int32 `json:"injury_id" validate:"required"`
}

type KamkGetInjuriesParams struct {
	UserID int32 `form:"user_id" validate:"required,gt=0"`
}

type KamkGetMaxIDParams struct {
	UserID int32 `form:"user_id" validate:"required,gt=0"`
}

type KamkDeleteInjuryParams struct {
	UserID   int32 `form:"user_id"   validate:"required,gt=0"`
	InjuryID int32 `form:"injury_id" validate:"required,gt=0"`
}

// AddInjury godoc
//
//	@Summary		Create injury
//	@Description	Creates a new injury row (status=0, date_start=NOW()) for a competitor
//	@Tags			KAMK - Injuries
//	@Accept			json
//	@Produce		json
//	@Param			body	body	swagger.KamkAddInjuryRequest	true	"Injury payload"
//	@Success		201		"Created: Injury stored (no content in response body)"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/kamk/injury [post]
func (h *InjuriesHandler) AddInjury(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var input KamkAddInjuryInput
	if err := utils.ReadJSON(w, r, &input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	if err := utils.GetValidator().Struct(input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	err := h.store.AddInjury(r.Context(), input.UserID, kamk.InjuryInput{
		InjuryType:  input.InjuryType,
		Severity:    input.Severity,
		PainLevel:   input.PainLevel,
		Description: input.Description,
		InjuryID:    input.InjuryID,
		Meta:        input.Meta,
	})
	if err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	invalidateKamkInjuries(r.Context(), h.cache, input.UserID)

	w.WriteHeader(http.StatusCreated)
}

// MarkRecovered godoc
//
//	@Summary		Mark injury recovered
//	@Description	Sets status=1 and date_end=NOW() for an injury recovery
//	@Tags			KAMK - Injuries
//	@Accept			json
//	@Produce		json
//	@Param			body	body	swagger.KamkMarkRecoveredRequest	true	"Recovery payload"
//	@Success		200		"OK: Marked recovered"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/kamk/injury-recovered [post]
func (h *InjuriesHandler) MarkRecovered(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	var input KamkMarkRecoveredInput
	if err := utils.ReadJSON(w, r, &input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	if err := utils.GetValidator().Struct(input); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	if _, err := h.store.MarkInjuryRecovered(r.Context(), input.UserID, input.InjuryID); err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	invalidateKamkInjuries(r.Context(), h.cache, input.UserID)

	w.WriteHeader(http.StatusOK)
}

// GetActive godoc
//
//	@Summary		List active injuries
//	@Description	Returns active injuries (status=0) for a competitor
//	@Tags			KAMK - Injuries
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		integer	true	"Competitor sportti_id"
//	@Success		200		{object}	swagger.KamkInjuriesListResponse
//	@Success		204		"No Content: no injuries"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/kamk/injury [get]
func (h *InjuriesHandler) GetActive(w http.ResponseWriter, r *http.Request) {
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

	params := KamkGetInjuriesParams{
		UserID: uid,
	}
	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	cacheKey := fmt.Sprintf("kamk:injury:list:%d", uid)
	if h.cache != nil {
		if cached, err := h.cache.Get(r.Context(), cacheKey); err == nil && cached != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(cached))
			return
		}
	}

	items, err := h.store.GetActiveInjuries(r.Context(), uid)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}
	if len(items) == 0 {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	resp := map[string]any{"injuries": items}
	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, resp, KAMKCacheTTL)
	utils.WriteJSON(w, http.StatusOK, resp)
}

// GetMaxID godoc
//
//	@Summary		Get next injury id helper
//	@Description	Returns the current maximum injury_id for a competitor (0 if none exist)
//	@Tags			KAMK - Injuries
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query		integer	true	"Competitor sportti_id"
//	@Success		200		{object}	swagger.KamkMaxInjuryIDResponse
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		401		{object}	swagger.UnauthorizedResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Failure		503		{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/kamk/injury-id [get]
func (h *InjuriesHandler) GetMaxID(w http.ResponseWriter, r *http.Request) {
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

	params := KamkGetMaxIDParams{
		UserID: uid,
	}
	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	id, err := h.store.GetMaxInjuryID(r.Context(), uid)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	resp := map[string]int32{"id": id}
	utils.WriteJSON(w, http.StatusOK, resp)
}

// DeleteInjury godoc
//
//	@Summary		Delete an injury by injury_id
//	@Description	Deletes a single injury for a competitor
//	@Tags			KAMK - Injuries
//	@Accept			json
//	@Produce		json
//	@Param			user_id		query	integer	true	"sportti_id"
//	@Param			injury_id	query	integer	true	"Injury ID"
//	@Success		200			"OK: deleted"
//	@Failure		400			{object}	swagger.ValidationErrorResponse
//	@Failure		401			{object}	swagger.UnauthorizedResponse
//	@Failure		403			{object}	swagger.ForbiddenResponse
//	@Failure		404			{object}	swagger.NotFoundResponse
//	@Failure		500			{object}	swagger.InternalServerErrorResponse
//	@Failure		503			{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/kamk/injury [delete]
func (h *InjuriesHandler) DeleteInjury(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}
	if err := utils.ValidateParams(r, []string{"user_id", "injury_id"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	var p KamkDeleteInjuryParams
	uid, err := utils.ParsePositiveInt32(r.URL.Query().Get("user_id"))
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	iid, err := utils.ParsePositiveInt32(r.URL.Query().Get("injury_id"))
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}
	p.UserID, p.InjuryID = uid, iid

	if err := utils.GetValidator().Struct(p); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	n, err := h.store.DeleteInjury(r.Context(), p.UserID, p.InjuryID)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}
	if n == 0 {
		utils.NotFoundResponse(w, r, fmt.Errorf("injury not found"))
		return
	}

	invalidateKamkInjuries(r.Context(), h.cache, p.UserID)
	w.WriteHeader(http.StatusOK)
}
