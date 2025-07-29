package tietoevry

import (
	"context"
	"database/sql"

	tietoevrysqlc "github.com/DeRuina/KUHA-REST-API/internal/db/tietoevry"
	"github.com/google/uuid"
)

// Interfaces
type Users interface {
	UpsertUser(ctx context.Context, arg tietoevrysqlc.UpsertUserParams) error
	DeleteUser(ctx context.Context, id uuid.UUID) (int64, error)
	GetUser(ctx context.Context, id uuid.UUID) (tietoevrysqlc.User, error)
	LogDeletedUser(ctx context.Context, userID uuid.UUID) error
	DeleteUserWithLogging(ctx context.Context, userID uuid.UUID) (int64, error)
	GetDeletedUsers(ctx context.Context) ([]tietoevrysqlc.DeletedUsersLog, error)
}

type Exercises interface {
	ValidateUsersExist(ctx context.Context, userIDs []uuid.UUID) error
	InsertExercisesBulk(ctx context.Context, exercises []ExercisePayload) error
	GetExercisesByUser(ctx context.Context, userID uuid.UUID) ([]tietoevrysqlc.Exercise, error)
	GetExerciseHRZones(ctx context.Context, id uuid.UUID) ([]tietoevrysqlc.ExerciseHrZone, error)
	GetExerciseSamples(ctx context.Context, id uuid.UUID) ([]tietoevrysqlc.ExerciseSample, error)
	GetExerciseSections(ctx context.Context, id uuid.UUID) ([]tietoevrysqlc.ExerciseSection, error)
}

type Symptoms interface {
	ValidateUsersExist(ctx context.Context, userIDs []uuid.UUID) error
	InsertSymptomsBulk(ctx context.Context, symptoms []tietoevrysqlc.InsertSymptomParams) error
	GetSymptomsByUser(ctx context.Context, userID uuid.UUID) ([]tietoevrysqlc.Symptom, error)
}

type Measurements interface {
	ValidateUsersExist(ctx context.Context, userIDs []uuid.UUID) error
	InsertMeasurementsBulk(ctx context.Context, measurements []tietoevrysqlc.InsertMeasurementParams) error
	GetMeasurementsByUser(ctx context.Context, userID uuid.UUID) ([]tietoevrysqlc.Measurement, error)
}

type TestResults interface {
	ValidateUsersExist(ctx context.Context, userIDs []uuid.UUID) error
	InsertTestResultsBulk(ctx context.Context, results []tietoevrysqlc.InsertTestResultParams) error
	GetTestResultsByUser(ctx context.Context, userID uuid.UUID) ([]tietoevrysqlc.TestResult, error)
}

type Questionnaires interface {
	ValidateUsersExist(ctx context.Context, userIDs []uuid.UUID) error
	InsertQuestionnaireAnswersBulk(ctx context.Context, answers []tietoevrysqlc.InsertQuestionnaireAnswerParams) error
	GetQuestionnairesByUser(ctx context.Context, userID uuid.UUID) ([]tietoevrysqlc.QuestionAnswer, error)
}

type ActivityZones interface {
	ValidateUsersExist(ctx context.Context, userIDs []uuid.UUID) error
	InsertActivityZonesBulk(ctx context.Context, zones []tietoevrysqlc.InsertActivityZoneParams) error
}

// TietoevryStorage
type TietoevryStorage struct {
	db             *sql.DB
	users          Users
	exercises      Exercises
	symptoms       Symptoms
	measurements   Measurements
	testResults    TestResults
	questionnaires Questionnaires
	activityZones  ActivityZones
}

// Methods
func (s *TietoevryStorage) Ping(ctx context.Context) error {
	return s.db.PingContext(ctx)
}

func (s *TietoevryStorage) Users() Users {
	return s.users
}

func (s *TietoevryStorage) Exercises() Exercises {
	return s.exercises
}

func (s *TietoevryStorage) Symptoms() Symptoms {
	return s.symptoms
}

func (s *TietoevryStorage) Measurements() Measurements {
	return s.measurements
}

func (s *TietoevryStorage) TestResults() TestResults {
	return s.testResults
}

func (s *TietoevryStorage) Questionnaires() Questionnaires {
	return s.questionnaires
}

func (s *TietoevryStorage) ActivityZones() ActivityZones {
	return s.activityZones
}

// NewTietoevryStorage creates a new TietoevryStorage instance
func NewTietoevryStorage(db *sql.DB) *TietoevryStorage {
	return &TietoevryStorage{
		db:             db,
		users:          &UserStore{db: db},
		exercises:      &ExercisesStore{db: db},
		symptoms:       &SymptomsStore{db: db},
		measurements:   &MeasurementsStore{db: db},
		testResults:    &TestResultsStore{db: db},
		questionnaires: &QuestionnairesStore{db: db},
		activityZones:  &ActivityZonesStore{db: db},
	}
}
