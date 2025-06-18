package utvapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authz"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/DeRuina/KUHA-REST-API/internal/store/utv"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
	"github.com/google/uuid"
)

// Validation structs
type LatestDataInput struct {
	UserID string `form:"user_id" validate:"required,uuid4"`
	Type   string `form:"type" validate:"required,key"`
	Device string `form:"device" validate:"omitempty,oneof=garmin oura polar suunto"`
	Limit  int32  `form:"limit" validate:"omitempty,min=1,max=100"`
}

// store and cache interfaces
type GeneralDataHandler struct {
	oura   utv.OuraData
	polar  utv.PolarData
	suunto utv.SuuntoData
	garmin utv.GarminData
	cache  *cache.Storage
}

// response structs
type LatestDataResponse struct {
	Device string          `json:"device"`
	Date   string          `json:"date"`
	Data   json.RawMessage `json:"data"`
}

// NewGeneralDataHandler initializes the handler
func NewGeneralDataHandler(oura utv.OuraData, polar utv.PolarData, suunto utv.SuuntoData, garmin utv.GarminData, cache *cache.Storage) *GeneralDataHandler {
	return &GeneralDataHandler{
		oura:   oura,
		polar:  polar,
		suunto: suunto,
		garmin: garmin,
		cache:  cache,
	}
}

// GetLatestData godoc
//
//	@Summary		Get latest data by type
//	@Description	Returns latest entries of a specific type for a user, optionally filtered by device and limited in number
//	@Tags			UTV - General
//	@Accept			json
//	@Produce		json
//	@Param			user_id	query	string						true	"User ID (UUID)"
//	@Param			type	query	string						true	"Data type (e.g., 'sleep', 'activity')"
//	@Param			device	query	string						false	"Device type (one of: 'garmin', 'oura', 'polar', 'suunto')"
//	@Param			limit	query	int							false	"Limit the number of results (default: 5, max: 100)"
//	@Success		200		{array}	swagger.LatestDataResponse	"Latest Data"
//	@Success		204		"No Content: No data found"
//	@Failure		400		{object}	swagger.ValidationErrorResponse
//	@Failure		403		{object}	swagger.ForbiddenResponse
//	@Failure		500		{object}	swagger.InternalServerErrorResponse
//	@Security		BearerAuth
//	@Router			/utv/latest [get]
func (h *GeneralDataHandler) GetLatestData(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	err := utils.ValidateParams(r, []string{"user_id", "type", "device", "limit"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := LatestDataInput{
		UserID: r.URL.Query().Get("user_id"),
		Type:   r.URL.Query().Get("type"),
		Device: r.URL.Query().Get("device"),
	}
	limitStr := r.URL.Query().Get("limit")
	if limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil {
			params.Limit = int32(parsedLimit)
		}
	}
	if params.Limit == 0 {
		params.Limit = 5 // default
	}

	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	userID, err := utils.ParseUUID(params.UserID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	cacheKey := fmt.Sprintf("latest:%s:%s:%s:%d", params.UserID, params.Type, params.Device, params.Limit)
	if h.cache != nil {
		if cached, err := h.cache.Get(r.Context(), cacheKey); err == nil && cached != "" {
			utils.WriteJSON(w, http.StatusOK, json.RawMessage(cached))
			return
		}
	}

	var results []LatestDataResponse

	// Helper to fetch from one device
	fetch := func(name string, store interface {
		GetLatestByType(ctx context.Context, userID uuid.UUID, typ string, limit int32) ([]utv.LatestDataEntry, error)
	}) {
		data, err := store.GetLatestByType(r.Context(), userID, params.Type, params.Limit)
		if err != nil {
			return // silently ignore errors per device
		}
		for _, row := range data {
			results = append(results, LatestDataResponse{
				Device: name,
				Date:   row.Date.Format("2006-01-02"),
				Data:   row.Data,
			})
		}
	}

	// Conditional fetch
	switch params.Device {
	case "garmin":
		fetch("garmin", h.garmin)
	case "oura":
		fetch("oura", h.oura)
	case "polar":
		fetch("polar", h.polar)
	case "suunto":
		fetch("suunto", h.suunto)
	default:
		fetch("garmin", h.garmin)
		fetch("oura", h.oura)
		fetch("polar", h.polar)
		fetch("suunto", h.suunto)
	}

	if len(results) == 0 {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	cache.SetCacheJSON(r.Context(), h.cache, cacheKey, results, 10*time.Minute)
	utils.WriteJSON(w, http.StatusOK, results)
}
