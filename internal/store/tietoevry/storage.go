package tietoevry

import (
	"context"
	"database/sql"

	tietoevrysqlc "github.com/DeRuina/KUHA-REST-API/internal/db/tietoevry"
	"github.com/google/uuid"
)

// Interface

type Users interface {
	UpsertUser(ctx context.Context, arg tietoevrysqlc.UpsertUserParams) error
	DeleteUser(ctx context.Context, id uuid.UUID) (int64, error)
}

// TietoevryStorage
type TietoevryStorage struct {
	db    *sql.DB
	users Users
}

// Methods
func (s *TietoevryStorage) Ping(ctx context.Context) error {
	return s.db.PingContext(ctx)
}

func (s *TietoevryStorage) Users() Users {
	return s.users
}

// NewTietoevryStorage creates a new TietoevryStorage instance
func NewTietoevryStorage(db *sql.DB) *TietoevryStorage {
	return &TietoevryStorage{
		db:    db,
		users: &UserStore{db: db},
	}
}
