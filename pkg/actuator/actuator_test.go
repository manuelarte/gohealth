package actuator

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/manuelarte/gohealth/pkg/health"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ActuatorSuite struct {
	suite.Suite
}

func (suite *ActuatorSuite) TestActuatorEmptyConfig() {
	config := &Config{}
	handler := GetHandler(config)
	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/health", nil)

	handler(writer, request)
	assert.Equal(suite.T(), health.DefaultHealthCheckers, config.HealthCheckers)
	assert.Equal(suite.T(), allEndpoints, config.Endpoints)
	assert.Equal(suite.T(), 200, writer.Result().StatusCode)
}

func (suite *ActuatorSuite) TestRootPath() {
	config := &Config{}
	handler := GetHandler(config)
	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/", nil)

	handler(writer, request)
	assert.Equal(suite.T(), 200, writer.Result().StatusCode)

	var result map[string][]string
	json.NewDecoder(writer.Body).Decode(&result)
	links := result["_links"]
	assert.Equal(suite.T(), len(allEndpoints), len(links))
}

func (suite *ActuatorSuite) TestActuatorNotFound() {
	config := &Config{}
	handler := GetHandler(config)
	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/test", nil)

	handler(writer, request)
	assert.Equal(suite.T(), 404, writer.Result().StatusCode)
}

func (suite *ActuatorSuite) TestActuatorUnsupportedMethod() {
	config := &Config{}
	handler := GetHandler(config)
	writer := httptest.NewRecorder()
	request := httptest.NewRequest("POST", "/", nil)

	handler(writer, request)
	assert.Equal(suite.T(), 405, writer.Result().StatusCode)
}

func TestActuatorSuite(t *testing.T) {
	suite.Run(t, new(ActuatorSuite))
}
