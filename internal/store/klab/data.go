package klab

import (
	"context"
	"database/sql"

	klabsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/klab"
)

type DataStore struct {
	db *sql.DB
}

type KlabDataPayload struct {
	Customer     []klabsqlc.InsertCustomerParams
	Measurements []klabsqlc.InsertMeasurementParams
	DirTests     []klabsqlc.InsertDirTestParams
	DirTestSteps []klabsqlc.InsertDirTestStepParams
	DirRawData   []klabsqlc.InsertDirRawDataParams
	DirReports   []klabsqlc.InsertDirReportParams
	DirResults   []klabsqlc.InsertDirResultsParams
}

func (s *DataStore) InsertKlabDataBulk(ctx context.Context, data []KlabDataPayload) error {
	ctx, cancel := context.WithTimeout(ctx, DataTimeout)
	defer cancel()

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	q := klabsqlc.New(tx)

	for _, item := range data {
		// Insert all customers
		for _, c := range item.Customer {
			if err := q.InsertCustomer(ctx, c); err != nil {
				return err
			}
		}

		// Insert measurements
		for _, m := range item.Measurements {
			if err := q.InsertMeasurement(ctx, m); err != nil {
				return err
			}
		}

		// Insert dirtests
		for _, t := range item.DirTests {
			if err := q.InsertDirTest(ctx, t); err != nil {
				return err
			}
		}

		// Insert dirteststeps
		for _, step := range item.DirTestSteps {
			if err := q.InsertDirTestStep(ctx, step); err != nil {
				return err
			}
		}

		// Insert dirrawdata
		for _, raw := range item.DirRawData {
			if err := q.InsertDirRawData(ctx, raw); err != nil {
				return err
			}
		}

		// Insert dirreports
		for _, rep := range item.DirReports {
			if err := q.InsertDirReport(ctx, rep); err != nil {
				return err
			}
		}

		// Insert dirresults
		for _, res := range item.DirResults {
			if err := q.InsertDirResults(ctx, res); err != nil {
				return err
			}
		}
	}

	return tx.Commit()
}
