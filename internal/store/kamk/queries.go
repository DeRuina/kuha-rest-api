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

func (s *QueriesStore) AddQuestionnaire(ctx context.Context, userID string, in QuestionnaireInput) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	sportti, err := utils.ParseSporttiID(userID)
	if err != nil {
		return err
	}
	competitorID, err := utils.ParsePositiveInt32(sportti)
	if err != nil {
		return err
	}

	q := kamksqlc.New(s.db)
	arg := kamksqlc.InsertQuestionnaireParams{
		CompetitorID: competitorID,
		QueryType:    utils.NullInt32Ptr(in.QueryType),
		Answers:      utils.NullStringPtr(in.Answers),
		Comment:      utils.NullStringPtr(in.Comment),
		Meta:         utils.NullStringPtr(in.Meta),
	}
	return q.InsertQuestionnaire(ctx, arg)
}

func (s *QueriesStore) GetQuestionnaires(ctx context.Context, userID string) ([]Questionnaire, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	sportti, err := utils.ParseSporttiID(userID)
	if err != nil {
		return nil, err
	}
	competitorID, err := utils.ParsePositiveInt32(sportti)
	if err != nil {
		return nil, err
	}

	q := kamksqlc.New(s.db)
	rows, err := q.GetQuestionnairesByUser(ctx, competitorID)
	if err != nil {
		return nil, err
	}

	out := make([]Questionnaire, 0, len(rows))
	for _, r := range rows {
		out = append(out, Questionnaire{
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

func (s *QueriesStore) IsQuizDoneToday(ctx context.Context, userID string, queryType int32) ([]Questionnaire, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	sportti, err := utils.ParseSporttiID(userID)
	if err != nil {
		return nil, err
	}
	competitorID, err := utils.ParsePositiveInt32(sportti)
	if err != nil {
		return nil, err
	}

	now := time.Now().UTC()
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	end := start.Add(24 * time.Hour)

	q := kamksqlc.New(s.db)
	rows, err := q.IsQuizDoneToday(ctx, kamksqlc.IsQuizDoneTodayParams{
		CompetitorID: competitorID,
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

func (s *QueriesStore) UpdateQuestionnaireByTimestamp(ctx context.Context, userID string, ts time.Time, answers string, comment *string) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	sportti, err := utils.ParseSporttiID(userID)
	if err != nil {
		return err
	}
	competitorID, err := utils.ParsePositiveInt32(sportti)
	if err != nil {
		return err
	}

	q := kamksqlc.New(s.db)
	return q.UpdateQuestionnaireByTimestamp(ctx, kamksqlc.UpdateQuestionnaireByTimestampParams{
		CompetitorID: competitorID,
		Timestamp:    ts,
		Answers:      utils.NullString(answers),
		Comment:      utils.NullStringPtr(comment),
	})
}
