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

func (h *AuthHandler) IssueToken(w http.ResponseWriter, r *http.Request) {
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
		utils.UnauthorizedResponse(w, r, err)
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]interface{}{
		"tokens": TokenResponse{
			JWT:          tokens.JWT,
			RefreshToken: tokens.RefreshToken,
		},
	})
}
