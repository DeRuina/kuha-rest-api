package utv

import (
	"context"
	"database/sql"
	"encoding/json"

	utvsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/utv"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
	"github.com/google/uuid"
)

type UserDataStore struct {
	db *sql.DB
}

func (s *UserDataStore) GetUserData(ctx context.Context, userID uuid.UUID) (json.RawMessage, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.GetUserData(ctx, userID)
}

func (s *UserDataStore) UpsertUserData(ctx context.Context, userID uuid.UUID, data json.RawMessage) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.UpsertUserData(ctx, utvsqlc.UpsertUserDataParams{
		UserID: userID,
		Data:   data,
	})
}

func (s *UserDataStore) DeleteUserData(ctx context.Context, userID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.DeleteUserData(ctx, userID)
}

func (s *UserDataStore) GetUserIDBySportID(ctx context.Context, sportID string) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)
	return queries.GetUserIDBySportID(ctx, sportID)
}
