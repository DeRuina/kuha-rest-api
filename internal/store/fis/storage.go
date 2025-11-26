package fis

import (
	"context"
	"database/sql"

	fissqlc "github.com/DeRuina/KUHA-REST-API/internal/db/fis"
)

// Resultcc interface
type Resultcc interface {
	GetLastRowResultCC(ctx context.Context) (fissqlc.AResultcc, error)
	InsertResultCC(ctx context.Context, in InsertResultCCClean) error
	UpdateResultCCByRecID(ctx context.Context, in UpdateResultCCClean) error
	DeleteResultCCByRecID(ctx context.Context, recid int32) error
	GetRaceResultsCCByRaceID(ctx context.Context, raceID int32) ([]fissqlc.AResultcc, error)
	GetAthleteResultsCC(ctx context.Context, competitorID int32, seasons []int32, disciplines, cats []string) ([]fissqlc.GetAthleteResultsCCRow, error)
}

// Resultjp interface
type Resultjp interface {
	GetLastRowResultJP(ctx context.Context) (fissqlc.AResultjp, error)
	InsertResultJP(ctx context.Context, in InsertResultJPClean) error
	UpdateResultJPByRecID(ctx context.Context, in UpdateResultJPClean) error
	DeleteResultJPByRecID(ctx context.Context, recid int32) error
	GetRaceResultsJPByRaceID(ctx context.Context, raceID int32) ([]fissqlc.AResultjp, error)
	GetAthleteResultsJP(ctx context.Context, competitorID int32, seasons []int32, disciplines, cats []string) ([]fissqlc.GetAthleteResultsJPRow, error)
}

// Resultnk interface
type Resultnk interface {
	GetLastRowResultNK(ctx context.Context) (fissqlc.AResultnk, error)
	InsertResultNK(ctx context.Context, in InsertResultNKClean) error
	UpdateResultNKByRecID(ctx context.Context, in UpdateResultNKClean) error
	DeleteResultNKByRecID(ctx context.Context, recid int32) error
	GetRaceResultsNKByRaceID(ctx context.Context, raceID int32) ([]fissqlc.AResultnk, error)
	GetAthleteResultsNK(ctx context.Context, competitorID int32, seasons []int32, disciplines, cats []string) ([]fissqlc.GetAthleteResultsNKRow, error)
}

// Racecc interface
type Racecc interface {
	GetCrossCountrySeasons(ctx context.Context) ([]int32, error)
	GetCrossCountryDisciplines(ctx context.Context) ([]string, error)
	GetCrossCountryCategories(ctx context.Context) ([]string, error)
	GetRacesCC(ctx context.Context, seasons []int32, disciplines, cats []string) ([]fissqlc.ARacecc, error)
	GetLastRowRaceCC(ctx context.Context) (fissqlc.ARacecc, error)
	InsertRaceCC(ctx context.Context, in InsertRaceCCClean) error
	UpdateRaceCCByID(ctx context.Context, in UpdateRaceCCClean) error
	DeleteRaceCCByID(ctx context.Context, raceID int32) error
	SearchRacesCC(ctx context.Context, seasoncode *int32, nationcode, gender, catcode *string) ([]fissqlc.SearchRacesCCRow, error)
	GetRacesByIDsCC(ctx context.Context, raceIDs []int32) ([]fissqlc.ARacecc, error)
	GetRaceCountsByCategoryCC(ctx context.Context, seasoncode int32, nationcode, gender *string) ([]fissqlc.GetRaceCountsByCategoryCCRow, error)
	GetRaceCountsByNationCC(ctx context.Context, seasoncode int32, catcode, gender *string) ([]fissqlc.GetRaceCountsByNationCCRow, error)
	GetRaceTotalCC(ctx context.Context, seasoncode int32, catcode, gender *string) (int64, error)
}

