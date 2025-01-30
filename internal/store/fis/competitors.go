package fis

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
)

// CompetitorsStore struct
type CompetitorsStore struct {
	db *sql.DB
}

// Competitor model with sql.NullString
type Competitor struct {
	FirstName sql.NullString `json:"-"`
	LastName  sql.NullString `json:"-"`
	FisCode   int32          `json:"fis_code"`
}

// Custom JSON Marshaller for handling NULL values
func (c Competitor) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		FisCode   int32  `json:"fis_code"`
	}{
		FirstName: nullStringToString(c.FirstName),
		LastName:  nullStringToString(c.LastName),
		FisCode:   c.FisCode,
	})
}

// Helper function to convert `sql.NullString` to a regular string
func nullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

// Ensure `CompetitorsStore` implements `Competitors` interface
func (s *CompetitorsStore) GetBySector(ctx context.Context, sectorCode string) ([]Competitor, error) {
	query := `SELECT Firstname, Lastname, Fiscode FROM A_competitor WHERE SectorCode = $1 ORDER BY Fiscode`

	rows, err := s.db.QueryContext(ctx, query, sectorCode)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var competitors []Competitor
	for rows.Next() {
		var c Competitor
		err := rows.Scan(&c.FirstName, &c.LastName, &c.FisCode)
		if err != nil {
			return nil, err
		}
		competitors = append(competitors, c)
	}

	if len(competitors) == 0 {
		log.Println("⚠️ No competitors found")
	}
	return competitors, nil
}

// Implement GetByFiscodeJP
func (s *CompetitorsStore) GetByFiscodeJP(ctx context.Context, fiscode int32) (int32, error) {
	query := `SELECT CompetitorID FROM A_competitor WHERE Fiscode = $1 AND SectorCode = 'JP'`
	log.Printf("🟡 Executing SQL Query: %s | Fiscode: %d", query, fiscode)

	var competitorID int32
	err := s.db.QueryRowContext(ctx, query, fiscode).Scan(&competitorID)
	if err != nil {
		log.Printf("❌ SQL Query failed: %v", err)
		return 0, err
	}

	log.Printf("✅ Found Competitor ID: %d", competitorID)
	return competitorID, nil
}

// Implement GetByFiscodeNK
func (s *CompetitorsStore) GetByFiscodeNK(ctx context.Context, fiscode int32) (int32, error) {
	query := `SELECT CompetitorID FROM A_competitor WHERE Fiscode = $1 AND SectorCode = 'NK'`
	log.Printf("🟡 Executing SQL Query: %s | Fiscode: %d", query, fiscode)

	var competitorID int32
	err := s.db.QueryRowContext(ctx, query, fiscode).Scan(&competitorID)
	if err != nil {
		log.Printf("❌ SQL Query failed: %v", err)
		return 0, err
	}

	log.Printf("✅ Found Competitor ID: %d", competitorID)
	return competitorID, nil
}

// Implement GetByGenderAndNationJP
func (s *CompetitorsStore) GetByGenderAndNationJP(ctx context.Context, gender, nation string) ([]int32, error) {
	query := `SELECT CompetitorID FROM A_competitor WHERE Gender = $1 AND NationCode = $2 AND SectorCode = 'JP'`
	log.Printf("🟡 Executing SQL Query: %s | Gender: %s, Nation: %s", query, gender, nation)

	rows, err := s.db.QueryContext(ctx, query, gender, nation)
	if err != nil {
		log.Printf("❌ SQL Query failed: %v", err)
		return nil, err
	}
	defer rows.Close()

	var competitorIDs []int32
	for rows.Next() {
		var competitorID int32
		if err := rows.Scan(&competitorID); err != nil {
			log.Printf("❌ Row Scan failed: %v", err)
			return nil, err
		}
		competitorIDs = append(competitorIDs, competitorID)
	}

	log.Printf("✅ Found %d Competitors", len(competitorIDs))
	return competitorIDs, nil
}
