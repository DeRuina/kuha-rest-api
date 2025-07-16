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

// SuuntoDataStore struct
type SuuntoDataStore struct {
	db *sql.DB
}

// Get available dates from Suunto data
func (s *SuuntoDataStore) GetDates(ctx context.Context, userID string, startDate *string, endDate *string) ([]string, error) {
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

	arg := utvsqlc.GetDatesFromSuuntoDataParams{
		UserID:     uid,
		AfterDate:  utils.NullTimeIfEmpty(start),
		BeforeDate: utils.NullTimeIfEmpty(end),
	}

	rawDates, err := queries.GetDatesFromSuuntoData(ctx, arg)
	if err != nil {
		return nil, err
	}

	var formattedDates []string
	for _, date := range rawDates {
		formattedDates = append(formattedDates, date.Format("2006-01-02"))
	}

	return formattedDates, nil
}

// Get all JSON keys (types) from Suunto data for a specific date
func (s *SuuntoDataStore) GetTypes(ctx context.Context, userID string, summaryDate string) ([]string, error) {
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

	arg := utvsqlc.GetTypesFromSuuntoDataParams{
		UserID: uid,
		Date:   date,
	}

	types, err := queries.GetTypesFromSuuntoData(ctx, arg)
	if err != nil {
		return nil, err
	}

	return types, nil
}

// Get all data for a specific date (or filter by key)
func (s *SuuntoDataStore) GetData(ctx context.Context, userID string, Date string, key *string) (json.RawMessage, error) {
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
		arg := utvsqlc.GetSpecificDataForDateSuuntoParams{
			UserID: uid,
			Date:   date,
			Key:    key,
		}
		data, err := queries.GetSpecificDataForDateSuunto(ctx, arg)
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
	arg := utvsqlc.GetAllDataForDateSuuntoParams{
		UserID: uid,
		Date:   date,
	}
	data, err := queries.GetAllDataForDateSuunto(ctx, arg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return data, nil
}

// insertData inserts Suunto data into the database
func (s *SuuntoDataStore) InsertData(ctx context.Context, userID uuid.UUID, date time.Time, data json.RawMessage) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)

	arg := utvsqlc.InsertSuuntoDataParams{
		UserID: userID,
		Date:   date,
		Data:   data,
	}

	return queries.InsertSuuntoData(ctx, arg)
}

// DeleteAllData deletes Suunto data for a specific user
func (s *SuuntoDataStore) DeleteAllData(ctx context.Context, userID uuid.UUID) (int64, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)

	return queries.DeleteAllSuuntoData(ctx, userID)
}

// GetLatestByType
func (s *SuuntoDataStore) GetLatestByType(ctx context.Context, userID uuid.UUID, typ string, limit int32) ([]LatestDataEntry, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)

	arg := utvsqlc.GetLatestSuuntoDataByTypeParams{
		UserID: userID,
		Type:   typ,
		Limit:  limit,
	}

	rows, err := queries.GetLatestSuuntoDataByType(ctx, arg)
	if err != nil {
		return nil, err
	}

	var entries []LatestDataEntry
	for _, row := range rows {
		entries = append(entries, LatestDataEntry{
			Device: "suunto",
			Date:   row.SummaryDate,
			Data:   row.Data,
		})
	}

	return entries, nil
}

// GetAllByType
func (s *SuuntoDataStore) GetAllByType(ctx context.Context, userID uuid.UUID, typ string, after, before *time.Time) ([]LatestDataEntry, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := utvsqlc.New(s.db)

	arg := utvsqlc.GetDataByTypeSuuntoParams{
		UserID:     userID,
		Type:       typ,
		AfterDate:  utils.NullTimeIfEmpty(after),
		BeforeDate: utils.NullTimeIfEmpty(before),
	}

	rows, err := queries.GetDataByTypeSuunto(ctx, arg)
	if err != nil {
		return nil, err
	}

	var result []LatestDataEntry
	for _, row := range rows {
		result = append(result, LatestDataEntry{
			Device: "suunto",
			Date:   row.SummaryDate,
			Data:   row.Data,
		})
	}

	return result, nil
}
