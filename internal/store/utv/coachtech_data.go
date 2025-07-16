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

type CoachtechDataStore struct {
	db *sql.DB
}

func (s *CoachtechDataStore) GetStatus(ctx context.Context, userID uuid.UUID) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.GetCoachtechStatus(ctx, userID)
}

func (s *CoachtechDataStore) GetData(ctx context.Context, userID uuid.UUID, after, before *time.Time) ([]json.RawMessage, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	arg := utvsqlc.GetCoachtechDataParams{
		UserID:     userID,
		AfterDate:  utils.NullTimeIfEmpty(after),
		BeforeDate: utils.NullTimeIfEmpty(before),
	}

	queries := utvsqlc.New(s.db)

	data, err := queries.GetCoachtechData(ctx, arg)
	if err != nil {
		return nil, err
	}

	return data, nil
}
