package utvapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DeRuina/KUHA-REST-API/internal/store/utv"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type OuraDataHandler struct {
	store utv.OuraData
}

// NewOuraDataHandler initializes OuraData handler
func NewOuraDataHandler(store utv.OuraData) *OuraDataHandler {
	return &OuraDataHandler{store: store}
}

// Get available dates from Oura data
func (h *OuraDataHandler) GetDates(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if userID == "" || startDate == "" || endDate == "" {
		utils.BadRequestResponse(w, r, fmt.Errorf("missing required parameters"))
		return
	}

	dates, err := h.store.GetDates(context.Background(), userID, startDate, endDate)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, dates)
}

// Get all JSON keys from Oura data
func (h *OuraDataHandler) GetTypes(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	summaryDate := r.URL.Query().Get("summary_date")

	if userID == "" || summaryDate == "" {
		utils.BadRequestResponse(w, r, fmt.Errorf("missing required parameters"))
		return
	}

	types, err := h.store.GetTypes(context.Background(), userID, summaryDate)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, types)
}

// Get a specific data point from Oura data
func (h *OuraDataHandler) GetDataPoint(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	summaryDate := r.URL.Query().Get("summary_date")
	key := r.URL.Query().Get("key")

	if userID == "" || summaryDate == "" || key == "" {
		utils.BadRequestResponse(w, r, fmt.Errorf("missing required parameters"))
		return
	}

	data, err := h.store.GetDataPoint(context.Background(), userID, summaryDate, key)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, data)
}

// Get unique JSON keys from Oura data over a date range
func (h *OuraDataHandler) GetUniqueTypes(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	startDate := r.URL.Query().Get("start_date")
	endDate := r.URL.Query().Get("end_date")

	if userID == "" || startDate == "" || endDate == "" {
		utils.BadRequestResponse(w, r, fmt.Errorf("missing required parameters"))
		return
	}

	types, err := h.store.GetUniqueTypes(context.Background(), userID, startDate, endDate)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, types)
}
