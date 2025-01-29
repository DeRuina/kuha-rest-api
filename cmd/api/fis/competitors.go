package fisapi

import (
	"context"
	"net/http"

	"github.com/DeRuina/KUHA-REST-API/internal/store/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type CompetitorsHandler struct {
	store fis.Competitors // ✅ Use the `Competitors` interface
}

// ✅ Modify to accept `fis.Competitors` interface
func NewCompetitorsHandler(store fis.Competitors) *CompetitorsHandler {
	return &CompetitorsHandler{store: store}
}

// GetBySector handles the GET request to fetch competitors by sector
func (h *CompetitorsHandler) GetBySector(w http.ResponseWriter, r *http.Request) {
	sectorCode := r.URL.Query().Get("sector")
	if sectorCode == "" {
		http.Error(w, "Sector code is required", http.StatusBadRequest)
		return
	}

	competitors, err := h.store.GetBySector(context.Background(), sectorCode)
	if err != nil {
		http.Error(w, "Error fetching competitors", http.StatusInternalServerError)
		return
	}

	utils.WriteJSON(w, http.StatusOK, competitors)
}
