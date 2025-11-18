package fis

import (
	"context"
	"database/sql"

	fissqlc "github.com/DeRuina/KUHA-REST-API/internal/db/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type RaceCCStore struct {
	db *sql.DB
}

func (s *RaceCCStore) GetCrossCountrySeasons(ctx context.Context) ([]int32, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	rows, err := q.GetCrossCountrySeasons(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]int32, 0, len(rows))
	for _, v := range rows {
		if v.Valid {
			out = append(out, v.Int32)
		}
	}
	return out, nil
}

func (s *RaceCCStore) GetCrossCountryDisciplines(ctx context.Context) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	rows, err := q.GetCrossCountryDisciplines(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]string, 0, len(rows))
	for _, v := range rows {
		if v.Valid {
			out = append(out, v.String)
		}
	}
	return out, nil
}

func (s *RaceCCStore) GetCrossCountryCategories(ctx context.Context) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	rows, err := q.GetCrossCountryCategories(ctx)
	if err != nil {
		return nil, err
	}

	out := make([]string, 0, len(rows))
	for _, v := range rows {
		if v.Valid {
			out = append(out, v.String)
		}
	}
	return out, nil
}

func (s *RaceCCStore) GetRacesCC(ctx context.Context, seasons []int32, disciplines, cats []string) ([]fissqlc.ARacecc, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	params := fissqlc.GetRacesCCParams{
		Column1: seasons,
		Column2: disciplines,
		Column3: cats,
	}
	return q.GetRacesCC(ctx, params)
}

func (s *RaceCCStore) GetLastRowRaceCC(ctx context.Context) (fissqlc.ARacecc, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.GetLastRowRaceCC(ctx)
}

func (s *RaceCCStore) InsertRaceCC(ctx context.Context, in InsertRaceCCClean) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.InsertRaceCC(ctx, mapInsertRaceCCToParams(in))
}

func (s *RaceCCStore) UpdateRaceCCByID(ctx context.Context, in UpdateRaceCCClean) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	_, err := q.UpdateRaceCCByID(ctx, mapUpdateRaceCCToParams(in))
	return err
}

func (s *RaceCCStore) DeleteRaceCCByID(ctx context.Context, raceID int32) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	_, err := q.DeleteRaceCCByID(ctx, raceID)
	return err
}
