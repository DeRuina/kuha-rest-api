package utv

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// OuraData interface
type OuraData interface {
	GetDates(ctx context.Context, userID string, startDate *string, endDate *string) ([]string, error)
	GetTypes(ctx context.Context, userID string, summaryDate string) ([]string, error)
	GetData(ctx context.Context, userID string, summaryDate string, key *string) (json.RawMessage, error)
	InsertData(ctx context.Context, userID uuid.UUID, date time.Time, data json.RawMessage) error
	DeleteAllData(ctx context.Context, userID uuid.UUID) (int64, error)
}

// PolarData interface
type PolarData interface {
	GetDates(ctx context.Context, userID string, startDate *string, endDate *string) ([]string, error)
	GetTypes(ctx context.Context, userID string, summaryDate string) ([]string, error)
	GetData(ctx context.Context, userID string, summaryDate string, key *string) (json.RawMessage, error)
	InsertData(ctx context.Context, userID uuid.UUID, date time.Time, data json.RawMessage) error
	DeleteAllData(ctx context.Context, userID uuid.UUID) (int64, error)
}

// SuuntoData interface
type SuuntoData interface {
	GetDates(ctx context.Context, userID string, startDate *string, endDate *string) ([]string, error)
	GetTypes(ctx context.Context, userID string, summaryDate string) ([]string, error)
	GetData(ctx context.Context, userID string, summaryDate string, key *string) (json.RawMessage, error)
	InsertData(ctx context.Context, userID uuid.UUID, date time.Time, data json.RawMessage) error
	DeleteAllData(ctx context.Context, userID uuid.UUID) (int64, error)
}

// GarminData interface
type GarminData interface {
	GetDates(ctx context.Context, userID string, startDate *string, endDate *string) ([]string, error)
	GetTypes(ctx context.Context, userID string, summaryDate string) ([]string, error)
	GetData(ctx context.Context, userID string, summaryDate string, key *string) (json.RawMessage, error)
	InsertData(ctx context.Context, userID uuid.UUID, date time.Time, data json.RawMessage) error
	DeleteAllData(ctx context.Context, userID uuid.UUID) (int64, error)
}

// UTVStorage struct to hold table-specific storage
type UTVStorage struct {
	db     *sql.DB
	oura   OuraData
	polar  PolarData
	suunto SuuntoData
	garmin GarminData
}

// Ping method
func (s *UTVStorage) Ping(ctx context.Context) error {
	return s.db.PingContext(ctx)
}

// Methods to return each table's storage interface
func (s *UTVStorage) Oura() OuraData {
	return s.oura
}

func (s *UTVStorage) Polar() PolarData {
	return s.polar
}

func (s *UTVStorage) Suunto() SuuntoData {
	return s.suunto
}

func (s *UTVStorage) Garmin() GarminData {
	return s.garmin
}

// Storage for UTV database tables
func NewUTVStorage(db *sql.DB) *UTVStorage {
	return &UTVStorage{
		db:     db,
		oura:   &OuraDataStore{db: db},
		polar:  &PolarDataStore{db: db},
		suunto: &SuuntoDataStore{db: db},
		garmin: &GarminDataStore{db: db},
	}
}
