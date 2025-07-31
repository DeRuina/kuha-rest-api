package authapi

import (
	"encoding/json"
	"net/http"

	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type TokenRequest struct {
	ClientToken string `json:"client_token" validate:"required"`
}

type TokenResponse struct {
	JWT          string `json:"jwt"`
	RefreshToken string `json:"refresh_token"`
}

// IssueTokens godoc
//
//	@Summary		Issue JWT and Refresh token
//	@Description	Authenticates the client_token and returns a JWT and refresh token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			client_token	body		TokenRequest	true	"Client Token"
//	@Success		200				{object}	TokenResponse	"Tokens"
//	@Failure		400				{object}	swagger.ValidationErrorResponse
//	@Failure		401				{object}	swagger.UnauthorizedResponse
//	@Failure		500				{object}	swagger.InternalServerErrorResponse
//	@Failure		503				{object}	swagger.ServiceUnavailableResponse
//	@Router			/auth/token [post]
func (h *AuthHandler) IssueTokens(w http.ResponseWriter, r *http.Request) {
	var req TokenRequest

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

	tokens, err := h.store.IssueToken(r.Context(), req.ClientToken, ip, userAgent)
	if err != nil {
		utils.HandleDatabaseError(w, r, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"tokens": TokenResponse{
			JWT:          tokens.JWT,
			RefreshToken: tokens.RefreshToken,
		},
	})
}
