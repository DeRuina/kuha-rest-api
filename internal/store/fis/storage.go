package fis

import (
	"context"
	"database/sql"
)

// Competitors interface
type Competitors interface {
	GetAthletesBySector(ctx context.Context, sectorCode string) ([]GetBySectorResponse, error)
	GetNationsBySector(ctx context.Context, sectorCode string) ([]string, error)
}

// FISStorage struct to hold table-specific storage
type FISStorage struct {
	db          *sql.DB
	competitors Competitors
}

// Ping method
func (s *FISStorage) Ping(ctx context.Context) error {
	return s.db.PingContext(ctx)
}

// Methods to return each table's storage interface
func (s *FISStorage) Competitors() Competitors {
	return s.competitors
}

// Storage for FIS database tables
func NewFISStorage(db *sql.DB) *FISStorage {
	return &FISStorage{
		db:          db,
		competitors: &CompetitorsStore{db: db},
	}
}
