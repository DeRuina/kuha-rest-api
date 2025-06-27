package utv

import (
	"context"
	"database/sql"
	"encoding/json"

	utvsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/utv"
	"github.com/google/uuid"
)

type PolarTokenStore struct {
	db *sql.DB
}

func (s *PolarTokenStore) GetStatus(ctx context.Context, userID uuid.UUID) (bool, bool, error) {
	queries := utvsqlc.New(s.db)
	row, err := queries.GetPolarStatus(ctx, userID)
	if err != nil {
		return false, false, err
	}
	return row.Connected, row.DataExists, nil
}

func (s *PolarTokenStore) UpsertToken(ctx context.Context, userID uuid.UUID, data json.RawMessage) error {
	queries := utvsqlc.New(s.db)
	return queries.UpsertPolarToken(ctx, utvsqlc.UpsertPolarTokenParams{
		UserID: userID,
		Data:   data,
	})
}

func (s *PolarTokenStore) GetTokenByPolarID(ctx context.Context, polarID string) (uuid.UUID, json.RawMessage, error) {
	queries := utvsqlc.New(s.db)
	row, err := queries.GetPolarTokenByPolarID(ctx, polarID)
	if err != nil {
		return uuid.Nil, nil, err
	}
	return row.UserID, row.Data, nil
}
