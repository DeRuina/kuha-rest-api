package tietoevry

import (
	"context"
	"database/sql"

	tietoevrysqlc "github.com/DeRuina/KUHA-REST-API/internal/db/tietoevry"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
	"github.com/google/uuid"
)

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) UpsertUser(ctx context.Context, arg tietoevrysqlc.UpsertUserParams) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := tietoevrysqlc.New(s.db)
	return q.UpsertUser(ctx, arg)
}

func (s *UserStore) DeleteUser(ctx context.Context, id uuid.UUID) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := tietoevrysqlc.New(s.db)
	return q.DeleteUser(ctx, id)
}
