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

type SymptomsStore struct {
	db *sql.DB
}

func NewSymptomsStore(db *sql.DB) *SymptomsStore {
	return &SymptomsStore{db: db}
}

func (s *SymptomsStore) ValidateUsersExist(ctx context.Context, userIDs []uuid.UUID) error {
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

func (s *SymptomsStore) InsertSymptomsBulk(ctx context.Context, symptoms []tietoevrysqlc.InsertSymptomParams) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := tietoevrysqlc.New(tx)

	for _, symptom := range symptoms {
		if err := q.InsertSymptom(ctx, symptom); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (s *SymptomsStore) GetSymptomsByUser(ctx context.Context, userID uuid.UUID) ([]tietoevrysqlc.Symptom, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	return tietoevrysqlc.New(s.db).GetSymptomsByUser(ctx, userID)
}
