package fis

import (
	"context"
	"database/sql"

	fissqlc "github.com/DeRuina/KUHA-REST-API/internal/db/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type RaceNKStore struct {
	db *sql.DB
}

func (s *RaceNKStore) GetNordicCombinedSeasons(ctx context.Context) ([]int32, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	rows, err := q.GetNordicCombinedSeasons(ctx)
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

func (s *RaceNKStore) GetNordicCombinedDisciplines(ctx context.Context) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	rows, err := q.GetNordicCombinedDisciplines(ctx)
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

func (s *RaceNKStore) GetNordicCombinedCategories(ctx context.Context) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	rows, err := q.GetNordicCombinedCategories(ctx)
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

func (s *RaceNKStore) GetRacesNK(ctx context.Context, seasons []int32, disciplines, cats []string) ([]fissqlc.ARacenk, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	params := fissqlc.GetRacesNKParams{
		Column1: seasons,
		Column2: disciplines,
		Column3: cats,
	}
	return q.GetRacesNK(ctx, params)
}

func (s *RaceNKStore) GetLastRowRaceNK(ctx context.Context) (fissqlc.ARacenk, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.GetLastRowRaceNK(ctx)
}

func (s *RaceNKStore) InsertRaceNK(ctx context.Context, in InsertRaceNKClean) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.InsertRaceNK(ctx, mapInsertRaceNKToParams(in))
}

func (s *RaceNKStore) UpdateRaceNKByID(ctx context.Context, in UpdateRaceNKClean) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	_, err := q.UpdateRaceNKByID(ctx, mapUpdateRaceNKToParams(in))
	return err
}

func (s *RaceNKStore) DeleteRaceNKByID(ctx context.Context, raceID int32) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	_, err := q.DeleteRaceNKByID(ctx, raceID)
	return err
}

func (s *RaceNKStore) SearchRacesNK(
	ctx context.Context,
	seasoncode *int32,
	nationcode, gender, catcode *string,
) ([]fissqlc.SearchRacesNKRow, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)

	params := fissqlc.SearchRacesNKParams{
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

	return q.SearchRacesNK(ctx, params)
}

func (s *RaceNKStore) GetRacesByIDsNK(
	ctx context.Context,
	raceIDs []int32,
) ([]fissqlc.ARacenk, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	if len(raceIDs) == 0 {
		return []fissqlc.ARacenk{}, nil
	}

	q := fissqlc.New(s.db)
	return q.GetRacesByIDsNK(ctx, raceIDs)
}

func (s *RaceNKStore) GetRaceCountsByCategoryNK(
	ctx context.Context,
	seasoncode int32,
	nationcode, gender *string,
) ([]fissqlc.GetRaceCountsByCategoryNKRow, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)

	params := fissqlc.GetRaceCountsByCategoryNKParams{
		Column1: seasoncode,
		Column2: "",
		Column3: "",
	}

	if nationcode != nil {
		params.Column2 = *nationcode
	}
	if gender != nil {
		params.Column3 = *gender
	}

	return q.GetRaceCountsByCategoryNK(ctx, params)
}

func (s *RaceNKStore) GetRaceCountsByNationNK(
	ctx context.Context,
	seasoncode int32,
	catcode, gender *string,
) ([]fissqlc.GetRaceCountsByNationNKRow, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)

	params := fissqlc.GetRaceCountsByNationNKParams{
		Column1: seasoncode,
		Column2: "",
		Column3: "",
	}

	if catcode != nil {
		params.Column2 = *catcode
	}
	if gender != nil {
		params.Column3 = *gender
	}

	return q.GetRaceCountsByNationNK(ctx, params)
}

func (s *RaceNKStore) GetRaceTotalNK(
	ctx context.Context,
	seasoncode int32,
	catcode, gender *string,
) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)

	params := fissqlc.GetRaceTotalNKParams{
		Column1: seasoncode,
		Column2: "",
		Column3: "",
	}

	if catcode != nil {
		params.Column2 = *catcode
	}

	if gender != nil {
		params.Column3 = *gender
	}

	return q.GetRaceTotalNK(ctx, params)
}
