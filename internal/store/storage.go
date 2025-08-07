package store

import (
	"context"
	"database/sql"

	"github.com/DeRuina/KUHA-REST-API/internal/db"
	"github.com/DeRuina/KUHA-REST-API/internal/store/archinisis"
	"github.com/DeRuina/KUHA-REST-API/internal/store/auth"
	"github.com/DeRuina/KUHA-REST-API/internal/store/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/store/kamk"
	"github.com/DeRuina/KUHA-REST-API/internal/store/klab"
	"github.com/DeRuina/KUHA-REST-API/internal/store/tietoevry"
	"github.com/DeRuina/KUHA-REST-API/internal/store/utv"
)

// Define database interfaces
type FIS interface {
	Ping(ctx context.Context) error
	Competitors() fis.Competitors
}

type UTV interface {
	Ping(ctx context.Context) error
	Oura() utv.OuraData
	Polar() utv.PolarData
	Suunto() utv.SuuntoData
	Garmin() utv.GarminData
	PolarToken() utv.PolarToken
	GarminToken() utv.GarminToken
	SuuntoToken() utv.SuuntoToken
	OuraToken() utv.OuraToken
	KlabToken() utv.KlabToken
	Coachtech() utv.CoachtechData
	UserData() utv.UserData
}

type Auth interface {
	Ping(ctx context.Context) error
	IssueToken(ctx context.Context, clientToken, ip, userAgent string) (*auth.Tokens, error)
	RefreshToken(ctx context.Context, refreshToken, ip, userAgent string) (string, error)
}

type Tietoevry interface {
	Ping(ctx context.Context) error
	Users() tietoevry.Users
	Exercises() tietoevry.Exercises
	Symptoms() tietoevry.Symptoms
	Measurements() tietoevry.Measurements
	TestResults() tietoevry.TestResults
	Questionnaires() tietoevry.Questionnaires
	ActivityZones() tietoevry.ActivityZones
}

type KAMK interface {
	Ping(ctx context.Context) error
}

type Klab interface {
	Ping(ctx context.Context) error
	Users() klab.Users
	Data() klab.Data
}

type Archinisis interface {
	Ping(ctx context.Context) error
	Users() archinisis.Users
}

// Storage struct for multiple databases
type Storage struct {
	FIS        FIS
	UTV        UTV
	Auth       Auth
	Tietoevry  Tietoevry
	KAMK       KAMK
	KLAB       Klab
	ARCHINISIS Archinisis
}

// Initializes storage for multiple databases
func NewStorage(databases *db.Database) *Storage {
	var fisStore FIS
	if databases.FIS != nil {
		fisStore = fis.NewFISStorage(databases.FIS)
	}

	var utvStore UTV
	if databases.UTV != nil {
		utvStore = utv.NewUTVStorage(databases.UTV)
	}

	var authStore Auth
	if databases.Auth != nil {
		authStore = auth.NewAuthStorage(databases.Auth)
	}

	var tietoevryStore Tietoevry
	if databases.Tietoevry != nil {
		tietoevryStore = tietoevry.NewTietoevryStorage(databases.Tietoevry)
	}

	var kamkStore KAMK
	if databases.KAMK != nil {
		kamkStore = kamk.NewKAMKStorage(databases.KAMK)
	}

	var klabStore Klab
	if databases.KLAB != nil {
		klabStore = klab.NewKLABStorage(databases.KLAB)
	}

	var archinisisStore Archinisis
	if databases.ARCHINISIS != nil {
		archinisisStore = archinisis.NewArchinisisStorage(databases.ARCHINISIS)
	}

	return &Storage{
		FIS:        fisStore,
		UTV:        utvStore,
		Auth:       authStore,
		Tietoevry:  tietoevryStore,
		KAMK:       kamkStore,
		KLAB:       klabStore,
		ARCHINISIS: archinisisStore,
	}
}

// Initializes storage for only the auth database
func NewAuthOnlyStore(authDB *sql.DB) *Storage {
	return &Storage{
		Auth: auth.NewAuthStorage(authDB),
	}
}
