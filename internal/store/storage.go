package store

import (
	"github.com/DeRuina/KUHA-REST-API/internal/db"
	authsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/auth"
	"github.com/DeRuina/KUHA-REST-API/internal/store/auth"
	"github.com/DeRuina/KUHA-REST-API/internal/store/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/store/utv"
)

// Define database interfaces
type FIS interface {
	Competitors() fis.Competitors
}

type UTV interface {
	Oura() utv.OuraData
	Polar() utv.PolarData
	Suunto() utv.SuuntoData
}

type Auth interface {
	Queries() *authsqlc.Queries
}

// Storage struct for multiple databases
type Storage struct {
	FIS  FIS
	UTV  UTV
	Auth Auth
}

// Initializes storage for multiple databases
func NewStorage(databases *db.Database) *Storage {
	return &Storage{
		FIS:  fis.NewFISStorage(databases.FIS),
		UTV:  utv.NewUTVStorage(databases.UTV),
		Auth: auth.NewAuthStorage(databases.Auth),
	}
}
