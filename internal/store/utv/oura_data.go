package utv

import (
	"context"
	"database/sql"
	"time"

	utvsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/utv"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
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

	rawDates, err := queries.GetDatesFromOuraData(ctx, arg)
	if err != nil {
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

	// Validate and convert inputs
	uid, err := utils.ParseUUID(userID)
	if err != nil {
		return nil, err
	}
	date, err := utils.ParseDate(summaryDate)
	if err != nil {
		return nil, err
	}

	// Call the SQLC-generated query
	arg := utvsqlc.GetTypesFromOuraDataParams{
		UserID: uid,
		Date:   date,
	}

	types, err := queries.GetTypesFromOuraData(ctx, arg)
	if err != nil {
		return nil, err
	}

	return types, nil
}

// // Get all data for a specific date (or filter by type)
// func (s *OuraDataStore) GetData(ctx context.Context, userID string, summaryDate string, key *string) (interface{}, error) {
// 	queries := utvsql.New(s.db)

// 	uid, err := utils.ParseUUID(userID)
// 	if err != nil {
// 		return nil, err
// 	}

// 	date, err := utils.ParseDate(summaryDate)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if key != nil {
// 		arg := utvsql.GetSpecificDataForDateParams{
// 			UserID:      uid,
// 			SummaryDate: date,
// 			Column3:     *key,
// 		}
// 		return queries.GetSpecificDataForDate(ctx, arg)
// 	}

// 	arg := utvsql.GetAllDataForDateParams{
// 		UserID:      uid,
// 		SummaryDate: date,
// 	}

// 	data, err := queries.GetAllDataForDate(ctx, arg)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var jsonData map[string]interface{}
// 	if err := json.Unmarshal(data, &jsonData); err != nil {
// 		return nil, err
// 	}

// 	return jsonData, nil
// }
