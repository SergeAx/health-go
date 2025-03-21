// Package uptime provides uptime-related health Checks.
package uptime

import (
	"net/http"
	"time"

	"github.com/SergeAx/health-go"
)

type system struct {
}

func (u *system) HealthChecks() map[string][]health.Checks {
	ut, err := upTime()
	now := time.Now().UTC().Format(time.RFC3339Nano)
	var uptime func() health.Checks
	if err != nil {
		uptime = func() health.Checks {
			return health.Checks{
				ComponentType: "system",
				Status:        health.Fail,
				Output:        err.Error(),
				Time:          now,
			}
		}
	} else {
		uptime = func() health.Checks {
			return health.Checks{
				ComponentType: "system",
				ObservedValue: ut,
				ObservedUnit:  "s",
				Status:        health.Pass,
				Time:          now,
			}
		}
	}
	return map[string][]health.Checks{
		"uptime": {
			uptime(),
		},
	}
}

func (*system) AuthorizeHealth(*http.Request) bool {
	return true
}

// System returns a ChecksProvider for health checks about the system uptime.
func System() health.ChecksProvider {
	return &system{}
}
