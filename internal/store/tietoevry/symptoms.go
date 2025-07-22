package tietoevry

import (
	"context"
	"database/sql"

	tietoevrysqlc "github.com/DeRuina/KUHA-REST-API/internal/db/tietoevry"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type SymptomsStore struct {
	db *sql.DB
}

func NewSymptomsStore(db *sql.DB) *SymptomsStore {
	return &SymptomsStore{db: db}
}

func (s *SymptomsStore) InsertSymptomsBulk(ctx context.Context, symptoms []tietoevrysqlc.InsertSymptomParams) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	qtx := tietoevrysqlc.New(tx)

	for _, symptom := range symptoms {
		if err := qtx.InsertSymptom(ctx, symptom); err != nil {
			return err
		}
	}

	return tx.Commit()
}
