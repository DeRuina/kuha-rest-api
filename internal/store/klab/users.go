package klab

import (
	"context"
	"database/sql"

	klabsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/klab"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) GetAllSporttiIDs(ctx context.Context) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := klabsqlc.New(s.db)
	return queries.GetAllSporttiIDs(ctx)
}

func (s *UsersStore) GetCustomerByID(ctx context.Context, idcustomer int32) (klabsqlc.Customer, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := klabsqlc.New(s.db)
	return queries.GetCustomerByID(ctx, idcustomer)
}

func (s *UsersStore) GetCustomerIDBySporttiID(ctx context.Context, sporttiID string) (int32, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	q := klabsqlc.New(s.db)
	return q.GetCustomerIDBySporttiID(ctx, sql.NullString{String: sporttiID, Valid: true})
}
