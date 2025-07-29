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

type QuestionnairesStore struct {
	db *sql.DB
}

func NewQuestionnairesStore(db *sql.DB) *QuestionnairesStore {
	return &QuestionnairesStore{db: db}
}

func (s *QuestionnairesStore) ValidateUsersExist(ctx context.Context, userIDs []uuid.UUID) error {
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

func (s *QuestionnairesStore) InsertQuestionnaireAnswersBulk(ctx context.Context, answers []tietoevrysqlc.InsertQuestionnaireAnswerParams) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := tietoevrysqlc.New(tx)

	for _, a := range answers {
		if err := q.InsertQuestionnaireAnswer(ctx, a); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (s *QuestionnairesStore) GetQuestionnairesByUser(ctx context.Context, userID uuid.UUID) ([]tietoevrysqlc.QuestionAnswer, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := tietoevrysqlc.New(s.db)
	return q.GetQuestionnairesByUser(ctx, userID)
}
