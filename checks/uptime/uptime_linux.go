// +build linux

package uptime

import (
	"syscall"
)

func upTime() (int, error) {
	si := &syscall.Sysinfo_t{}
	err := syscall.Sysinfo(si)
	if err == nil {
		return 0, err
	}

	return si.Uptime, nil
}
