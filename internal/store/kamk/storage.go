package kamk

import (
	"context"
	"database/sql"
)

// KAMKStorage
type KAMKStorage struct {
	db *sql.DB
}

// Methods
func (s *KAMKStorage) Ping(ctx context.Context) error {
	return s.db.PingContext(ctx)
}

// NewKAMKStorage creates a new KAMKStorage instance
func NewKAMKStorage(db *sql.DB) *KAMKStorage {
	return &KAMKStorage{
		db: db,
	}
}
