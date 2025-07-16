package klab

import (
	"context"
	"database/sql"
)

// kLABStorage
type KLABStorage struct {
	db *sql.DB
}

// Methods
func (s *KLABStorage) Ping(ctx context.Context) error {
	return s.db.PingContext(ctx)
}

// NewKLABStorage creates a new KLABStorage instance
func NewKLABStorage(db *sql.DB) *KLABStorage {
	return &KLABStorage{
		db: db,
	}
}
