package archinisis

import (
	"context"
	"database/sql"

	archsqlc "github.com/DeRuina/KUHA-REST-API/internal/db/archinisis"
	"github.com/DeRuina/KUHA-REST-API/internal/utils"
)

type UsersStore struct {
	db *sql.DB
}

func (s *UsersStore) GetAllSporttiIDs(ctx context.Context) ([]string, error) {
	ctx, cancel := context.WithTimeout(ctx, utils.QueryTimeout)
	defer cancel()

	queries := archsqlc.New(s.db)
	return queries.GetSporttiIDs(ctx)
}
