package fisapi

import (
	"context"
	"fmt"
	"net/http"

	"github.com/DeRuina/KUHA-REST-API/internal/store/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type CompetitorsHandler struct {
	store fis.Competitors
}

func NewCompetitorsHandler(store fis.Competitors) *CompetitorsHandler {
	return &CompetitorsHandler{store: store}
}

// GetBySector Getter
func (h *CompetitorsHandler) GetAthletesBySector(w http.ResponseWriter, r *http.Request) {
	sectorCode := r.URL.Query().Get("sector")

	if sectorCode == "" {
		utils.BadRequestResponse(w, r, fmt.Errorf("sector code is required"))
		return
	}

	validSectors := map[string]bool{"JP": true, "NK": true, "CC": true}
	if !validSectors[sectorCode] {
		utils.BadRequestResponse(w, r, fmt.Errorf("invalid sector code. Allowed values: JP, NK, CC"))
		return
	}

	competitors, err := h.store.GetAthletesBySector(context.Background(), sectorCode)
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	if len(competitors) == 0 {
		utils.NotFoundResponse(w, r, fmt.Errorf("no competitors found for sector %s", sectorCode))
		return
	}

	utils.WriteJSON(w, http.StatusOK, competitors)
}
