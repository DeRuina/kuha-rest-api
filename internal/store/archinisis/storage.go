package archinisis

import (
	"context"
	"database/sql"
	"time"
)

const DataTimeout = 30 * time.Second

// Interfaces
type Users interface {
	GetAllSporttiIDs(ctx context.Context) ([]string, error)
}

type Data interface {
	GetRaceReportSessions(ctx context.Context, sporttiID string) ([]int32, error)
	GetRaceReport(ctx context.Context, sporttiID string, sessionID int32) (string, error)
}

// ArchinisisStorage
type ArchinisisStorage struct {
	db    *sql.DB
	users Users
	data  Data
}

// Methods
func (s *ArchinisisStorage) Ping(ctx context.Context) error {
	return s.db.PingContext(ctx)
}

func (s *ArchinisisStorage) Users() Users {
	return s.users
}

func (s *ArchinisisStorage) Data() Data {
	return s.data
}

// NewArchinisisStorage creates a new ArchinisisStorage instance
func NewArchinisisStorage(db *sql.DB) *ArchinisisStorage {
	return &ArchinisisStorage{
		db:    db,
		users: &UsersStore{db: db},
		data:  &DataStore{db: db},
	}
}
