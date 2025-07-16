package fis

import (
	"context"
	"database/sql"

	fissqlc "github.com/DeRuina/KUHA-REST-API/internal/db/fis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

// CompetitorsStore struct
type CompetitorsStore struct {
	db *sql.DB
}

// GetAthletesBySector

type GetBySectorResponse struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	FisCode   int32  `json:"fis_code"`
}

func (s *CompetitorsStore) GetAthletesBySector(ctx context.Context, sectorCode string) ([]GetBySectorResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := fissqlc.New(s.db)

	dbSectorCode := sql.NullString{String: sectorCode, Valid: true}

	competitors, err := queries.GetAthletesBySector(ctx, dbSectorCode)
	if err != nil {
		return nil, err
	}

	var response []GetBySectorResponse
	for _, c := range competitors {
		response = append(response, GetBySectorResponse{
			FirstName: c.Firstname.String,
			LastName:  c.Lastname.String,
			FisCode:   c.Fiscode.Int32,
		})
	}

	return response, nil
}

// GetNationsBySector

func (s *CompetitorsStore) GetNationsBySector(ctx context.Context, sectorCode string) ([]string, error) {
	queries := fissqlc.New(s.db)

	dbSectorCode := sql.NullString{String: sectorCode, Valid: true}

	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	nations, err := queries.GetNationsBySector(ctx, dbSectorCode)
	if err != nil {
		return nil, err
	}

	var response []string
	for _, n := range nations {
		if n.Valid {
			response = append(response, n.String)
		}
	}

	return response, nil
}
