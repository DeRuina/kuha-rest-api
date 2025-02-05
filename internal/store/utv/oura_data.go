package utv

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	utvsql "github.com/DeRuina/KUHA-REST-API/internal/db/utv"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

// OuraDataStore struct
type OuraDataStore struct {
	db *sql.DB
}

// Get available dates for Oura data
func (s *OuraDataStore) GetDates(ctx context.Context, userID string, startDate string, endDate string) ([]time.Time, error) {
	queries := utvsql.New(s.db)

	// Validate and convert inputs
	uid, err := utils.ParseUUID(userID)
	if err != nil {
		return nil, err
	}
	start, err := utils.ParseDate(startDate)
	if err != nil {
		return nil, err
	}
	end, err := utils.ParseDate(endDate)
	if err != nil {
		return nil, err
	}

	arg := utvsql.GetDatesFromOuraDataParams{
		UserID:        uid,
		SummaryDate:   start,
		SummaryDate_2: end,
	}

	return queries.GetDatesFromOuraData(ctx, arg)
}

// Get all JSON keys from Oura data
func (s *OuraDataStore) GetTypes(ctx context.Context, userID string, summaryDate string) ([]string, error) {
	queries := utvsql.New(s.db)

	// Validate and convert inputs
	uid, err := utils.ParseUUID(userID)
	if err != nil {
		return nil, err
	}
	date, err := utils.ParseDate(summaryDate)
	if err != nil {
		return nil, err
	}

	arg := utvsql.GetTypesFromOuraDataParams{
		SummaryDate: date,
		UserID:      uid,
	}

	types, err := queries.GetTypesFromOuraData(ctx, arg)
	if err != nil {
		return nil, err
	}

	return types, nil
}

// Get a specific data point from Oura data
func (s *OuraDataStore) GetDataPoint(ctx context.Context, userID string, summaryDate string, key string) (interface{}, error) {
	queries := utvsql.New(s.db)

	// Validate and convert inputs
	uid, err := utils.ParseUUID(userID)
	if err != nil {
		return nil, err
	}
	date, err := utils.ParseDate(summaryDate)
	if err != nil {
		return nil, err
	}

	arg := utvsql.GetDataPointFromOuraDataParams{
		Data:        json.RawMessage(utils.ParseJSONKey(key)),
		SummaryDate: date,
		UserID:      uid,
	}

	data, err := queries.GetDataPointFromOuraData(ctx, arg)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Get unique JSON keys from Oura data over a date range
func (s *OuraDataStore) GetUniqueTypes(ctx context.Context, userID string, startDate string, endDate string) ([]string, error) {
	queries := utvsql.New(s.db)

	// Validate and convert inputs
	uid, err := utils.ParseUUID(userID)
	if err != nil {
		return nil, err
	}
	start, err := utils.ParseTimestamp(startDate)
	if err != nil {
		return nil, err
	}
	end, err := utils.ParseTimestamp(endDate)
	if err != nil {
		return nil, err
	}

	arg := utvsql.GetUniqueOuraDataTypesParams{
		UserID:        uid,
		ToTimestamp:   start,
		ToTimestamp_2: end,
	}

	types, err := queries.GetUniqueOuraDataTypes(ctx, arg)
	if err != nil {
		return nil, err
	}

	return types, nil
}
