package actuator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/manuelarte/gohealth/pkg/health"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ServerSuite struct {
	suite.Suite
}

func (suite *ServerSuite) TestRootEndpoint() {
	server := httptest.NewServer(GetHandler(&Config{}))
	res, err := http.Get(server.URL + "/")
	if err != nil {
		assert.FailNow(suite.T(), "request error", err)
	}
	defer res.Body.Close()

	assert.Equal(suite.T(), 200, res.StatusCode)
	var resBody map[string][]string
	_ = json.NewDecoder(res.Body).Decode(&resBody)

	assert.Equal(suite.T(), len(allEndpoints), len(resBody["_links"]))
}

func (suite *ServerSuite) TestHealthEndpoint() {
	server := httptest.NewServer(GetHandler(&Config{}))
	res, err := http.Get(server.URL + "/health")
	if err != nil {
		assert.FailNow(suite.T(), "request error", err)
	}
	defer res.Body.Close()

	assert.Equal(suite.T(), 200, res.StatusCode)
	var response health.HealthCheckerResponse
	_ = json.NewDecoder(res.Body).Decode(&response)

	assert.NotNil(suite.T(), response)
	assert.Equal(suite.T(), health.UP, response.Status)
	assert.Equal(suite.T(), 2, len(response.Components))
}

func (suite *ServerSuite) TestEnvEndpoint() {
	server := httptest.NewServer(GetHandler(&Config{}))
	res, err := http.Get(server.URL + "/env")
	if err != nil {
		assert.FailNow(suite.T(), "request error", err)
	}
	defer res.Body.Close()

	assert.Equal(suite.T(), 200, res.StatusCode)
	var response map[string]string
	_ = json.NewDecoder(res.Body).Decode(&response)
	assert.NotNil(suite.T(), response)
	assert.Greater(suite.T(), len(response), 0)
}

func (suite *ServerSuite) TestThreaddumpEndpoint() {
	server := httptest.NewServer(GetHandler(&Config{}))
	res, err := http.Get(server.URL + "/threaddump")
	if err != nil {
		assert.FailNow(suite.T(), "request error", err)
	}
	defer res.Body.Close()

	assert.Equal(suite.T(), 200, res.StatusCode)
	resBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		assert.FailNow(suite.T(), "read error", err)
	}
	fmt.Println(string(resBytes))
}

func (suite *ServerSuite) TestMetricsEndpoint() {
	server := httptest.NewServer(GetHandler(&Config{}))
	res, err := http.Get(server.URL + "/metrics")
	if err != nil {
		assert.FailNow(suite.T(), "request error", err)
	}
	defer res.Body.Close()

	assert.Equal(suite.T(), 200, res.StatusCode)
	buff, err := ioutil.ReadAll(res.Body)
	if err != nil {
		assert.FailNow(suite.T(), "read error", err)
	}
	metrics := string(buff)
	fmt.Println(metrics)
	assert.Greater(suite.T(), len(metrics), 0)
}

func (suite *ServerSuite) TestInfoEndpoint() {
	server := httptest.NewServer(GetHandler(&Config{}))
	res, err := http.Get(server.URL + "/info")
	if err != nil {
		assert.FailNow(suite.T(), "request error", err)
	}
	defer res.Body.Close()

	assert.Equal(suite.T(), 200, res.StatusCode)
	buff, err := ioutil.ReadAll(res.Body)
	if err != nil {
		assert.FailNow(suite.T(), "read error", err)
	}
	info := string(buff)
	fmt.Println(info)
	assert.Greater(suite.T(), len(info), 0)
}

func TestServerSuite(t *testing.T) {
	suite.Run(t, new(ServerSuite))
}
