package utvapi

import (
	"context"
	"errors"
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

// Get available dates from Oura data (with optional filtering)
func (h *OuraDataHandler) GetDates(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	afterDate := r.URL.Query().Get("after_date")
	beforeDate := r.URL.Query().Get("before_date")

	if userID == "" {
		utils.BadRequestResponse(w, r, utils.ErrMissingUserID)
		return
	}

	err := utils.ValidateParams(r, []string{"user_id", "after_date", "before_date"})
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	if afterDate != "" && beforeDate != "" && afterDate > beforeDate {
		utils.UnprocessableEntityResponse(w, r, utils.ErrInvalidDateRange)
		return
	}

	dates, err := h.store.GetDates(context.Background(), userID, &afterDate, &beforeDate)
	if err != nil {
		switch {
		case errors.Is(err, utils.ErrInvalidUUID):
			utils.BadRequestResponse(w, r, err)
		case errors.Is(err, utils.ErrInvalidDate):
			utils.BadRequestResponse(w, r, err)
		default:
			utils.InternalServerError(w, r, err)
		}
		return
	}

	if len(dates) == 0 {
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	utils.WriteJSON(w, http.StatusOK, dates)
}

// // Get all JSON keys from Oura data (with optional filtering)
// func (h *OuraDataHandler) GetTypes(w http.ResponseWriter, r *http.Request) {
// 	userID := r.URL.Query().Get("user_id")
// 	specificDate := r.URL.Query().Get("specific_date")
// 	startDate := r.URL.Query().Get("start_date")
// 	endDate := r.URL.Query().Get("end_date")

// 	if userID == "" {
// 		utils.BadRequestResponse(w, r, fmt.Errorf("user_id is required"))
// 		return
// 	}

// 	types, err := h.store.GetTypes(context.Background(), userID, &specificDate, &startDate, &endDate)
// 	if err != nil {
// 		utils.InternalServerError(w, r, err)
// 		return
// 	}

// 	utils.WriteJSON(w, http.StatusOK, types)
// }

// // Get all data for a specific date (or filter by type)
// func (h *OuraDataHandler) GetData(w http.ResponseWriter, r *http.Request) {
// 	userID := r.URL.Query().Get("user_id")
// 	summaryDate := r.URL.Query().Get("summary_date")
// 	key := r.URL.Query().Get("key")

// 	if userID == "" || summaryDate == "" {
// 		utils.BadRequestResponse(w, r, fmt.Errorf("user_id and summary_date are required"))
// 		return
// 	}

// 	data, err := h.store.GetData(context.Background(), userID, summaryDate, &key)
// 	if err != nil {
// 		utils.InternalServerError(w, r, err)
// 		return
// 	}

// 	utils.WriteJSON(w, http.StatusOK, data)
// }
