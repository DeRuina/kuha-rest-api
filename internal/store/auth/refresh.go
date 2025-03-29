package auth

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authn"
	authsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/auth"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
	"github.com/sqlc-dev/pqtype"
)

func (a *AuthStorage) RefreshToken(ctx context.Context, refreshToken, ip, userAgent string) (string, error) {
	revoked, err := a.queries.IsRevokedRefreshToken(ctx, refreshToken)
	if err != nil {
		return "", err
	}
	if revoked {
		return "", errors.New("refresh token revoked")
	}

	tokenData, err := a.queries.GetRefreshToken(ctx, refreshToken)
	if err != nil {
		return "", errors.New("invalid refresh token")
	}
	if tokenData.ExpiresAt.Before(time.Now()) {
		return "", errors.New("refresh token expired")
	}

	client, err := a.queries.GetClientByToken(ctx, tokenData.ClientToken)
	if err != nil {
		return "", errors.New("client not found")
	}

	jwt, err := authn.GenerateJWT(client.ClientName, client.Role, 24*time.Hour)
	if err != nil {
		return "", err
	}

	metaRefresh := pqtype.NullRawMessage{Valid: true}
	if err := metaRefresh.Scan([]byte(`{"reason":"used refresh"}`)); err != nil {
		return "", err
	}

	metaJWT := pqtype.NullRawMessage{Valid: true}
	if err := metaJWT.Scan([]byte(`{"reason":"new jwt"}`)); err != nil {
		return "", err
	}

	if err := a.queries.InsertTokenLog(ctx, authsqlc.InsertTokenLogParams{
		ClientToken: tokenData.ClientToken,
		TokenType:   "refresh",
		Action:      "used",
		Token:       utils.NullString(refreshToken),
		IpAddress:   sql.NullString{String: ip, Valid: true},
		UserAgent:   sql.NullString{String: userAgent, Valid: true},
		Metadata:    metaRefresh,
	}); err != nil {
		return "", err
	}

	if err := a.queries.InsertTokenLog(ctx, authsqlc.InsertTokenLogParams{
		ClientToken: tokenData.ClientToken,
		TokenType:   "jwt",
		Action:      "issued",
		Token:       utils.NullString(refreshToken),
		IpAddress:   sql.NullString{String: ip, Valid: true},
		UserAgent:   sql.NullString{String: userAgent, Valid: true},
		Metadata:    metaJWT,
	}); err != nil {
		return "", err
	}

	return jwt, nil
}
