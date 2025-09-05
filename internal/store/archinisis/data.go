package archinisis

import (
	"context"
	"database/sql"

	archsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/archinisis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type DataStore struct {
	db *sql.DB
}

func (s *DataStore) GetRaceReportSessions(ctx context.Context, sporttiID string) ([]int32, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := archsqlc.New(s.db)

	rows, err := q.GetRaceReportSessionIDsBySporttiID(ctx, sql.NullString{String: sporttiID, Valid: true})
	if err != nil {
		return nil, err
	}

	out := make([]int32, 0, len(rows))
	for _, v := range rows {
		if v.Valid {
			out = append(out, v.Int32)
		}
	}
	return out, nil
}

func (s *DataStore) GetRaceReport(ctx context.Context, sporttiID string, sessionID int32) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := archsqlc.New(s.db)

	res, err := q.GetRaceReport(ctx, archsqlc.GetRaceReportParams{
		SporttiID: sql.NullString{String: sporttiID, Valid: true},
		SessionID: sql.NullInt32{Int32: sessionID, Valid: true},
	})
	if err != nil {
		return "", err
	}
	if !res.Valid {
		return "", sql.ErrNoRows
	}
	return res.String, nil
}
