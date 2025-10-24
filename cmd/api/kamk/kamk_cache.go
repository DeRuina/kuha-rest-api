package kamkapi

import (
	"context"
	"fmt"
	"time"

	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
)

const (
	KAMKCacheTTL = 6 * time.Hour

	kamkInjuryListPrefix  = "kamk:injury:list"
	kamkQueriesListPrefix = "kamk:queries:list"
	kamkQuizDonePrefix    = "kamk:is-done"
)

func invalidateKamkInjuries(ctx context.Context, c *cache.Storage, sporttiID int32) {
	if c == nil {
		return
	}
	_ = c.DeleteByPrefixes(ctx, fmt.Sprintf("%s:%d", kamkInjuryListPrefix, sporttiID))
}

func invalidateKamkQueries(ctx context.Context, c *cache.Storage, sporttiID int32) {
	if c == nil {
		return
	}
	_ = c.DeleteByPrefixes(ctx, fmt.Sprintf("%s:%d", kamkQueriesListPrefix, sporttiID), fmt.Sprintf("%s:%d:", kamkQuizDonePrefix, sporttiID))
}
