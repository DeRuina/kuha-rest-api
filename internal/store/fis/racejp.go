package fis

import (
	"context"
	"database/sql"

	fissqlc "github.com/DeRuina/KUHA-REST-API/internal/db/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type RaceJPStore struct {
	db *sql.DB
}

func (s *RaceJPStore) GetSkiJumpingSeasons(ctx context.Context) ([]int32, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	rows, err := q.GetSkiJumpingSeasons(ctx)
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

func (s *RaceJPStore) GetSkiJumpingDisciplines(ctx context.Context) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	rows, err := q.GetSkiJumpingDisciplines(ctx)
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

func (s *RaceJPStore) GetSkiJumpingCategories(ctx context.Context) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	rows, err := q.GetSkiJumpingCategories(ctx)
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

func (s *RaceJPStore) GetRacesJP(ctx context.Context, seasons []int32, disciplines, cats []string) ([]fissqlc.ARacejp, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	params := fissqlc.GetRacesJPParams{
		Column1: seasons,
		Column2: disciplines,
		Column3: cats,
	}
	return q.GetRacesJP(ctx, params)
}

func (s *RaceJPStore) GetLastRowRaceJP(ctx context.Context) (fissqlc.ARacejp, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.GetLastRowRaceJP(ctx)
}

func (s *RaceJPStore) InsertRaceJP(ctx context.Context, in InsertRaceJPClean) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.InsertRaceJP(ctx, mapInsertRaceJPToParams(in))
}

func (s *RaceJPStore) UpdateRaceJPByID(ctx context.Context, in UpdateRaceJPClean) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	_, err := q.UpdateRaceJPByID(ctx, mapUpdateRaceJPToParams(in))
	return err
}

func (s *RaceJPStore) DeleteRaceJPByID(ctx context.Context, raceID int32) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	_, err := q.DeleteRaceJPByID(ctx, raceID)
	return err
}

func (s *RaceJPStore) SearchRacesJP(
	ctx context.Context,
	seasoncode *int32,
	nationcode, gender, catcode *string,
) ([]fissqlc.SearchRacesJPRow, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)

	params := fissqlc.SearchRacesJPParams{
		Column1: 0,
		Column2: "",
		Column3: "",
		Column4: "",
	}

	if seasoncode != nil {
		params.Column1 = *seasoncode
	}
	if nationcode != nil {
		params.Column2 = *nationcode
	}
	if gender != nil {
		params.Column3 = *gender
	}
	if catcode != nil {
		params.Column4 = *catcode
	}

	return q.SearchRacesJP(ctx, params)
}

func (s *RaceJPStore) GetRacesByIDsJP(
	ctx context.Context,
	raceIDs []int32,
) ([]fissqlc.ARacejp, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	if len(raceIDs) == 0 {
		return []fissqlc.ARacejp{}, nil
	}

	q := fissqlc.New(s.db)
	return q.GetRacesByIDsJP(ctx, raceIDs)
}
