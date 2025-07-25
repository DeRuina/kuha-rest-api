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

func (s *UserStore) GetUser(ctx context.Context, id uuid.UUID) (tietoevrysqlc.User, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := tietoevrysqlc.New(s.db)
	return queries.GetUser(ctx, id)
}

func (s *UserStore) LogDeletedUser(ctx context.Context, userID uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := tietoevrysqlc.New(s.db)
	return q.LogDeletedUser(ctx, userID)
}

func (s *UserStore) DeleteUserWithLogging(ctx context.Context, userID uuid.UUID) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	queries := tietoevrysqlc.New(tx)

	if err := queries.LogDeletedUser(ctx, userID); err != nil {
		return 0, err
	}

	rows, err := queries.DeleteUser(ctx, userID)
	if err != nil {
		return 0, err
	}

	if err := tx.Commit(); err != nil {
		return 0, err
	}

	return rows, nil
}

func (s *UserStore) GetDeletedUsers(ctx context.Context) ([]tietoevrysqlc.DeletedUsersLog, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := tietoevrysqlc.New(s.db)
	return q.GetDeletedUsers(ctx)
}
