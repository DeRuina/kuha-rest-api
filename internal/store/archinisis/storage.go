package archinisis

import (
	"context"
	"database/sql"
)

// ArchinisisStorage
type ArchinisisStorage struct {
	db *sql.DB
}

// Methods
func (s *ArchinisisStorage) Ping(ctx context.Context) error {
	return s.db.PingContext(ctx)
}

// NewArchinisisStorage creates a new ArchinisisStorage instance
func NewArchinisisStorage(db *sql.DB) *ArchinisisStorage {
	return &ArchinisisStorage{
		db: db,
	}
}
