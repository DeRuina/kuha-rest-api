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

type GarminTokenStore struct {
	db *sql.DB
}

func (s *GarminTokenStore) GetStatus(ctx context.Context, userID uuid.UUID) (bool, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	row, err := queries.GetGarminStatus(ctx, userID)
	if err != nil {
		return false, false, err
	}

	return row.Connected, row.DataExists, nil
}

func (s *GarminTokenStore) UpsertToken(ctx context.Context, userID uuid.UUID, data json.RawMessage) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.UpsertGarminToken(ctx, utvsqlc.UpsertGarminTokenParams{
		UserID: userID,
		Data:   data,
	})
}

func (s *GarminTokenStore) TokenExists(ctx context.Context, token string) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.GarminTokenExists(ctx, token)
}

func (s *GarminTokenStore) GetUserIDByToken(ctx context.Context, token string) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.GetGarminUserIDByToken(ctx, token)
}

func (s *GarminTokenStore) DeleteToken(ctx context.Context, userID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.DeleteGarminToken(ctx, userID)
}

func (s *GarminTokenStore) GetTokensForUpdate(ctx context.Context, cutoff time.Time) ([]utvsqlc.GarminToken, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.GetGarminTokensForUpdate(ctx, cutoff)
}

func (s *GarminTokenStore) GetDataForUpdate(ctx context.Context, cutoff time.Time) ([]utvsqlc.GarminToken, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.GetGarminDataForUpdate(ctx, cutoff)
}

func (s *GarminTokenStore) GetTokenJSON(ctx context.Context, userID uuid.UUID) (json.RawMessage, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.GetGarminTokenJSON(ctx, userID)
}
