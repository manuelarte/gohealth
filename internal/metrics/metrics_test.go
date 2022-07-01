package metrics

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMetrics(t *testing.T) {
	testWriter := httptest.NewRecorder()
	testRequest := httptest.NewRequest("GET", "/metrics", nil)

	HandleMetrics(testWriter, *testRequest)
	assert.Equal(t, 200, testWriter.Result().StatusCode)

	var res MemStats
	json.NewDecoder(testWriter.Body).Decode(&res)
	assert.NotNil(t, res)
}
