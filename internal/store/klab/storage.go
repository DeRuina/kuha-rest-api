package klab

import (
	"context"
	"database/sql"
	"time"

	klabsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/klab"
)

const DataTimeout = 30 * time.Second

// Interfaces
type Users interface {
	GetAllSporttiIDs(ctx context.Context) ([]string, error)
	GetCustomerByID(ctx context.Context, idcustomer int32) (klabsqlc.Customer, error)
	GetCustomerIDBySporttiID(ctx context.Context, sporttiID string) (int32, error)
}

type Data interface {
	InsertKlabDataBulk(ctx context.Context, payloads []KlabDataPayload) error
	GetDataByCustomerIDNoCustomer(ctx context.Context, idcustomer int32) (*KlabDataNoCustomerResponse, error)
	GetCustomerIDBySporttiID(ctx context.Context, sporttiID string) (int32, error)
}

// kLABStorage
type KLABStorage struct {
	db    *sql.DB
	users Users
	data  Data
}

// Methods
func (s *KLABStorage) Ping(ctx context.Context) error {
	return s.db.PingContext(ctx)
}

func (s *KLABStorage) Users() Users {
	return s.users
}

func (s *KLABStorage) Data() Data {
	return s.data
}

// NewKLABStorage creates a new KLABStorage instance
func NewKLABStorage(db *sql.DB) *KLABStorage {
	return &KLABStorage{
		db:    db,
		users: &UsersStore{db: db},
		data:  &DataStore{db: db},
	}
}
