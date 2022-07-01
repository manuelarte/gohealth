package health

import (
	"syscall"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DiskUsageCheckerSuite struct {
	suite.Suite
}

func (suite *DiskUsageCheckerSuite) TestDiskStatus() {
	checker := DiskStatusChecker{}
	result := checker.CheckHealth()

	assert.Equal(suite.T(), UP, result.Status)
}

func (suite *DiskUsageCheckerSuite) TestDiskStatusError() {
	realFsStat := fsStats
	defer func() { fsStats = realFsStat }()

	fsStats = func(path string, buf *syscall.Statfs_t) (err error) { return assert.AnError }

	checker := DiskStatusChecker{}
	result := checker.CheckHealth()
	assert.Equal(suite.T(), DOWN, result.Status)
}

func TestDiskUsageSuite(t *testing.T) {
	suite.Run(t, new(DiskUsageCheckerSuite))
}
