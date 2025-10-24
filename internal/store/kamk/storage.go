package kamk

import (
	"context"
	"database/sql"
	"time"
)

type Injuries interface {
	AddInjury(ctx context.Context, userID int32, in InjuryInput) error
	MarkInjuryRecovered(ctx context.Context, userID int32, injuryID int32) (int64, error)
	GetActiveInjuries(ctx context.Context, userID int32) ([]Injury, error)
	GetMaxInjuryID(ctx context.Context, userID int32) (int32, error)
}

type Queries interface {
	AddQuestionnaire(ctx context.Context, userID int32, in QuestionnaireInput) error
	GetQuestionnaires(ctx context.Context, userID int32) ([]Questionnaire, error)
	IsQuizDoneToday(ctx context.Context, userID int32, queryType int32) ([]Questionnaire, error)
	UpdateQuestionnaireByTimestamp(ctx context.Context, userID int32, ts time.Time, answers string, comment *string) error
}

// KAMKStorage
type KAMKStorage struct {
	db       *sql.DB
	injuries Injuries
	queries  Queries
}

// Methods
func (s *KAMKStorage) Ping(ctx context.Context) error {
	return s.db.PingContext(ctx)
}

func (s *KAMKStorage) Injuries() Injuries {
	return s.injuries
}

func (s *KAMKStorage) Queries() Queries {
	return s.queries
}

// NewKAMKStorage creates a new KAMKStorage instance
func NewKAMKStorage(db *sql.DB) *KAMKStorage {
	return &KAMKStorage{
		db:       db,
		injuries: &InjuriesStore{db: db},
		queries:  &QueriesStore{db: db},
	}
}
