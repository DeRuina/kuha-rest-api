package utv

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	utvsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/utv"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
	"github.com/google/uuid"
)

// OuraDataStore struct
type OuraDataStore struct {
	db *sql.DB
}

// Get available dates from Oura data
func (s *OuraDataStore) GetDates(ctx context.Context, userID string, startDate *string, endDate *string) ([]string, error) {
	queries := utvsqlc.New(s.db)

	uid, err := utils.ParseUUID(userID)
	if err != nil {
		return nil, err
	}

	var start, end *time.Time
	start, err = utils.ParseDatePtr(startDate)
	if err != nil {
		return nil, err
	}

	end, err = utils.ParseDatePtr(endDate)
	if err != nil {
		return nil, err
	}

	arg := utvsqlc.GetDatesFromOuraDataParams{
		UserID:     uid,
		AfterDate:  utils.NullTimeIfEmpty(start),
		BeforeDate: utils.NullTimeIfEmpty(end),
	}

	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	rawDates, err := queries.GetDatesFromOuraData(ctx, arg)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, utils.ErrQueryTimeOut
		}
		return nil, err
	}

	var formattedDates []string
	for _, date := range rawDates {
		formattedDates = append(formattedDates, date.Format("2006-01-02"))
	}

	return formattedDates, nil
}

// Get all JSON keys (types) from Oura data for a specific date
func (s *OuraDataStore) GetTypes(ctx context.Context, userID string, summaryDate string) ([]string, error) {
	queries := utvsqlc.New(s.db)

	uid, err := utils.ParseUUID(userID)
	if err != nil {
		return nil, err
	}
	date, err := utils.ParseDate(summaryDate)
	if err != nil {
		return nil, err
	}

	arg := utvsqlc.GetTypesFromOuraDataParams{
		UserID: uid,
		Date:   date,
	}

	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	types, err := queries.GetTypesFromOuraData(ctx, arg)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, utils.ErrQueryTimeOut
		}
		return nil, err
	}

	return types, nil
}

// Get all data for a specific date (or filter by key)
func (s *OuraDataStore) GetData(ctx context.Context, userID string, Date string, key *string) (json.RawMessage, error) {
	queries := utvsqlc.New(s.db)

	uid, err := utils.ParseUUID(userID)
	if err != nil {
		return nil, err
	}
	date, err := utils.ParseDate(Date)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	// If a key is provided, fetch only that specific type
	if key != nil {
		arg := utvsqlc.GetSpecificDataForDateOuraParams{
			UserID: uid,
			Date:   date,
			Key:    key,
		}
		data, err := queries.GetSpecificDataForDateOura(ctx, arg)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
			if errors.Is(err, context.DeadlineExceeded) {
				return nil, utils.ErrQueryTimeOut
			}
			return nil, err
		}

		switch v := data.(type) {
		case nil:
			return nil, nil
		case string:
			return json.RawMessage(v), nil
		case []byte:
			return json.RawMessage(v), nil
		default:
			return json.Marshal(data)

		}
	}

	// If no key is provided, return all data
	arg := utvsqlc.GetAllDataForDateOuraParams{
		UserID: uid,
		Date:   date,
	}
	data, err := queries.GetAllDataForDateOura(ctx, arg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, utils.ErrQueryTimeOut
		}
		return nil, err
	}

	return data, nil
}

// insertData inserts Oura data into the database
func (s *OuraDataStore) InsertData(ctx context.Context, userID uuid.UUID, date time.Time, data json.RawMessage) error {
	queries := utvsqlc.New(s.db)

	arg := utvsqlc.InsertOuraDataParams{
		UserID: userID,
		Date:   date,
		Data:   data,
	}

	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	return queries.InsertOuraData(ctx, arg)
}

// DeleteAllData deletes Oura data for a specific user
func (s *OuraDataStore) DeleteAllData(ctx context.Context, userID uuid.UUID) (int64, error) {
	queries := utvsqlc.New(s.db)

	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	return queries.DeleteAllOuraData(ctx, userID)
}

// GetLatestByType
func (s *OuraDataStore) GetLatestByType(ctx context.Context, userID uuid.UUID, typ string, limit int32) ([]LatestDataEntry, error) {
	queries := utvsqlc.New(s.db)

	arg := utvsqlc.GetLatestOuraDataByTypeParams{
		UserID: userID,
		Type:   typ,
		Limit:  limit,
	}

	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	rows, err := queries.GetLatestOuraDataByType(ctx, arg)
	if err != nil {
		return nil, err
	}

	var entries []LatestDataEntry
	for _, row := range rows {
		entries = append(entries, LatestDataEntry{
			Device: "oura",
			Date:   row.SummaryDate,
			Data:   row.Data,
		})
	}

	return entries, nil
}
