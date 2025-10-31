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

// PolarDataStore struct
type PolarDataStore struct {
	db *sql.DB
}

// Get available dates from Polar data
func (s *PolarDataStore) GetDates(ctx context.Context, userID string, startDate *string, endDate *string) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

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

	arg := utvsqlc.GetDatesFromPolarDataParams{
		UserID:     uid,
		AfterDate:  utils.NullTimeIfEmpty(start),
		BeforeDate: utils.NullTimeIfEmpty(end),
	}

	rawDates, err := queries.GetDatesFromPolarData(ctx, arg)
	if err != nil {
		return nil, err
	}

	var formattedDates []string
	for _, date := range rawDates {
		formattedDates = append(formattedDates, date.Format("2006-01-02"))
	}

	return formattedDates, nil
}

// Get all JSON keys (types) from Polar data for a specific date
func (s *PolarDataStore) GetTypes(ctx context.Context, userID string, summaryDate string) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)

	uid, err := utils.ParseUUID(userID)
	if err != nil {
		return nil, err
	}
	date, err := utils.ParseDate(summaryDate)
	if err != nil {
		return nil, err
	}

	arg := utvsqlc.GetTypesFromPolarDataParams{
		UserID: uid,
		Date:   date,
	}

	types, err := queries.GetTypesFromPolarData(ctx, arg)
	if err != nil {
		return nil, err
	}

	return types, nil
}

// Get all data for a specific date (or filter by key)
func (s *PolarDataStore) GetData(ctx context.Context, userID string, Date string, key *string) (json.RawMessage, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)

	uid, err := utils.ParseUUID(userID)
	if err != nil {
		return nil, err
	}
	date, err := utils.ParseDate(Date)
	if err != nil {
		return nil, err
	}
	// If a key is provided, fetch only that specific type
	if key != nil {
		arg := utvsqlc.GetSpecificDataForDatePolarParams{
			UserID: uid,
			Date:   date,
			Key:    key,
		}
		data, err := queries.GetSpecificDataForDatePolar(ctx, arg)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
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
	arg := utvsqlc.GetAllDataForDatePolarParams{
		UserID: uid,
		Date:   date,
	}
	data, err := queries.GetAllDataForDatePolar(ctx, arg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return data, nil
}

// insertData inserts Polar data into the database
func (s *PolarDataStore) InsertData(ctx context.Context, userID uuid.UUID, date time.Time, data json.RawMessage) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)

	arg := utvsqlc.InsertPolarDataParams{
		UserID: userID,
		Date:   date,
		Data:   data,
	}

	return queries.InsertPolarData(ctx, arg)
}

// DeleteAllData deletes Polar data for a specific user
func (s *PolarDataStore) DeleteAllData(ctx context.Context, userID uuid.UUID) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)

	return queries.DeleteAllPolarData(ctx, userID)
}

// GetLatestByType
func (s *PolarDataStore) GetLatestByType(ctx context.Context, userID uuid.UUID, typ string, limit int32) ([]LatestDataEntry, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)

	arg := utvsqlc.GetLatestPolarDataByTypeParams{
		UserID: userID,
		Type:   typ,
		Limit:  limit,
	}

	rows, err := queries.GetLatestPolarDataByType(ctx, arg)
	if err != nil {
		return nil, err
	}

	var entries []LatestDataEntry
	for _, row := range rows {
		entries = append(entries, LatestDataEntry{
			Device: "polar",
			Date:   row.SummaryDate,
			Data:   row.Data,
		})
	}

	return entries, nil
}

// GetAllByType
func (s *PolarDataStore) GetAllByType(ctx context.Context, userID uuid.UUID, typ string, after, before *time.Time, limit, offset int32) ([]LatestDataEntry, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)

	arg := utvsqlc.GetDataByTypePolarParams{
		UserID:     userID,
		Type:       typ,
		AfterDate:  utils.NullTimeIfEmpty(after),
		BeforeDate: utils.NullTimeIfEmpty(before),
		Limit:      limit,
		Offset:     offset,
	}

	rows, err := queries.GetDataByTypePolar(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result []LatestDataEntry
	for _, row := range rows {
		result = append(result, LatestDataEntry{
			Device: "polar",
			Date:   row.SummaryDate,
			Data:   row.Data,
		})
	}

	return result, nil
}
