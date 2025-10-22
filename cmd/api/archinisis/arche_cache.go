package archapi

import (
	"context"
	"fmt"
	"time"

	"github.com/DeRuina/KUHA-REST-API/internal/store/cache"
)

const (
	ARCHCacheTTL = 6 * time.Hour

	archSessionsPrefix = "arch:race-report:sessions"
	archHTMLPrefix     = "arch:race-report:html"
	archDataPrefix     = "arch:data"
)

func invalidateArchRaceReport(ctx context.Context, c *cache.Storage, sporttiID string, sessionID *int32) {
	if c == nil {
		return
	}
	pf := []string{
		fmt.Sprintf("%s:%s", archSessionsPrefix, sporttiID),
	}
	if sessionID != nil {
		// just this one HTML
		pf = append(pf, fmt.Sprintf("%s:%s:%d", archHTMLPrefix, sporttiID, *sessionID))
	} else {
		// all HTML for this athlete
		pf = append(pf, fmt.Sprintf("%s:%s:", archHTMLPrefix, sporttiID))
	}
	_ = c.DeleteByPrefixes(ctx, pf...)
}

func invalidateArchData(ctx context.Context, c *cache.Storage, sporttiID string) {
	if c == nil {
		return
	}
	_ = c.DeleteByPrefixes(ctx, fmt.Sprintf("%s:%s", archDataPrefix, sporttiID))
}

func invalidateArchAll(ctx context.Context, c *cache.Storage, sporttiID string) {
	if c == nil {
		return
	}
	_ = c.DeleteByPrefixes(
		ctx,
		fmt.Sprintf("%s:%s", archDataPrefix, sporttiID),
		fmt.Sprintf("%s:%s", archSessionsPrefix, sporttiID),
		fmt.Sprintf("%s:%s:", archHTMLPrefix, sporttiID),
	)
}
