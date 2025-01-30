package fis

import (
	"context"
	"database/sql"
)

// Define Competitors interface
type Competitors interface {
	GetBySector(ctx context.Context, sectorCode string) ([]Competitor, error)
	GetByFiscodeJP(ctx context.Context, fiscode int32) (int32, error)
	GetByFiscodeNK(ctx context.Context, fiscode int32) (int32, error)
	GetByGenderAndNationJP(ctx context.Context, gender, nation string) ([]int32, error)
}

// Ensure `FISStorage` implements `FIS`
type FISStorage struct {
	competitors Competitors // ✅ Private field holding the actual implementation
}

// Implement the `Competitors()` method to return the interface
func (s *FISStorage) Competitors() Competitors {
	return s.competitors
}

// `NewFISStorage` initializes storage for FIS database tables
func NewFISStorage(db *sql.DB) *FISStorage {
	return &FISStorage{
		competitors: &CompetitorsStore{db}, // ✅ Correctly assign CompetitorsStore
	}
}
