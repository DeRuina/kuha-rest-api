package utv

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"

	utvsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/utv"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

// SuuntoDataStore struct
type SuuntoDataStore struct {
	db *sql.DB
}

// Get available dates from Suunto data
func (s *SuuntoDataStore) GetDates(ctx context.Context, userID string, startDate *string, endDate *string) ([]string, error) {
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

	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	rawDates, err := queries.GetDatesFromSuuntoData(ctx, arg)
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

// Get all JSON keys (types) from Suunto data for a specific date
func (s *SuuntoDataStore) GetTypes(ctx context.Context, userID string, summaryDate string) ([]string, error) {
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

	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	types, err := queries.GetTypesFromSuuntoData(ctx, arg)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			return nil, utils.ErrQueryTimeOut
		}
		return nil, err
	}

	return types, nil
}

// Get all data for a specific date (or filter by key)
func (s *SuuntoDataStore) GetData(ctx context.Context, userID string, Date string, key *string) (json.RawMessage, error) {
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
	arg := utvsqlc.GetAllDataForDateSuuntoParams{
		UserID: uid,
		Date:   date,
	}
	data, err := queries.GetAllDataForDateSuunto(ctx, arg)
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
