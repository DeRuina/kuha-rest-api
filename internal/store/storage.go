package store

import (
	"github.com/DeRuina/KUHA-REST-API/internal/db"
	"github.com/DeRuina/KUHA-REST-API/internal/store/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/store/utv"
)

// Define database interfaces
type FIS interface {
	Competitors() fis.Competitors
}

type UTV interface {
	Oura() utv.OuraData
}

// Define Storage struct for multiple databases
type Storage struct {
	FIS FIS
	UTV UTV
}

// NewStorage initializes storage for multiple databases
func NewStorage(databases *db.Database) *Storage {
	return &Storage{
		FIS: fis.NewFISStorage(databases.FIS),
		UTV: utv.NewUTVStorage(databases.UTV),
	}
}
