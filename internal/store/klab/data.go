package klab

import (
	"context"
	"database/sql"

	klabsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/klab"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type DataStore struct {
	db *sql.DB
}

type KlabDataPayload struct {
	Customers    []klabsqlc.UpsertCustomerParams
	Measurements []klabsqlc.InsertMeasurementParams
	DirTests     []klabsqlc.InsertDirTestParams
	DirTestSteps []klabsqlc.InsertDirTestStepParams
	DirRawData   []klabsqlc.InsertDirRawDataParams
	DirReports   []klabsqlc.InsertDirReportParams
	DirResults   []klabsqlc.InsertDirResultsParams
}

type KlabDataNoCustomer struct {
	CustomerID   int32
	Measurements []klabsqlc.MeasurementList
	DirTests     []klabsqlc.Dirtest
	DirTestSteps []klabsqlc.Dirteststep
	DirReports   []klabsqlc.Dirreport
	DirRawData   []klabsqlc.Dirrawdatum
	DirResults   []klabsqlc.Dirresult
}

func (s *DataStore) InsertKlabDataBulk(ctx context.Context, payloads []KlabDataPayload) error {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := klabsqlc.New(tx)

	for _, b := range payloads {
		// 1) Customers
		for _, c := range b.Customers {
			if err := q.UpsertCustomer(ctx, c); err != nil {
				return err
			}
		}

		// 2) Measurements
		for _, m := range b.Measurements {
			if err := q.InsertMeasurement(ctx, m); err != nil {
				return err
			}
		}

		// 3) Child tables
		for _, t := range b.DirTests {
			if err := q.InsertDirTest(ctx, t); err != nil {
				return err
			}
		}
		for _, st := range b.DirTestSteps {
			if err := q.InsertDirTestStep(ctx, st); err != nil {
				return err
			}
		}
		for _, rd := range b.DirRawData {
			if err := q.InsertDirRawData(ctx, rd); err != nil {
				return err
			}
		}
		for _, rp := range b.DirReports {
			if err := q.InsertDirReport(ctx, rp); err != nil {
				return err
			}
		}
		for _, rs := range b.DirResults {
			if err := q.InsertDirResults(ctx, rs); err != nil {
				return err
			}
		}
	}

	return tx.Commit()
}

func (s *DataStore) GetCustomerByID(ctx context.Context, idcustomer int32) (klabsqlc.Customer, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := klabsqlc.New(s.db)
	return queries.GetCustomerByID(ctx, idcustomer)
}

func (s *DataStore) GetDataByCustomerIDNoCustomer(ctx context.Context, idcustomer int32) (*KlabDataNoCustomerResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	// Check if customer exists first
	_, err := s.GetCustomerByID(ctx, idcustomer)
	if err != nil {
		return nil, err
	}

	q := klabsqlc.New(s.db)

	// Get all measurements for the customer
	meas, err := q.GetMeasurementsByCustomer(ctx, idcustomer)
	if err != nil {
		return nil, err
	}

	// Collect measurement IDs for bulk fetches
	mids := make([]int32, 0, len(meas))
	for _, m := range meas {
		mids = append(mids, m.Idmeasurement)
	}

	var (
		tests   []klabsqlc.Dirtest
		steps   []klabsqlc.Dirteststep
		reps    []klabsqlc.Dirreport
		raws    []klabsqlc.Dirrawdatum
		results []klabsqlc.Dirresult
	)

	if len(mids) > 0 {
		if tests, err = q.GetDirTestsByMeasurementIDs(ctx, mids); err != nil {
			return nil, err
		}
		if steps, err = q.GetDirTestStepsByMeasurementIDs(ctx, mids); err != nil {
			return nil, err
		}
		if reps, err = q.GetDirReportsByMeasurementIDs(ctx, mids); err != nil {
			return nil, err
		}
		if raws, err = q.GetDirRawDataByMeasurementIDs(ctx, mids); err != nil {
			return nil, err
		}
		if results, err = q.GetDirResultsByMeasurementIDs(ctx, mids); err != nil {
			return nil, err
		}
	}

	// Convert to clean response structs
	cleanMeasurements := make([]KlabMeasurementResponse, len(meas))
	for i, m := range meas {
		cleanMeasurements[i] = convertMeasurement(m)
	}

	cleanTests := make([]KlabDirTestResponse, len(tests))
	for i, t := range tests {
		cleanTests[i] = convertDirTest(t)
	}

	cleanSteps := make([]KlabDirTestStepResponse, len(steps))
	for i, s := range steps {
		cleanSteps[i] = convertDirTestStep(s)
	}

	cleanReports := make([]KlabDirReportResponse, len(reps))
	for i, r := range reps {
		cleanReports[i] = convertDirReport(r)
	}

	cleanRawData := make([]KlabDirRawDataResponse, len(raws))
	for i, r := range raws {
		cleanRawData[i] = convertDirRawData(r)
	}

	cleanResults := make([]KlabDirResultsResponse, len(results))
	for i, r := range results {
		cleanResults[i] = convertDirResults(r)
	}

	return &KlabDataNoCustomerResponse{
		CustomerID:   idcustomer,
		Measurements: cleanMeasurements,
		DirTests:     cleanTests,
		DirTestSteps: cleanSteps,
		DirReports:   cleanReports,
		DirRawData:   cleanRawData,
		DirResults:   cleanResults,
	}, nil
}

func (s *DataStore) GetCustomerIDBySporttiID(ctx context.Context, sporttiID string) (int32, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := klabsqlc.New(s.db)
	return q.GetCustomerIDBySporttiID(ctx, sql.NullString{String: sporttiID, Valid: true})
}
