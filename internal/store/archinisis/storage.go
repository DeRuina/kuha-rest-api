package archinisis

import (
	"context"
	"database/sql"
	"time"
)

const DataTimeout = 30 * time.Second

// Interfaces
type Users interface {
	GetAllSporttiIDs(ctx context.Context) ([]string, error)
}

// ArchinisisStorage
type ArchinisisStorage struct {
	db    *sql.DB
	users Users
}

// Methods
func (s *ArchinisisStorage) Ping(ctx context.Context) error {
	return s.db.PingContext(ctx)
}

func (s *ArchinisisStorage) Users() Users {
	return s.users
}

// NewArchinisisStorage creates a new ArchinisisStorage instance
func NewArchinisisStorage(db *sql.DB) *ArchinisisStorage {
	return &ArchinisisStorage{
		db:    db,
		users: &UsersStore{db: db},
	}
}
