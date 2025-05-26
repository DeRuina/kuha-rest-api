package auth

import (
	"context"
	"database/sql"

	authsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/auth"
)

type AuthStorage struct {
	db      *sql.DB
	queries *authsqlc.Queries
}

func (a *AuthStorage) Queries() *authsqlc.Queries {
	return a.queries
}

func (s *AuthStorage) Ping(ctx context.Context) error {
	return s.db.PingContext(ctx)
}

func NewAuthStorage(db *sql.DB) *AuthStorage {
	return &AuthStorage{
		db:      db,
		queries: authsqlc.New(db),
	}
}
