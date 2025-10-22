package utvapi

import (
	"context"
	"fmt"
	"time"

	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
	"github.com/google/uuid"
)

const (
	UTVCacheTTL = 6 * time.Hour

	// per-source prefixes
	utvGarminDatesPrefix = "utv:garmin:dates"
	utvGarminTypesPrefix = "utv:garmin:types"
	utvGarminDataPrefix  = "utv:garmin:data"

	utvOuraDatesPrefix = "utv:oura:dates"
	utvOuraTypesPrefix = "utv:oura:types"
	utvOuraDataPrefix  = "utv:oura:data"

	utvPolarDatesPrefix = "utv:polar:dates"
	utvPolarTypesPrefix = "utv:polar:types"
	utvPolarDataPrefix  = "utv:polar:data"

	utvSuuntoDatesPrefix = "utv:suunto:dates"
	utvSuuntoTypesPrefix = "utv:suunto:types"
	utvSuuntoDataPrefix  = "utv:suunto:data"

	// cross-source/general prefixes
	utvLatestPrefix = "utv:latest"
	utvAllPrefix    = "utv:all"

	// coachtech
	utvCoachtechDataPrefix = "utv:coachtech:data:user"
)

// invalidate per-source keys + general views (latest, all)
func invalidateUTVSource(ctx context.Context, c *cache.Storage, userID uuid.UUID, src string) {
	if c == nil {
		return
	}

	var pfx []string
	switch src {
	case "garmin":
		pfx = []string{
			fmt.Sprintf("%s:%s", utvGarminDatesPrefix, userID),
			fmt.Sprintf("%s:%s", utvGarminTypesPrefix, userID),
			fmt.Sprintf("%s:%s", utvGarminDataPrefix, userID),
		}
	case "oura":
		pfx = []string{
			fmt.Sprintf("%s:%s", utvOuraDatesPrefix, userID),
			fmt.Sprintf("%s:%s", utvOuraTypesPrefix, userID),
			fmt.Sprintf("%s:%s", utvOuraDataPrefix, userID),
		}
	case "polar":
		pfx = []string{
			fmt.Sprintf("%s:%s", utvPolarDatesPrefix, userID),
			fmt.Sprintf("%s:%s", utvPolarTypesPrefix, userID),
			fmt.Sprintf("%s:%s", utvPolarDataPrefix, userID),
		}
	case "suunto":
		pfx = []string{
			fmt.Sprintf("%s:%s", utvSuuntoDatesPrefix, userID),
			fmt.Sprintf("%s:%s", utvSuuntoTypesPrefix, userID),
			fmt.Sprintf("%s:%s", utvSuuntoDataPrefix, userID),
		}
	default:
		return
	}

	// also nuke the general views for this user
	pfx = append(pfx,
		fmt.Sprintf("%s:%s:", utvLatestPrefix, userID),
		fmt.Sprintf("%s:%s:", utvAllPrefix, userID),
	)

	_ = c.DeleteByPrefixes(ctx, pfx...)
}

func invalidateUTVCoachtech(ctx context.Context, c *cache.Storage, userID uuid.UUID) {
	if c == nil {
		return
	}
	_ = c.DeleteByPrefixes(ctx,
		fmt.Sprintf("%s:%s", utvCoachtechDataPrefix, userID),
	)
}
