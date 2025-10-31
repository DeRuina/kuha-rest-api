package kamk

import (
	"context"
	"database/sql"
	"time"

	kamksqlc "github.com/DeRuina/KUHA-REST-API/internal/db/kamk"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type QueriesStore struct {
	db *sql.DB
}

type Questionnaire struct {
	ID        int64     `json:"id"`
	UserID    int32     `json:"user_id"`
	QueryType int32     `json:"query_type"`
	Answers   string    `json:"answers"`
	Comment   string    `json:"comment"`
	Timestamp time.Time `json:"timestamp"`
	Meta      string    `json:"meta"`
}

type QuestionnaireInput struct {
	QueryType int32
	Answers   string
	Comment   string
	Meta      string
}

func (s *QueriesStore) AddQuestionnaire(ctx context.Context, userID int32, in QuestionnaireInput) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := kamksqlc.New(s.db)
	arg := kamksqlc.InsertQuestionnaireParams{
		UserID:    userID,
		QueryType: in.QueryType,
		Answers:   in.Answers,
		Comment:   in.Comment,
		Meta:      in.Meta,
	}
	return q.InsertQuestionnaire(ctx, arg)
}

func (s *QueriesStore) GetQuestionnaires(ctx context.Context, userID int32) ([]Questionnaire, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := kamksqlc.New(s.db)
	rows, err := q.GetQuestionnairesByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	out := make([]Questionnaire, 0, len(rows))
	for _, r := range rows {
		out = append(out, Questionnaire{
			ID:        r.ID,
			UserID:    r.UserID,
			QueryType: r.QueryType,
			Answers:   r.Answers,
			Comment:   r.Comment,
			Timestamp: r.Timestamp,
			Meta:      r.Meta,
		})
	}
	return out, nil
}

func (s *QueriesStore) IsQuizDoneToday(ctx context.Context, userID int32, queryType int32) ([]Questionnaire, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	now := time.Now().UTC()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	end := start.Add(24 * time.Hour)

	q := kamksqlc.New(s.db)
	rows, err := q.IsQuizDoneToday(ctx, kamksqlc.IsQuizDoneTodayParams{
		UserID:      userID,
		QueryType:   queryType,
		Timestamp:   start,
		Timestamp_2: end,
	})
	if err != nil {
		return nil, err
	}

	out := make([]Questionnaire, 0, len(rows))
	for _, r := range rows {
		out = append(out, Questionnaire{
			ID:        r.ID,
			UserID:    r.UserID,
			QueryType: r.QueryType,
			Answers:   r.Answers,
			Comment:   r.Comment,
			Timestamp: r.Timestamp,
			Meta:      r.Meta,
		})
	}
	return out, nil
}

func (s *QueriesStore) UpdateQuestionnaireByID(ctx context.Context, userID int32, id int64, answers string, comment string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := kamksqlc.New(s.db)
	n, err := q.UpdateQuestionnaireByID(ctx, kamksqlc.UpdateQuestionnaireByIDParams{
		UserID:  userID,
		ID:      id,
		Answers: answers,
		Comment: comment,
	})
	if err != nil {
		return 0, err
	}
	return n, nil
}

func (s *QueriesStore) DeleteQuestionnaireByID(ctx context.Context, userID int32, id int64) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := kamksqlc.New(s.db)
	n, err := q.DeleteQuestionnaireByID(ctx, kamksqlc.DeleteQuestionnaireByIDParams{
		UserID: userID,
		ID:     id,
	})
	if err != nil {
		return 0, err
	}
	return n, nil
}
