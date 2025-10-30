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
	ID           int64     `json:"id"`
	CompetitorID int32     `json:"competitor_id"`
	QueryType    *int32    `json:"query_type,omitempty"`
	Answers      *string   `json:"answers,omitempty"`
	Comment      *string   `json:"comment,omitempty"`
	Timestamp    time.Time `json:"timestamp"`
	Meta         *string   `json:"meta,omitempty"`
}

type QuestionnaireInput struct {
	QueryType *int32
	Answers   *string
	Comment   *string
	Meta      *string
}

func (s *QueriesStore) AddQuestionnaire(ctx context.Context, userID int32, in QuestionnaireInput) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := kamksqlc.New(s.db)
	arg := kamksqlc.InsertQuestionnaireParams{
		CompetitorID: userID,
		QueryType:    utils.NullInt32Ptr(in.QueryType),
		Answers:      utils.NullStringPtr(in.Answers),
		Comment:      utils.NullStringPtr(in.Comment),
		Meta:         utils.NullStringPtr(in.Meta),
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
			ID:           r.ID,
			CompetitorID: r.CompetitorID,
			QueryType:    utils.Int32PtrOrNil(r.QueryType),
			Answers:      utils.StringPtrOrNil(r.Answers),
			Comment:      utils.StringPtrOrNil(r.Comment),
			Timestamp:    r.Timestamp,
			Meta:         utils.StringPtrOrNil(r.Meta),
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
		CompetitorID: userID,
		QueryType:    utils.NullInt32(queryType),
		Timestamp:    start,
		Timestamp_2:  end,
	})
	if err != nil {
		return nil, err
	}

	out := make([]Questionnaire, 0, len(rows))
	for _, r := range rows {
		out = append(out, Questionnaire{
			ID:           r.ID,
			CompetitorID: r.CompetitorID,
			QueryType:    utils.Int32PtrOrNil(r.QueryType),
			Answers:      utils.StringPtrOrNil(r.Answers),
			Comment:      utils.StringPtrOrNil(r.Comment),
			Timestamp:    r.Timestamp,
			Meta:         utils.StringPtrOrNil(r.Meta),
		})
	}
	return out, nil
}

func (s *QueriesStore) UpdateQuestionnaireByID(ctx context.Context, userID int32, id int64, answers string, comment *string) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := kamksqlc.New(s.db)
	n, err := q.UpdateQuestionnaireByID(ctx, kamksqlc.UpdateQuestionnaireByIDParams{
		CompetitorID: userID,
		ID:           id,
		Answers:      utils.NullString(answers),
		Comment:      utils.NullStringPtr(comment),
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
		CompetitorID: userID,
		ID:           id,
	})
	if err != nil {
		return 0, err
	}
	return n, nil
}
