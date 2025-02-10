package utv

import (
	"context"
	"database/sql"
)

// OuraData interface
type OuraData interface {
	GetDates(ctx context.Context, userID string, startDate *string, endDate *string) ([]string, error)
	GetTypes(ctx context.Context, userID string, summaryDate string) ([]string, error)
	// GetData(ctx context.Context, userID string, summaryDate string, key *string) (interface{}, error)
}

// type PolarData interface {
// 	GetDates(ctx context.Context, userID string, startDate string, endDate string) ([]string, error)
// 	GetTypes(ctx context.Context, userID string, summaryDate string) ([]string, error)
// 	GetDataPoint(ctx context.Context, userID string, summaryDate string, key string) (interface{}, error)
// 	GetUniqueTypes(ctx context.Context, userID string, startDate string, endDate string) ([]string, error)
// }

// type SuuntoData interface {
// 	GetDates(ctx context.Context, userID string, startDate string, endDate string) ([]string, error)
// 	GetTypes(ctx context.Context, userID string, summaryDate string) ([]string, error)
// 	GetDataPoint(ctx context.Context, userID string, summaryDate string, key string) (interface{}, error)
// 	GetUniqueTypes(ctx context.Context, userID string, startDate string, endDate string) ([]string, error)
// }

// UTVStorage struct to hold table-specific storage
type UTVStorage struct {
	oura OuraData
	// polar  PolarData
	// suunto SuuntoData
}

// Implement methods to return each table's storage interface
func (s *UTVStorage) Oura() OuraData {
	return s.oura
}

// func (s *UTVStorage) Polar() PolarData {
// 	return s.polar
// }

// func (s *UTVStorage) Suunto() SuuntoData {
// 	return s.suunto
// }

// `NewUTVStorage` initializes storage for UTV database tables
func NewUTVStorage(db *sql.DB) *UTVStorage {
	return &UTVStorage{
		oura: &OuraDataStore{db: db},
		// polar:  &PolarDataStore{db: db},
		// suunto: &SuuntoDataStore{db: db},
	}
}
