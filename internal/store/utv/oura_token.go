package utv

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	utvsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/utv"
	"github.com/google/uuid"
)

type OuraTokenStore struct {
	db *sql.DB
}

func (s *OuraTokenStore) GetStatus(ctx context.Context, userID uuid.UUID) (bool, bool, error) {
	queries := utvsqlc.New(s.db)
	row, err := queries.GetOuraStatus(ctx, userID)
	if err != nil {
		return false, false, err
	}
	return row.Connected, row.DataExists, nil
}

func (s *OuraTokenStore) UpsertToken(ctx context.Context, userID uuid.UUID, data json.RawMessage) error {
	queries := utvsqlc.New(s.db)
	return queries.UpsertOuraToken(ctx, utvsqlc.UpsertOuraTokenParams{
		UserID: userID,
		Data:   data,
	})
}

func (s *OuraTokenStore) GetTokenByOuraID(ctx context.Context, ouraID string) (uuid.UUID, json.RawMessage, error) {
	queries := utvsqlc.New(s.db)
	row, err := queries.GetOuraTokenByOuraID(ctx, ouraID)
	if err != nil {
		return uuid.Nil, nil, err
	}
	return row.UserID, row.Data, nil
}

func (s *OuraTokenStore) DeleteToken(ctx context.Context, userID uuid.UUID) error {
	queries := utvsqlc.New(s.db)
	return queries.DeleteOuraToken(ctx, userID)
}

func (s *OuraTokenStore) GetTokensForUpdate(ctx context.Context, cutoff time.Time) ([]utvsqlc.OuraToken, error) {
	queries := utvsqlc.New(s.db)
	return queries.GetOuraTokensForUpdate(ctx, cutoff)
}
