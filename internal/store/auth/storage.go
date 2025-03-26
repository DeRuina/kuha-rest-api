package auth

import (
	"database/sql"

	authsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/auth"
)

type AuthStorage struct {
	queries *authsqlc.Queries
}

func (a *AuthStorage) Queries() *authsqlc.Queries {
	return a.queries
}

func NewAuthStorage(db *sql.DB) *AuthStorage {
	return &AuthStorage{
		queries: authsqlc.New(db),
	}
}
