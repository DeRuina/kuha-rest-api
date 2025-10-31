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

type MeasurementsStore struct {
	db *sql.DB
}

func (s *MeasurementsStore) ValidateUsersExist(ctx context.Context, userIDs []uuid.UUID) error {
	if len(userIDs) == 0 {
		return nil
	}

	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	uniqueUserIDs := make(map[uuid.UUID]bool)
	for _, id := range userIDs {
		uniqueUserIDs[id] = true
	}

	uniqueList := make([]uuid.UUID, 0, len(uniqueUserIDs))
	for id := range uniqueUserIDs {
		uniqueList = append(uniqueList, id)
	}

	query := `SELECT id FROM users WHERE id = ANY($1)`
	rows, err := s.db.QueryContext(ctx, query, pq.Array(uniqueList))
	if err != nil {
		return fmt.Errorf("failed to validate users: %w", err)
	}
	defer rows.Close()

	existingUsers := make(map[uuid.UUID]bool)
	for rows.Next() {
		var userID uuid.UUID
		if err := rows.Scan(&userID); err != nil {
			return fmt.Errorf("failed to scan user ID: %w", err)
		}
		existingUsers[userID] = true
	}

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

func (s *MeasurementsStore) InsertMeasurementsBulk(ctx context.Context, measurements []tietoevrysqlc.InsertMeasurementParams) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := tietoevrysqlc.New(tx)
	for _, m := range measurements {
		if err := q.InsertMeasurement(ctx, m); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (s *MeasurementsStore) GetMeasurementsByUser(ctx context.Context, userID uuid.UUID) ([]tietoevrysqlc.Measurement, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := tietoevrysqlc.New(s.db)
	return q.GetMeasurementsByUser(ctx, userID)
}
