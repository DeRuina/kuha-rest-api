package fis

import (
	"context"
	"database/sql"

	fissqlc "github.com/DeRuina/KUHA-REST-API/internal/db/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type ResultNKStore struct {
	db *sql.DB
}

func (s *ResultNKStore) GetLastRowResultNK(ctx context.Context) (fissqlc.AResultnk, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.GetLastRowResultNK(ctx)
}

func (s *ResultNKStore) InsertResultNK(ctx context.Context, in InsertResultNKClean) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.InsertResultNK(ctx, mapInsertResultNKToParams(in))
}

func (s *ResultNKStore) UpdateResultNKByRecID(ctx context.Context, in UpdateResultNKClean) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	_, err := q.UpdateResultNKByRecID(ctx, mapUpdateResultNKToParams(in))
	return err
}

func (s *ResultNKStore) DeleteResultNKByRecID(ctx context.Context, recid int32) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	_, err := q.DeleteResultNKByRecID(ctx, recid)
	return err
}

func (s *ResultNKStore) GetRaceResultsNKByRaceID(ctx context.Context, raceID int32) ([]fissqlc.AResultnk, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.GetRaceResultsNKByRaceID(ctx, sql.NullInt32{Int32: raceID, Valid: true})
}

func (s *ResultNKStore) GetAthleteResultsNK(
	ctx context.Context,
	competitorID int32,
	seasons []int32,
	disciplines, cats []string,
) ([]fissqlc.GetAthleteResultsNKRow, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	params := fissqlc.GetAthleteResultsNKParams{
		Competitorid: sql.NullInt32{Int32: competitorID, Valid: true},
		Column2:      seasons,
		Column3:      disciplines,
		Column4:      cats,
	}
	return q.GetAthleteResultsNK(ctx, params)
}

func (s *ResultNKStore) GetSeasonsCatcodesNKByCompetitor(
	ctx context.Context,
	fiscode int32,
) ([]fissqlc.GetSeasonsCatcodesNKByCompetitorRow, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.GetSeasonsCatcodesNKByCompetitor(ctx, fiscode)
}

func (s *ResultNKStore) GetLatestResultsNK(
	ctx context.Context,
	fiscode int32,
	seasoncode *int32,
	catcodes []string,
	limit *int32,
) ([]fissqlc.GetLatestResultsNKRow, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)

	params := fissqlc.GetLatestResultsNKParams{
		Column1: fiscode,
		Column2: 0,
		Column3: nil,
		Column4: 50,
	}

	if seasoncode != nil {
		params.Column2 = *seasoncode
	}

	if len(catcodes) > 0 {
		params.Column3 = catcodes
	}

	if limit != nil && *limit > 0 {
		params.Column4 = *limit
	}

	return q.GetLatestResultsNK(ctx, params)
}
