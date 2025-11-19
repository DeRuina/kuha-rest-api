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

	fisRaceCCLastRowPrefix = "fis:lastrow:racecc"
	fisRaceCCCodesPrefix   = "fis:racecc:codes"
	fisRaceCCListPrefix    = "fis:racecc:list"

	fisRaceJPLastRowPrefix = "fis:lastrow:racejp"
	fisRaceJPCodesPrefix   = "fis:racejp:codes"
	fisRaceJPListPrefix    = "fis:racejp:list"

	fisRaceNKLastRowPrefix = "fis:lastrow:racenk"
	fisRaceNKCodesPrefix   = "fis:racenk:codes"
	fisRaceNKListPrefix    = "fis:racenk:list"

	fisResultCCLastRowPrefix = "fis:lastrow:resultcc"
	fisResultCCRacePrefix    = "fis:resultcc:race"
	fisResultCCAthletePrefix = "fis:resultcc:athlete"

	fisResultJPLastRowPrefix = "fis:lastrow:resultjp"
	fisResultJPRacePrefix    = "fis:resultjp:race"
	fisResultJPAthletePrefix = "fis:resultjp:athlete"
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

func invalidateRaceCC(ctx context.Context, c *cache.Storage, raceID int32) {
	if c == nil {
		return
	}
	_ = c.DeleteByPrefixes(ctx, fisRaceCCLastRowPrefix, fisRaceCCListPrefix)
}

func invalidateRaceJP(ctx context.Context, c *cache.Storage, raceID int32) {
	if c == nil {
		return
	}
	_ = c.DeleteByPrefixes(ctx, fisRaceJPLastRowPrefix, fisRaceJPListPrefix)
}

func invalidateRaceNK(ctx context.Context, c *cache.Storage, raceID int32) {
	if c == nil {
		return
	}
	_ = c.DeleteByPrefixes(ctx, fisRaceNKLastRowPrefix, fisRaceNKListPrefix)
}

func invalidateResultCC(ctx context.Context, c *cache.Storage, recid int32) {
	if c == nil {
		return
	}
	_ = c.DeleteByPrefixes(
		ctx,
		fisResultCCLastRowPrefix,
		fisResultCCRacePrefix,
		fisResultCCAthletePrefix,
	)
}

func invalidateResultJP(ctx context.Context, c *cache.Storage, recid int32) {
	if c == nil {
		return
	}
	_ = c.DeleteByPrefixes(
		ctx,
		fisResultJPLastRowPrefix,
		fisResultJPRacePrefix,
		fisResultJPAthletePrefix,
	)
}
