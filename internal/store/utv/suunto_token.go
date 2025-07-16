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

type SuuntoTokenStore struct {
	db *sql.DB
}

func (s *SuuntoTokenStore) GetStatus(ctx context.Context, userID uuid.UUID) (bool, bool, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	row, err := queries.GetSuuntoStatus(ctx, userID)
	if err != nil {
		return false, false, err
	}
	return row.Connected, row.DataExists, nil
}

func (s *SuuntoTokenStore) UpsertToken(ctx context.Context, userID uuid.UUID, data json.RawMessage) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.UpsertSuuntoToken(ctx, utvsqlc.UpsertSuuntoTokenParams{
		UserID: userID,
		Data:   data,
	})
}

func (s *SuuntoTokenStore) GetTokenByUsername(ctx context.Context, username string) (uuid.UUID, json.RawMessage, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	row, err := queries.GetSuuntoTokenByUsername(ctx, username)
	if err != nil {
		return uuid.Nil, nil, err
	}
	return row.UserID, row.Data, nil
}

func (s *SuuntoTokenStore) DeleteToken(ctx context.Context, userID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.DeleteSuuntoToken(ctx, userID)
}

func (s *SuuntoTokenStore) GetTokensForUpdate(ctx context.Context, cutoff time.Time) ([]utvsqlc.SuuntoToken, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.GetSuuntoTokensForUpdate(ctx, cutoff)
}

func (s *SuuntoTokenStore) GetDataForUpdate(ctx context.Context, cutoff time.Time) ([]utvsqlc.SuuntoToken, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.GetSuuntoDataForUpdate(ctx, cutoff)
}
