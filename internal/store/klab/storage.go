package klab

import (
	"context"
	"database/sql"
)

// kLABStorage
type kLABStorage struct {
	db *sql.DB
}

// Methods
func (s *kLABStorage) Ping(ctx context.Context) error {
	return s.db.PingContext(ctx)
}

// NewkLABStorage creates a new kLABStorage instance
func NewkLABStorage(db *sql.DB) *kLABStorage {
	return &kLABStorage{
		db: db,
	}
}
