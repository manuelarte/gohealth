package health

import "syscall"

const diskServiceName = "diskspace"

type DiskStatus struct {
	All       uint64 `json:"all"`
	Used      uint64 `json:"used"`
	Free      uint64 `json:"free"`
	Available uint64 `json:"available"`
}

type DiskStatusChecker struct{}

var fsStats = syscall.Statfs

func (a *DiskStatusChecker) CheckHealth() (result HealthCheckResult) {
	fs := syscall.Statfs_t{}
	if err := fsStats("/", &fs); err != nil {
		return HandleHealthcheckError(diskServiceName, err)
	}
	status := DiskStatus{
		All:       fs.Blocks * uint64(fs.Bsize),
		Free:      fs.Bfree * uint64(fs.Bsize),
		Available: fs.Bavail * uint64(fs.Bsize),
	}
	status.Used = status.All - status.Free

	result.Service = diskServiceName
	result.Details = status
	result.Status = UP
	return
}
