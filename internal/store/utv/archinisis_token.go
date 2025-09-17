package utv

import (
	"context"
	"database/sql"
	"encoding/json"

	utvsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/utv"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
	"github.com/google/uuid"
)

type ArchinisisTokenStore struct {
	db *sql.DB
}

func (s *ArchinisisTokenStore) GetStatus(ctx context.Context, userID uuid.UUID) (bool, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.GetArchinisisStatus(ctx, userID)
}

func (s *ArchinisisTokenStore) UpsertToken(ctx context.Context, userID uuid.UUID, data json.RawMessage) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.UpsertArchinisisToken(ctx, utvsqlc.UpsertArchinisisTokenParams{
		UserID: userID,
		Data:   data,
	})
}

func (s *ArchinisisTokenStore) DeleteToken(ctx context.Context, userID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.DeleteArchinisisToken(ctx, userID)
}

func (s *ArchinisisTokenStore) GetSportIDs(ctx context.Context) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := utvsqlc.New(s.db)
	return q.GetArchinisisSportIDs(ctx)
}
