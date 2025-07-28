package tietoevry

import (
	"context"
	"database/sql"
	"fmt"

	tietoevrysqlc "github.com/DeRuina/KUHA-REST-API/internal/db/tietoevry"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
	"github.com/google/uuid"
	"github.com/lib/pq"
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

func (s *ExercisesStore) ValidateUsersExist(ctx context.Context, userIDs []uuid.UUID) error {
	if len(userIDs) == 0 {
		return nil
	}

	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	// Get unique user IDs
	uniqueUserIDs := make(map[uuid.UUID]bool)
	for _, id := range userIDs {
		uniqueUserIDs[id] = true
	}

	uniqueList := make([]uuid.UUID, 0, len(uniqueUserIDs))
	for id := range uniqueUserIDs {
		uniqueList = append(uniqueList, id)
	}

	// Query to get existing user IDs
	query := `SELECT id FROM users WHERE id = ANY($1)`
	rows, err := s.db.QueryContext(ctx, query, pq.Array(uniqueList))
	if err != nil {
		return fmt.Errorf("failed to validate users: %w", err)
	}
	defer rows.Close()

	// Collect existing user IDs
	existingUsers := make(map[uuid.UUID]bool)
	for rows.Next() {
		var userID uuid.UUID
		if err := rows.Scan(&userID); err != nil {
			return fmt.Errorf("failed to scan user ID: %w", err)
		}
		existingUsers[userID] = true
	}

	// Find missing users
	var missingUsers []uuid.UUID
	for _, userID := range uniqueList {
		if !existingUsers[userID] {
			missingUsers = append(missingUsers, userID)
		}
	}

	if len(missingUsers) > 0 {
		return fmt.Errorf("users do not exist, please create them first: %v", missingUsers)
	}

	return nil
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

func (s *ExercisesStore) GetExercisesByUser(ctx context.Context, userID uuid.UUID) ([]tietoevrysqlc.Exercise, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()
	return tietoevrysqlc.New(s.db).GetExercisesByUser(ctx, userID)
}

func (s *ExercisesStore) GetExerciseHRZones(ctx context.Context, id uuid.UUID) ([]tietoevrysqlc.ExerciseHrZone, error) {
	return tietoevrysqlc.New(s.db).GetExerciseHRZones(ctx, id)
}

func (s *ExercisesStore) GetExerciseSamples(ctx context.Context, id uuid.UUID) ([]tietoevrysqlc.ExerciseSample, error) {
	return tietoevrysqlc.New(s.db).GetExerciseSamples(ctx, id)
}

func (s *ExercisesStore) GetExerciseSections(ctx context.Context, id uuid.UUID) ([]tietoevrysqlc.ExerciseSection, error) {
	return tietoevrysqlc.New(s.db).GetExerciseSections(ctx, id)
}
