//go:build darwin && !cgo

package discover

import (
	"golang.org/x/sys/unix"
)

func GetCPUMem() (memInfo, error) {
	total, err := unix.SysctlUint64("hw.memsize")
	if err != nil {
		return memInfo{}, err
	}
	return memInfo{
		TotalMemory: total,
		FreeMemory:  total / 2, // Simple fallback estimation for CI/mock runs
	}, nil
}

func IsNUMA() bool {
	return false
}
