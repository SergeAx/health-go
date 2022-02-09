//go:build windows

package uptime

import (
	"syscall"
	"time"

	"golang.org/x/sys/windows"
)

// A little copying from https://github.com/shirou/gopsutil

var (
	ModKernel32        = windows.NewLazySystemDLL("kernel32.dll")
	procGetTickCount32 = ModKernel32.NewProc("GetTickCount")
	procGetTickCount64 = ModKernel32.NewProc("GetTickCount64")
)

func upTime() (int64, error) {
	procGetTickCount := procGetTickCount64
	err := procGetTickCount64.Find()
	if err != nil {
		procGetTickCount = procGetTickCount32 // handle WinXP, but keep in mind that "the time will wrap around to zero if the system is run continuously for 49.7 days." from MSDN
	}
	r1, _, lastErr := syscall.Syscall(procGetTickCount.Addr(), 0, 0, 0, 0)
	if lastErr != 0 {
		return 0, lastErr
	}
	return int64((time.Duration(r1) * time.Millisecond).Seconds()), nil
}
