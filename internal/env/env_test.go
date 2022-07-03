package env

import (
	"encoding/json"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleEnv(t *testing.T) {
	testWriter := httptest.NewRecorder()
	testRequest := httptest.NewRequest("GET", "/env", nil)

	os.Setenv("testname", "envtest")
	defer func() { os.Unsetenv("testname") }()

	HandleEnv(testWriter, testRequest)
	assert.Equal(t, 200, testWriter.Result().StatusCode)

	var result map[string]string
	json.NewDecoder(testWriter.Body).Decode(&result)
	assert.Equal(t, "envtest", result["testname"])
}
