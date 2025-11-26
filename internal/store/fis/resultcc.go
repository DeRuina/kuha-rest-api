package fis

import (
	"context"
	"database/sql"

	fissqlc "github.com/DeRuina/KUHA-REST-API/internal/db/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type ResultCCStore struct {
	db *sql.DB
}

func (s *ResultCCStore) GetLastRowResultCC(ctx context.Context) (fissqlc.AResultcc, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.GetLastRowResultCC(ctx)
}

func (s *ResultCCStore) InsertResultCC(ctx context.Context, in InsertResultCCClean) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.InsertResultCC(ctx, mapInsertResultCCToParams(in))
}

func (s *ResultCCStore) UpdateResultCCByRecID(ctx context.Context, in UpdateResultCCClean) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	_, err := q.UpdateResultCCByRecID(ctx, mapUpdateResultCCToParams(in))
	return err
}

func (s *ResultCCStore) DeleteResultCCByRecID(ctx context.Context, recid int32) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	_, err := q.DeleteResultCCByRecID(ctx, recid)
	return err
}

func (s *ResultCCStore) GetRaceResultsCCByRaceID(ctx context.Context, raceID int32) ([]fissqlc.AResultcc, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.GetRaceResultsCCByRaceID(ctx, sql.NullInt32{Int32: raceID, Valid: true})
}

func (s *ResultCCStore) GetAthleteResultsCC(
	ctx context.Context,
	competitorID int32,
	seasons []int32,
	disciplines, cats []string,
) ([]fissqlc.GetAthleteResultsCCRow, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	params := fissqlc.GetAthleteResultsCCParams{
		Competitorid: sql.NullInt32{Int32: competitorID, Valid: true},
		Column2:      seasons,
		Column3:      disciplines,
		Column4:      cats,
	}
	return q.GetAthleteResultsCC(ctx, params)
}

func (s *ResultCCStore) GetSeasonsCatcodesCCByCompetitor(
	ctx context.Context,
	competitorID int32,
) ([]fissqlc.GetSeasonsCatcodesCCByCompetitorRow, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.GetSeasonsCatcodesCCByCompetitor(ctx, competitorID)
}

func (s *ResultCCStore) GetLatestResultsCC(
	ctx context.Context,
	competitorID int32,
	seasoncode *int32,
	catcodes []string,
	limit *int32,
) ([]fissqlc.GetLatestResultsCCRow, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)

	params := fissqlc.GetLatestResultsCCParams{
		Column1: competitorID,
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

	return q.GetLatestResultsCC(ctx, params)
}
