package fis

import (
	"context"
	"database/sql"

	fissqlc "github.com/DeRuina/KUHA-REST-API/internal/db/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type ResultJPStore struct {
	db *sql.DB
}

func (s *ResultJPStore) GetLastRowResultJP(ctx context.Context) (fissqlc.AResultjp, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.GetLastRowResultJP(ctx)
}

func (s *ResultJPStore) InsertResultJP(ctx context.Context, in InsertResultJPClean) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.InsertResultJP(ctx, mapInsertResultJPToParams(in))
}

func (s *ResultJPStore) UpdateResultJPByRecID(ctx context.Context, in UpdateResultJPClean) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	_, err := q.UpdateResultJPByRecID(ctx, mapUpdateResultJPToParams(in))
	return err
}

func (s *ResultJPStore) DeleteResultJPByRecID(ctx context.Context, recid int32) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	_, err := q.DeleteResultJPByRecID(ctx, recid)
	return err
}

func (s *ResultJPStore) GetRaceResultsJPByRaceID(ctx context.Context, raceID int32) ([]fissqlc.AResultjp, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.GetRaceResultsJPByRaceID(ctx, sql.NullInt32{Int32: raceID, Valid: true})
}

func (s *ResultJPStore) GetAthleteResultsJP(
	ctx context.Context,
	competitorID int32,
	seasons []int32,
	disciplines, cats []string,
) ([]fissqlc.GetAthleteResultsJPRow, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	params := fissqlc.GetAthleteResultsJPParams{
		Competitorid: sql.NullInt32{Int32: competitorID, Valid: true},
		Column2:      seasons,
		Column3:      disciplines,
		Column4:      cats,
	}
	return q.GetAthleteResultsJP(ctx, params)
}

func (s *ResultJPStore) GetSeasonsCatcodesJPByCompetitor(
	ctx context.Context,
	fiscode int32,
) ([]fissqlc.GetSeasonsCatcodesJPByCompetitorRow, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.GetSeasonsCatcodesJPByCompetitor(ctx, fiscode)
}

func (s *ResultJPStore) GetLatestResultsJP(
	ctx context.Context,
	fiscode int32,
	seasoncode *int32,
	catcodes []string,
	limit *int32,
) ([]fissqlc.GetLatestResultsJPRow, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)

	params := fissqlc.GetLatestResultsJPParams{
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

	return q.GetLatestResultsJP(ctx, params)
}
