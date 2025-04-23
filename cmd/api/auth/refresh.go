package authapi

import (
	"encoding/json"
	"net/http"

	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshResponse struct {
	JWT string `json:"jwt"`
}

// RefreshToken godoc
//
//	@Summary		Issue a new JWT token
//	@Description	Authenticates the refresh token and returns a new JWT token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			refresh_token	body		RefreshRequest	true	"Refresh Token"
//	@Success		200				{object}	RefreshResponse	"New JWT token"
//	@Failure		400				{object}	swagger.ValidationErrorResponse
//	@Failure		401				{object}	swagger.UnauthorizedResponse
//	@Failure		500				{object}	swagger.InternalServerErrorResponse
//	@Router			/auth/refresh [post]
func (h *AuthHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	var req RefreshRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	if err := utils.GetValidator().Struct(req); err != nil {
		utils.BadRequestResponse(w, r, err)
		return
	}

	ip := r.RemoteAddr
	userAgent := r.UserAgent()

	jwt, err := h.store.RefreshToken(r.Context(), req.RefreshToken, ip, userAgent)
	if err != nil {
		utils.UnauthorizedResponse(w, r, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"tokens": RefreshResponse{JWT: jwt},
	})
}
