package fis

import (
	"context"
	"database/sql"

	fissqlc "github.com/DeRuina/KUHA-REST-API/internal/db/fis"
)

// CompetitorsStore struct
type CompetitorsStore struct {
	db *sql.DB
}

// GetBySector

type GetBySectorResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	FisCode   int32  `json:"fis_code"`
}

func (s *CompetitorsStore) GetBySector(ctx context.Context, sectorCode string) ([]GetBySectorResponse, error) {
	query := `SELECT Firstname, Lastname, Fiscode FROM A_competitor WHERE SectorCode = $1 ORDER BY Fiscode`

	rows, err := s.db.QueryContext(ctx, query, sectorCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var competitors []GetBySectorResponse
	for rows.Next() {
		var c fissqlc.ACompetitor
		err := rows.Scan(&c.Firstname, &c.Lastname, &c.Fiscode)
		if err != nil {
			return nil, err
		}

		competitors = append(competitors, GetBySectorResponse{
			FirstName: c.Firstname.String,
			LastName:  c.Lastname.String,
			FisCode:   c.Fiscode.Int32,
		})
	}

	return competitors, nil
}

// Implement GetByFiscodeJP
func (s *CompetitorsStore) GetByFiscodeJP(ctx context.Context, fiscode int32) (int32, error) {
	query := `SELECT CompetitorID FROM A_competitor WHERE Fiscode = $1 AND SectorCode = 'JP'`

	var competitorID int32
	err := s.db.QueryRowContext(ctx, query, fiscode).Scan(&competitorID)
	if err != nil {
		return 0, err
	}

	return competitorID, nil
}

// Implement GetByFiscodeNK
func (s *CompetitorsStore) GetByFiscodeNK(ctx context.Context, fiscode int32) (int32, error) {
	query := `SELECT CompetitorID FROM A_competitor WHERE Fiscode = $1 AND SectorCode = 'NK'`

	var competitorID int32
	err := s.db.QueryRowContext(ctx, query, fiscode).Scan(&competitorID)
	if err != nil {
		return 0, err
	}

	return competitorID, nil
}

// Implement GetByGenderAndNationJP
func (s *CompetitorsStore) GetByGenderAndNationJP(ctx context.Context, gender, nation string) ([]int32, error) {
	query := `SELECT CompetitorID FROM A_competitor WHERE Gender = $1 AND NationCode = $2 AND SectorCode = 'JP'`

	rows, err := s.db.QueryContext(ctx, query, gender, nation)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var competitorIDs []int32
	for rows.Next() {
		var competitorID int32
		if err := rows.Scan(&competitorID); err != nil {
			return nil, err
		}
		competitorIDs = append(competitorIDs, competitorID)
	}

	return competitorIDs, nil
}
