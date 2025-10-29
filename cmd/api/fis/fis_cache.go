package fisapi

import (
	"context"
	"fmt"
	"time"

	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
)

const (
	FISCacheTTL = 6 * time.Hour

	fisAthletesPrefix   = "fis:athletes"
	fisLastRowPrefix    = "fis:lastrow"
	fisCompetitorPrefix = "fis:competitor"
	fisNationsPrefix    = "fis:nations"
)

func invalidateCompetitor(ctx context.Context, c *cache.Storage, competitorID int32) {
	if c == nil {
		return
	}
	_ = c.DeleteByPrefixes(
		ctx,
		fmt.Sprintf("%s", fisLastRowPrefix),
		fmt.Sprintf("%s:%d", fisCompetitorPrefix, competitorID),
	)
}

func invalidateAthletesSector(ctx context.Context, c *cache.Storage, sector string) {
	if c == nil {
		return
	}
	_ = c.DeleteByPrefixes(
		ctx,
		fmt.Sprintf("%s:%s", fisAthletesPrefix, sector),
	)
}

func invalidateNationsSector(ctx context.Context, c *cache.Storage, sector string) {
	if c == nil {
		return
	}
	_ = c.DeleteByPrefixes(ctx, fmt.Sprintf("%s:%s", fisNationsPrefix, sector))
}

func invalidateSector(ctx context.Context, c *cache.Storage, sector string) {
	invalidateAthletesSector(ctx, c, sector)
	invalidateNationsSector(ctx, c, sector)
}
