package fis

import (
	"context"
	"database/sql"
	"time"

	fissqlc "github.com/DeRuina/KUHA-REST-API/internal/db/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type CompetitorsStore struct {
	db *sql.DB
}

func (s *CompetitorsStore) GetAthletesBySector(ctx context.Context, sector string) ([]AthleteRow, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	rows, err := q.GetAthletesBySector(ctx, utils.NullString(sector))
	if err != nil {
		return nil, err
	}
	out := make([]AthleteRow, 0, len(rows))
	for _, r := range rows {
		out = append(out, AthleteRow{
			Firstname: utils.StringPtrOrNil(r.Firstname),
			Lastname:  utils.StringPtrOrNil(r.Lastname),
			Fiscode:   utils.Int32PtrOrNil(r.Fiscode),
		})
	}
	return out, nil
}

func (s *CompetitorsStore) GetNationsBySector(ctx context.Context, sector string) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	rows, err := q.GetNationsBySector(ctx, utils.NullString(sector))
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

func (s *CompetitorsStore) GetLastRowCompetitor(ctx context.Context) (fissqlc.ACompetitor, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.GetLastRowCompetitor(ctx)
}

func (s *CompetitorsStore) InsertCompetitor(ctx context.Context, in InsertCompetitorClean) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.InsertCompetitor(ctx, mapInsertToParams(in))
}

func (s *CompetitorsStore) UpdateCompetitorByID(ctx context.Context, in UpdateCompetitorClean) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	_, err := q.UpdateCompetitorByID(ctx, mapUpdateToParams(in))
	return err
}

func (s *CompetitorsStore) DeleteCompetitorByID(ctx context.Context, competitorID int32) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	_, err := q.DeleteCompetitorByID(ctx, competitorID)
	return err
}

func (s *CompetitorsStore) GetCompetitorIDByFiscodeCC(ctx context.Context, fiscode int32) (int32, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.GetCompetitorIDByFiscodeCC(ctx, sql.NullInt32{Int32: fiscode, Valid: true})
}

func (s *CompetitorsStore) GetCompetitorIDByFiscodeJP(ctx context.Context, fiscode int32) (int32, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.GetCompetitorIDByFiscodeJP(ctx, sql.NullInt32{Int32: fiscode, Valid: true})
}

func (s *CompetitorsStore) GetCompetitorIDByFiscodeNK(ctx context.Context, fiscode int32) (int32, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)
	return q.GetCompetitorIDByFiscodeNK(ctx, sql.NullInt32{Int32: fiscode, Valid: true})
}

func (s *CompetitorsStore) SearchCompetitors(
	ctx context.Context,
	nationcode, sectorcode, gender *string,
	birthdateMin, birthdateMax *time.Time,
) ([]fissqlc.ACompetitor, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := fissqlc.New(s.db)

	params := fissqlc.SearchCompetitorsParams{
		Column1: "",
		Column2: "",
		Column3: "",
		Column4: time.Time{},
		Column5: time.Time{},
	}

	if nationcode != nil {
		params.Column1 = *nationcode
	}
	if sectorcode != nil {
		params.Column2 = *sectorcode
	}
	if gender != nil {
		params.Column3 = *gender
	}
	if birthdateMin != nil {
		params.Column4 = *birthdateMin
	}
	if birthdateMax != nil {
		params.Column5 = *birthdateMax
	}

	return q.SearchCompetitors(ctx, params)
}
