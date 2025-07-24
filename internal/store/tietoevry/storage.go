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
}

type Exercises interface {
	ValidateUsersExist(ctx context.Context, userIDs []uuid.UUID) error
	InsertExercisesBulk(ctx context.Context, exercises []ExercisePayload) error
}

type Symptoms interface {
	ValidateUsersExist(ctx context.Context, userIDs []uuid.UUID) error
	InsertSymptomsBulk(ctx context.Context, symptoms []tietoevrysqlc.InsertSymptomParams) error
}

type Measurements interface {
	ValidateUsersExist(ctx context.Context, userIDs []uuid.UUID) error
	InsertMeasurementsBulk(ctx context.Context, measurements []tietoevrysqlc.InsertMeasurementParams) error
}

type TestResults interface {
	ValidateUsersExist(ctx context.Context, userIDs []uuid.UUID) error
	InsertTestResultsBulk(ctx context.Context, results []tietoevrysqlc.InsertTestResultParams) error
}

type Questionnaires interface {
	ValidateUsersExist(ctx context.Context, userIDs []uuid.UUID) error
	InsertQuestionnaireAnswersBulk(ctx context.Context, answers []tietoevrysqlc.InsertQuestionnaireAnswerParams) error
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
	}
}
