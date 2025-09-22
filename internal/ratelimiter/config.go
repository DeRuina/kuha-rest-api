package ratelimiter

import "time"

var RoleLimits = map[string]struct {
	Limit  int
	Window time.Duration
}{
	"admin":      {Limit: 5000, Window: time.Minute},
	"fis":        {Limit: 1000, Window: time.Minute},
	"utv":        {Limit: 3000, Window: time.Minute},
	"kamk":       {Limit: 1000, Window: time.Minute},
	"klab":       {Limit: 1000, Window: time.Minute},
	"tietoevry":  {Limit: 5000, Window: time.Minute},
	"coachtech":  {Limit: 1000, Window: time.Minute},
	"archinisis": {Limit: 1000, Window: time.Minute},
	"default":    {Limit: 500, Window: time.Minute},
}

func GetLimitForRole(role string) (int, time.Duration) {
	if val, ok := RoleLimits[role]; ok {
		return val.Limit, val.Window
	}
	return RoleLimits["default"].Limit, RoleLimits["default"].Window
}
