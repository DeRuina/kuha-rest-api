package utv

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

// LatestDataEntry struct to hold latest data entry
type LatestDataEntry struct {
	Device string
	Date   time.Time
	Data   json.RawMessage
}

// OuraData interface
type OuraData interface {
	GetDates(ctx context.Context, userID string, startDate *string, endDate *string) ([]string, error)
	GetTypes(ctx context.Context, userID string, summaryDate string) ([]string, error)
	GetData(ctx context.Context, userID string, summaryDate string, key *string) (json.RawMessage, error)
	InsertData(ctx context.Context, userID uuid.UUID, date time.Time, data json.RawMessage) error
	DeleteAllData(ctx context.Context, userID uuid.UUID) (int64, error)
	GetLatestByType(ctx context.Context, userID uuid.UUID, typ string, limit int32) ([]LatestDataEntry, error)
	GetAllByType(ctx context.Context, userID uuid.UUID, typ string, after, before *time.Time) ([]LatestDataEntry, error)
}

// OuraToken interface
type OuraToken interface {
	GetStatus(ctx context.Context, userID uuid.UUID) (bool, bool, error)
	UpsertToken(ctx context.Context, userID uuid.UUID, data json.RawMessage) error
	GetTokenByOuraID(ctx context.Context, ouraID string) (uuid.UUID, json.RawMessage, error)
}

// PolarData interface
type PolarData interface {
	GetDates(ctx context.Context, userID string, startDate *string, endDate *string) ([]string, error)
	GetTypes(ctx context.Context, userID string, summaryDate string) ([]string, error)
	GetData(ctx context.Context, userID string, summaryDate string, key *string) (json.RawMessage, error)
	InsertData(ctx context.Context, userID uuid.UUID, date time.Time, data json.RawMessage) error
	DeleteAllData(ctx context.Context, userID uuid.UUID) (int64, error)
	GetLatestByType(ctx context.Context, userID uuid.UUID, typ string, limit int32) ([]LatestDataEntry, error)
	GetAllByType(ctx context.Context, userID uuid.UUID, typ string, after, before *time.Time) ([]LatestDataEntry, error)
}

// PolarToken interface
type PolarToken interface {
	GetStatus(ctx context.Context, userID uuid.UUID) (bool, bool, error)
	UpsertToken(ctx context.Context, userID uuid.UUID, data json.RawMessage) error
	GetTokenByPolarID(ctx context.Context, polarID string) (uuid.UUID, json.RawMessage, error)
}

// SuuntoData interface
type SuuntoData interface {
	GetDates(ctx context.Context, userID string, startDate *string, endDate *string) ([]string, error)
	GetTypes(ctx context.Context, userID string, summaryDate string) ([]string, error)
	GetData(ctx context.Context, userID string, summaryDate string, key *string) (json.RawMessage, error)
	InsertData(ctx context.Context, userID uuid.UUID, date time.Time, data json.RawMessage) error
	DeleteAllData(ctx context.Context, userID uuid.UUID) (int64, error)
	GetLatestByType(ctx context.Context, userID uuid.UUID, typ string, limit int32) ([]LatestDataEntry, error)
	GetAllByType(ctx context.Context, userID uuid.UUID, typ string, after, before *time.Time) ([]LatestDataEntry, error)
}

// SuuntoToken interface
type SuuntoToken interface {
	GetStatus(ctx context.Context, userID uuid.UUID) (bool, bool, error)
	UpsertToken(ctx context.Context, userID uuid.UUID, data json.RawMessage) error
	GetTokenByUsername(ctx context.Context, username string) (uuid.UUID, json.RawMessage, error)
}

// GarminData interface
type GarminData interface {
	GetDates(ctx context.Context, userID string, startDate *string, endDate *string) ([]string, error)
	GetTypes(ctx context.Context, userID string, summaryDate string) ([]string, error)
	GetData(ctx context.Context, userID string, summaryDate string, key *string) (json.RawMessage, error)
	InsertData(ctx context.Context, userID uuid.UUID, date time.Time, data json.RawMessage) error
	DeleteAllData(ctx context.Context, userID uuid.UUID) (int64, error)
	GetLatestByType(ctx context.Context, userID uuid.UUID, typ string, limit int32) ([]LatestDataEntry, error)
	GetAllByType(ctx context.Context, userID uuid.UUID, typ string, after, before *time.Time) ([]LatestDataEntry, error)
}

// GarminToken interface
type GarminToken interface {
	GetStatus(ctx context.Context, userID uuid.UUID) (bool, bool, error)
	UpsertToken(ctx context.Context, userID uuid.UUID, data json.RawMessage) error
	TokenExists(ctx context.Context, token string) (bool, error)
	GetUserIDByToken(ctx context.Context, token string) (uuid.UUID, error)
}

// KlabToken interface
type KlabToken interface {
	GetStatus(ctx context.Context, userID uuid.UUID) (bool, error)
	UpsertToken(ctx context.Context, userID uuid.UUID, data json.RawMessage) error
}

// UTVStorage struct to hold table-specific storage
type UTVStorage struct {
	db          *sql.DB
	oura        OuraData
	polar       PolarData
	suunto      SuuntoData
	garmin      GarminData
	polarToken  PolarToken
	garminToken GarminToken
	suuntoToken SuuntoToken
	ouraToken   OuraToken
	klabToken   KlabToken
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

func (s *UTVStorage) PolarToken() PolarToken {
	return s.polarToken
}

func (s *UTVStorage) GarminToken() GarminToken {
	return s.garminToken
}

func (s *UTVStorage) OuraToken() OuraToken {
	return s.ouraToken
}

func (s *UTVStorage) SuuntoToken() SuuntoToken {
	return s.suuntoToken
}

func (s *UTVStorage) KlabToken() KlabToken {
	return s.klabToken
}

// Storage for UTV database tables
func NewUTVStorage(db *sql.DB) *UTVStorage {
	return &UTVStorage{
		db:          db,
		oura:        &OuraDataStore{db: db},
		polar:       &PolarDataStore{db: db},
		suunto:      &SuuntoDataStore{db: db},
		garmin:      &GarminDataStore{db: db},
		polarToken:  &PolarTokenStore{db: db},
		garminToken: &GarminTokenStore{db: db},
		suuntoToken: &SuuntoTokenStore{db: db},
		ouraToken:   &OuraTokenStore{db: db},
		klabToken:   &KlabTokenStore{db: db},
	}
}
