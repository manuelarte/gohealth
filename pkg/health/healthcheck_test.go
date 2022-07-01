package health

import (
	"encoding/json"
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type testChecker struct {
	Status HealthStatus
}

func (c *testChecker) CheckHealth() (result HealthCheckResult) {
	result.Service = "test"
	result.Status = c.Status
	return
}

func TestHandleHealthcheckError(t *testing.T) {
	err := errors.New("test")
	result := HandleHealthcheckError("test", err)
	assert.Equal(t, DOWN, result.Status)
}

func TestAddChecker(t *testing.T) {
	realCheckers := checkers
	defer func() { checkers = realCheckers }()
	checkers = make([]HealthChecker, 0)

	checker := testChecker{}
	Add(&checker)
	assert.Equal(t, 1, len(checkers))
}

type HealthcheckerSuite struct {
	suite.Suite
	checkers []HealthChecker
}

func (suite *HealthcheckerSuite) SetupTest() {
	suite.checkers = checkers
}

func (suite *HealthcheckerSuite) TearDownTest() {
	checkers = suite.checkers
}

func (suite *HealthcheckerSuite) TestHealthCheckerSingleChecker() {
	checkers = make([]HealthChecker, 0)
	checkers = append(checkers, &testChecker{UP})

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/health", nil)

	HandleHealthchecks(writer, request)

	assert.Equal(suite.T(), 200, writer.Result().StatusCode)

	var response HealthCheckerResponse
	json.NewDecoder(writer.Body).Decode(&response)

	assert.Equal(suite.T(), UP, response.Status)
}

func (suite *HealthcheckerSuite) TestHealthCheckerBadChecker() {
	checkers = make([]HealthChecker, 0)
	checkers = append(checkers, &testChecker{DOWN})

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/health", nil)

	HandleHealthchecks(writer, request)

	assert.Equal(suite.T(), 500, writer.Result().StatusCode)

	var response HealthCheckerResponse
	json.NewDecoder(writer.Body).Decode(&response)

	assert.Equal(suite.T(), DOWN, response.Status)
}

func TestHealthcheckerSuite(t *testing.T) {
	suite.Run(t, new(HealthcheckerSuite))
}