// Racejp interface
type Racejp interface {
	GetSkiJumpingSeasons(ctx context.Context) ([]int32, error)
	GetSkiJumpingDisciplines(ctx context.Context) ([]string, error)
	GetSkiJumpingCategories(ctx context.Context) ([]string, error)
	GetRacesJP(ctx context.Context, seasons []int32, disciplines, cats []string) ([]fissqlc.ARacejp, error)
	GetLastRowRaceJP(ctx context.Context) (fissqlc.ARacejp, error)
	InsertRaceJP(ctx context.Context, in InsertRaceJPClean) error
	UpdateRaceJPByID(ctx context.Context, in UpdateRaceJPClean) error
	DeleteRaceJPByID(ctx context.Context, raceID int32) error
	SearchRacesJP(ctx context.Context, seasoncode *int32, nationcode, gender, catcode *string) ([]fissqlc.SearchRacesJPRow, error)
	GetRacesByIDsJP(ctx context.Context, raceIDs []int32) ([]fissqlc.ARacejp, error)
	GetRaceCountsByCategoryJP(ctx context.Context, seasoncode int32, nationcode, gender *string) ([]fissqlc.GetRaceCountsByCategoryJPRow, error)
	GetRaceCountsByNationJP(ctx context.Context, seasoncode int32, catcode, gender *string) ([]fissqlc.GetRaceCountsByNationJPRow, error)
	GetRaceTotalJP(ctx context.Context, seasoncode int32, catcode, gender *string) (int64, error)
}

// Racenk interface
type Racenk interface {
	GetNordicCombinedSeasons(ctx context.Context) ([]int32, error)
	GetNordicCombinedDisciplines(ctx context.Context) ([]string, error)
	GetNordicCombinedCategories(ctx context.Context) ([]string, error)
	GetRacesNK(ctx context.Context, seasons []int32, disciplines, cats []string) ([]fissqlc.ARacenk, error)
	GetLastRowRaceNK(ctx context.Context) (fissqlc.ARacenk, error)
	InsertRaceNK(ctx context.Context, in InsertRaceNKClean) error
	UpdateRaceNKByID(ctx context.Context, in UpdateRaceNKClean) error
	DeleteRaceNKByID(ctx context.Context, raceID int32) error
	SearchRacesNK(ctx context.Context, seasoncode *int32, nationcode, gender, catcode *string) ([]fissqlc.SearchRacesNKRow, error)
	GetRacesByIDsNK(ctx context.Context, raceIDs []int32) ([]fissqlc.ARacenk, error)
	GetRaceCountsByCategoryNK(ctx context.Context, seasoncode int32, nationcode, gender *string) ([]fissqlc.GetRaceCountsByCategoryNKRow, error)
	GetRaceCountsByNationNK(ctx context.Context, seasoncode int32, catcode, gender *string) ([]fissqlc.GetRaceCountsByNationNKRow, error)
	GetRaceTotalNK(ctx context.Context, seasoncode int32, catcode, gender *string) (int64, error)
}

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

// Athlete interface
type Athlete interface {
	GetAthletesBySporttiID(ctx context.Context, sporttiid int32) ([]AthleteRow, error)
	InsertAthlete(ctx context.Context, in InsertAthleteClean) error
	UpdateAthleteByFiscode(ctx context.Context, in UpdateAthleteClean) error
	DeleteAthleteByFiscode(ctx context.Context, fiscode int32) error
}

// FISStorage struct to hold table-specific storage
type FISStorage struct {
	db          *sql.DB
	competitors Competitors
	racecc      Racecc
	racejp      Racejp
	racenk      Racenk
	resultcc    Resultcc
	resultjp    Resultjp
	resultnk    Resultnk
	athlete     Athlete
}

// Ping method
func (s *FISStorage) Ping(ctx context.Context) error {
	return s.db.PingContext(ctx)
}

// Methods to return each table's storage interface
func (s *FISStorage) Competitors() Competitors {
	return s.competitors
}

func (s *FISStorage) RaceCC() Racecc {
	return s.racecc
}

func (s *FISStorage) RaceJP() Racejp {
	return s.racejp
}

func (s *FISStorage) RaceNK() Racenk {
	return s.racenk
}

func (s *FISStorage) ResultCC() Resultcc {
	return s.resultcc
}

func (s *FISStorage) ResultJP() Resultjp {
	return s.resultjp
}

func (s *FISStorage) ResultNK() Resultnk {
	return s.resultnk
}

func (s *FISStorage) Athlete() Athlete {
	return s.athlete
}

// Storage for FIS database tables
func NewFISStorage(db *sql.DB) *FISStorage {
	return &FISStorage{
		db:          db,
		competitors: &CompetitorsStore{db: db},
		racecc:      &RaceCCStore{db: db},
		racejp:      &RaceJPStore{db: db},
		racenk:      &RaceNKStore{db: db},
		resultcc:    &ResultCCStore{db: db},
		resultjp:    &ResultJPStore{db: db},
		resultnk:    &ResultNKStore{db: db},
		athlete:     &AthleteStore{db: db},
	}
}
