package archapi

import (
	"database/sql"
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

type SporttiIDParam struct {
	SporttiID string `form:"sportti_id" validate:"required,numeric"`
}

// DeleteUser godoc
//
//	@Summary		Delete an athlete (hard delete)
//	@Description	Removes an athlete by sportti_id. Related measurements and reports are deleted via FK cascades.
//	@Tags			Archinisis - User
//	@Accept			json
//	@Produce		json
//	@Param			sportti_id	query	string	true	"Sportti ID"
//	@Success		200
//	@Failure		400	{object}	swagger.ValidationErrorResponse
//	@Failure		401	{object}	swagger.UnauthorizedResponse
//	@Failure		403	{object}	swagger.ForbiddenResponse
//	@Failure		404	{object}	swagger.NotFoundResponse
//	@Failure		500	{object}	swagger.InternalServerErrorResponse
//	@Failure		503	{object}	swagger.ServiceUnavailableResponse
//	@Security		BearerAuth
//	@Router			/archinisis/user [delete]
func (h *UserDataHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	if !authz.Authorize(r) {
		utils.ForbiddenResponse(w, r, fmt.Errorf("access denied"))
		return
	}

	if err := utils.ValidateParams(r, []string{"sportti_id"}); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	params := SporttiIDParam{
		SporttiID: r.URL.Query().Get("sportti_id"),
	}
	if err := utils.GetValidator().Struct(params); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	parsed, err := utils.ParseSporttiID(params.SporttiID)
	if err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	_, err = h.store.DeleteUserBySporttiID(r.Context(), parsed)
	if err == sql.ErrNoRows {
		utils.NotFoundResponse(w, r, err)
		return
	}
	if err != nil {
		utils.InternalServerError(w, r, err)
		return
	}

	invalidateArchAll(r.Context(), h.cache, parsed)

	w.WriteHeader(http.StatusOK)
}
