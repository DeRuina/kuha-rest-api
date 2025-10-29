package fis

import (
	"context"
	"database/sql"

	fissqlc "github.com/DeRuina/KUHA-REST-API/internal/db/fis"
)

// Competitors interface
type Competitors interface {
	GetAthletesBySector(ctx context.Context, sector string) ([]AthleteRow, error)
	GetNationsBySector(ctx context.Context, sector string) ([]string, error)

	GetLastRowCompetitor(ctx context.Context) (fissqlc.ACompetitor, error)
	InsertCompetitor(ctx context.Context, in InsertCompetitorClean) error
	UpdateCompetitorByID(ctx context.Context, in UpdateCompetitorClean) error
	DeleteCompetitorByID(ctx context.Context, competitorID int32) error

	GetCompetitorIDByFiscodeCC(ctx context.Context, fiscode int32) (int32, error)
	GetCompetitorIDByFiscodeJP(ctx context.Context, fiscode int32) (int32, error)
	GetCompetitorIDByFiscodeNK(ctx context.Context, fiscode int32) (int32, error)
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
