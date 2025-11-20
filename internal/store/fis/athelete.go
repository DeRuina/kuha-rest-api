package fis

import (
	"context"
	"database/sql"

	fissqlc "github.com/DeRuina/KUHA-REST-API/internal/db/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type AthleteStore struct {
	db *sql.DB
}

func (s *AthleteStore) GetAthletesBySporttiID(ctx context.Context, sporttiid int32) ([]AthleteRow, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	rows, err := q.GetAthletesBySporttiID(ctx, sql.NullInt32{Int32: sporttiid, Valid: true})
	if err != nil {
		return nil, err
	}

	out := make([]AthleteRow, 0, len(rows))
	for _, r := range rows {
		out = append(out, AthleteRow{
			Firstname: utils.StringPtrOrNil(r.Firstname),
			Lastname:  utils.StringPtrOrNil(r.Lastname),
			Fiscode:   &r.Fiscode,
		})
	}
	return out, nil
}

func (s *AthleteStore) InsertAthlete(ctx context.Context, in InsertAthleteClean) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.InsertAthlete(ctx, mapInsertAthleteToParams(in))
}

func (s *AthleteStore) UpdateAthleteByFiscode(ctx context.Context, in UpdateAthleteClean) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	_, err := q.UpdateAthleteByFiscode(ctx, mapUpdateAthleteToParams(in))
	return err
}

func (s *AthleteStore) DeleteAthleteByFiscode(ctx context.Context, fiscode int32) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	_, err := q.DeleteAthleteByFiscode(ctx, fiscode)
	return err
}
