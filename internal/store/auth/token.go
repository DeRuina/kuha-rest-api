package auth

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"errors"
	"time"

	"github.com/DeRuina/KUHA-REST-API/internal/auth/authn"
	authsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/auth"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
	"github.com/sqlc-dev/pqtype"
)

type Tokens struct {
	JWT          string
	RefreshToken string
}

func (a *AuthStorage) IssueToken(ctx context.Context, clientTokenRaw, ip, userAgent string) (*Tokens, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	hashed := sha256.Sum256([]byte(clientTokenRaw))
	clientToken := hex.EncodeToString(hashed[:])

	revoked, err := a.queries.IsRevokedToken(ctx, clientToken)
	if err != nil {
		return nil, err
	}
	if revoked {
		return nil, errors.New("client_token is revoked")
	}

	client, err := a.queries.GetClientByToken(ctx, clientToken)
	if err != nil {
		return nil, errors.New("invalid client_token")
	}

	// Start a transaction to ensure atomicity
	tx, err := a.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback() // This will be a no-op if the transaction is committed

	queries := authsqlc.New(tx)

	// Check and revoke old refresh token
	old, err := queries.GetRefreshTokenByClient(ctx, clientToken)
	if err == nil {
		if err := queries.DeleteRefreshToken(ctx, old.Token); err != nil {
			return nil, err
		}
		if err := queries.CreateRevokedRefreshToken(ctx, old.Token); err != nil {
			return nil, err
		}

		meta := pqtype.NullRawMessage{Valid: true}
		if err := meta.Scan([]byte(`{"reason":"rotation"}`)); err != nil {
			return nil, err
		}

		if err := queries.InsertTokenLog(ctx, authsqlc.InsertTokenLogParams{
			ClientToken: clientToken,
			TokenType:   "refresh",
			Action:      "revoked",
			Token:       utils.NullString(old.Token),
			IpAddress:   sql.NullString{String: ip, Valid: true},
			UserAgent:   sql.NullString{String: userAgent, Valid: true},
			Metadata:    meta,
		}); err != nil {
			return nil, err
		}
	}

	// Generate new refresh
	refresh, err := authn.GenerateRandomToken()
	if err != nil {
		return nil, err
	}
	expires := time.Now().Add(90 * 24 * time.Hour)
	if err := queries.CreateRefreshToken(ctx, authsqlc.CreateRefreshTokenParams{
		ClientToken: clientToken,
		Token:       refresh,
		ExpiresAt:   expires,
	}); err != nil {
		return nil, err
	}

	// Generate JWT
	jwt, err := authn.GenerateJWT(client.ClientName, client.Role, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	metaRefresh := pqtype.NullRawMessage{Valid: true}
	if err := metaRefresh.Scan([]byte(`{"reason":"new refresh"}`)); err != nil {
		return nil, err
	}

	metaJWT := pqtype.NullRawMessage{Valid: true}
	if err := metaJWT.Scan([]byte(`{"reason":"new jwt"}`)); err != nil {
		return nil, err
	}

	if err := queries.InsertTokenLog(ctx, authsqlc.InsertTokenLogParams{
		ClientToken: clientToken,
		TokenType:   "refresh",
		Action:      "issued",
		Token:       utils.NullString(refresh),
		IpAddress:   sql.NullString{String: ip, Valid: true},
		UserAgent:   sql.NullString{String: userAgent, Valid: true},
		Metadata:    metaRefresh,
	}); err != nil {
		return nil, err
	}

	if err := queries.InsertTokenLog(ctx, authsqlc.InsertTokenLogParams{
		ClientToken: clientToken,
		TokenType:   "jwt",
		Action:      "issued",
		Token:       utils.NullString(refresh),
		IpAddress:   sql.NullString{String: ip, Valid: true},
		UserAgent:   sql.NullString{String: userAgent, Valid: true},
		Metadata:    metaJWT,
	}); err != nil {
		return nil, err
	}

	// Commit the transaction if all operations succeeded
	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &Tokens{
		JWT:          jwt,
		RefreshToken: refresh,
	}, nil
}
