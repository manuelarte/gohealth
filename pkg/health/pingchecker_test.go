package health

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingHandler(t *testing.T) {
	checker := PingChecker{}
	result := checker.CheckHealth()
	assert.Equal(t, UP, result.Status)
	assert.Equal(t, "ping", result.Service)
}
