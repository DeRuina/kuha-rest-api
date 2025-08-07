package archapi

import (
	"fmt"
	"net/http"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authz"
	"github.com/DeRuina/KUHA-REST-API/internal/store/archinisis"
	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type UserDataHandler struct {
	store archinisis.Users
	cache *cache.Storage
}

func NewUserDataHandler(store archinisis.Users, cache *cache.Storage) *UserDataHandler {
	return &UserDataHandler{store: store, cache: cache}
}

// GetSporttiIDs godoc
//
//	@Summary		Get all Sportti IDs
//	@Description	Returns a list of all Sportti IDs
//	@Tags			ARCHINISIS - User
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	swagger.UserDataArchinisisResponse
//	@Failure		401	{object}	swagger.UnauthorizedResponse
//	@Failure		403	{object}	swagger.ForbiddenResponse
//	@Failure		500	{object}	swagger.InternalServerErrorResponse
//	@Failure		503	{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/archinisis/sport-ids [get]
func (h *UserDataHandler) GetSporttiIDs(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	sporttiIDs, err := h.store.GetAllSporttiIDs(r.Context())
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]any{
		"sportti_ids": sporttiIDs,
	})
}
