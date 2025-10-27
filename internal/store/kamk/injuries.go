package kamk

import (
	"context"
	"database/sql"
	"time"

	kamksqlc "github.com/DeRuina/KUHA-REST-API/internal/db/kamk"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type InjuriesStore struct {
	db *sql.DB
}

type Injury struct {
	CompetitorID int32      `json:"competitor_id"`
	InjuryType   int32      `json:"injury_type"`
	Severity     *int32     `json:"severity,omitempty"`
	PainLevel    *int32     `json:"pain_level,omitempty"`
	Description  *string    `json:"description,omitempty"`
	DateStart    time.Time  `json:"date_start"`
	Status       int32      `json:"status"`
	DateEnd      *time.Time `json:"date_end,omitempty"`
	InjuryID     *int32     `json:"injury_id,omitempty"`
	Meta         *string    `json:"meta,omitempty"`
}

type InjuryInput struct {
	InjuryType  int32
	Severity    *int32
	PainLevel   *int32
	Description *string
	InjuryID    int32
	Meta        *string
}

func (s *InjuriesStore) AddInjury(ctx context.Context, userID int32, in InjuryInput) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := kamksqlc.New(s.db)

	arg := kamksqlc.InsertInjuryParams{
		CompetitorID: userID,
		InjuryType:   in.InjuryType,
		Severity:     utils.NullInt32Ptr(in.Severity),
		PainLevel:    utils.NullInt32Ptr(in.PainLevel),
		Description:  utils.NullStringPtr(in.Description),
		InjuryID:     utils.NullInt32(in.InjuryID),
		Meta:         utils.NullStringPtr(in.Meta),
	}

	return q.InsertInjury(ctx, arg)
}

func (s *InjuriesStore) MarkInjuryRecovered(ctx context.Context, userID int32, injuryID int32) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := kamksqlc.New(s.db)
	if err := q.MarkInjuryRecoveredByID(ctx, kamksqlc.MarkInjuryRecoveredByIDParams{
		InjuryID:     utils.NullInt32(injuryID),
		CompetitorID: userID,
	}); err != nil {
		return 0, err
	}

	return 1, nil
}

func (s *InjuriesStore) GetActiveInjuries(ctx context.Context, userID int32) ([]Injury, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := kamksqlc.New(s.db)
	rows, err := q.GetActiveInjuriesByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	out := make([]Injury, 0, len(rows))
	for _, r := range rows {
		out = append(out, Injury{
			CompetitorID: r.CompetitorID,
			InjuryType:   r.InjuryType,
			Severity:     utils.Int32PtrOrNil(r.Severity),
			PainLevel:    utils.Int32PtrOrNil(r.PainLevel),
			Description:  utils.StringPtrOrNil(r.Description),
			DateStart:    r.DateStart,
			Status:       r.Status,
			DateEnd:      utils.TimePtrOrNil(r.DateEnd),
			InjuryID:     utils.Int32PtrOrNil(r.InjuryID),
			Meta:         utils.StringPtrOrNil(r.Meta),
		})
	}
	return out, nil
}

func (s *InjuriesStore) GetMaxInjuryID(ctx context.Context, userID int32) (int32, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := kamksqlc.New(s.db)
	return q.GetMaxInjuryIDForUser(ctx, userID)
}

func (s *InjuriesStore) DeleteInjury(ctx context.Context, userID int32, injuryID int32) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := kamksqlc.New(s.db)
	n, err := q.DeleteInjuryByID(ctx, kamksqlc.DeleteInjuryByIDParams{
		CompetitorID: userID,
		InjuryID:     utils.NullInt32(injuryID),
	})
	if err != nil {
		return 0, err
	}
	return n, nil
}
