package utv

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	utvsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/utv"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
	"github.com/google/uuid"
)

type PolarTokenStore struct {
	db *sql.DB
}

func (s *PolarTokenStore) GetStatus(ctx context.Context, userID uuid.UUID) (bool, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	row, err := queries.GetPolarStatus(ctx, userID)
	if err != nil {
		return false, false, err
	}
	return row.Connected, row.DataExists, nil
}

func (s *PolarTokenStore) UpsertToken(ctx context.Context, userID uuid.UUID, data json.RawMessage) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.UpsertPolarToken(ctx, utvsqlc.UpsertPolarTokenParams{
		UserID: userID,
		Data:   data,
	})
}

func (s *PolarTokenStore) GetTokenByPolarID(ctx context.Context, polarID string) (uuid.UUID, json.RawMessage, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	row, err := queries.GetPolarTokenByPolarID(ctx, polarID)
	if err != nil {
		return uuid.Nil, nil, err
	}
	return row.UserID, row.Data, nil
}

func (s *PolarTokenStore) DeleteToken(ctx context.Context, userID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.DeletePolarToken(ctx, userID)
}

func (s *PolarTokenStore) GetTokensForUpdate(ctx context.Context, cutoff time.Time) ([]utvsqlc.PolarToken, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.GetPolarTokensForUpdate(ctx, cutoff)
}

func (s *PolarTokenStore) GetDataForUpdate(ctx context.Context, cutoff time.Time) ([]utvsqlc.PolarToken, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.GetPolarDataForUpdate(ctx, cutoff)
}

func (s *PolarTokenStore) GetTokenJSON(ctx context.Context, userID uuid.UUID) (json.RawMessage, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.GetPolarTokenJSON(ctx, userID)
}
