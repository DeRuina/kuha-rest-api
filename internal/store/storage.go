package store

import (
	"database/sql"

	"github.com/DeRuina/KUHA-REST-API/internal/store/fis"
)

// Define FIS interface
type FIS interface {
	Competitors() fis.Competitors
}

// Define Storage struct for multiple databases
type Storage struct {
	FIS FIS
}

// NewStorage initializes storage for multiple databases
func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		FIS: fis.NewFISStorage(db),
	}
}
