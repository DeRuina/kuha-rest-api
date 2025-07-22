package tietoevry

import (
	"context"
	"database/sql"

	tietoevrysqlc "github.com/DeRuina/KUHA-REST-API/internal/db/tietoevry"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type ExercisePayload struct {
	Exercise tietoevrysqlc.InsertExerciseParams
	HRZones  []tietoevrysqlc.InsertExerciseHRZoneParams
	Samples  []tietoevrysqlc.InsertExerciseSampleParams
	Sections []tietoevrysqlc.InsertExerciseSectionParams
}

type ExercisesStore struct {
	db *sql.DB
}

func NewExercisesStore(db *sql.DB) *ExercisesStore {
	return &ExercisesStore{
		db: db,
	}
}

func (s *ExercisesStore) InsertExercisesBulk(ctx context.Context, exercises []ExercisePayload) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	// Start a transaction to ensure atomicity for ALL exercises
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := tietoevrysqlc.New(tx)

	// Process each exercise bundle
	for _, exercise := range exercises {
		// 1. Insert base exercise
		if err := q.InsertExercise(ctx, exercise.Exercise); err != nil {
			return err
		}

		// 2. Insert HR zones
		for _, zone := range exercise.HRZones {
			if err := q.InsertExerciseHRZone(ctx, zone); err != nil {
				return err
			}
		}

		// 3. Insert samples
		for _, sample := range exercise.Samples {
			if err := q.InsertExerciseSample(ctx, sample); err != nil {
				return err
			}
		}

		// 4. Insert sections
		for _, section := range exercise.Sections {
			if err := q.InsertExerciseSection(ctx, section); err != nil {
				return err
			}
		}
	}

	// Commit the transaction if all exercises succeeded
	return tx.Commit()
}
