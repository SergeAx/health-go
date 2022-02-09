package uptime

import (
	"testing"
)

func TestUptimeFunc(t *testing.T) {
	u, err := upTime()
	if err != nil {
		t.Error(err)
	}

	t.Log("Got uptime() = ", u)
}
