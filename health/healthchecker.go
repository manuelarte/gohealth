package health

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HealthStatus int

const (
	DOWN HealthStatus = 0 // DOWN	- The service is not healthy
	UP   HealthStatus = 1 // UP		- The service is healthy
)

type ComponentDetails struct {
	Status  HealthStatus `json:"status"`  // HealthStatus - An indicator for whether the service is healthy or not
	Details interface{}  `json:"details"` // Details - a series of values which will be serialized into JSON in the final request
}

type HealthCheckResult struct {
	ComponentDetails
	Service string // Service - The name of the service being checked
}

/*
	HealthChecker - An interface for a listener which is utilized to check the running status of a service
*/
type HealthChecker interface {
	// CheckHealth - checks the service for its current running status
	CheckHealth() HealthCheckResult
}

var checkers = make([]HealthChecker, 0)

// Add - Adds a HealthChecker to the global slice of checkers
func Add(checker HealthChecker) {
	checkers = append(checkers, checker)
}

type HealthCheckerResponse struct {
	Status     HealthStatus                `json:"status"`
	Components map[string]ComponentDetails `json:"components"`
}

type HealthCheckError struct {
	Message string
}

// HandleHealthcheckError utility function for handling errors since we're not intending to use a logger for this module
func HandleHealthcheckError(service string, err error) (result HealthCheckResult) {
	result.Service = service
	result.Details = &HealthCheckError{Message: fmt.Sprintf("error: %v", err)}
	result.Status = DOWN
	return
}

/*
	HandleHealthchecks - Composes all of the individual healthchecks into a single response object and sends it out
*/
func HandleHealthchecks(writer http.ResponseWriter, request *http.Request) {
	// Collect the healthcheck results
	topLevelStatus := UP
	payload := HealthCheckerResponse{Components: make(map[string]ComponentDetails)}
	for _, checker := range checkers {
		result := checker.CheckHealth()

		// Gather fields for the payload
		if result.Status == DOWN {
			topLevelStatus = DOWN
		}
		payload.Components[result.Service] = ComponentDetails{
			Status:  result.Status,
			Details: result.Details,
		}
	}

	payload.Status = topLevelStatus
	status := http.StatusOK
	if topLevelStatus == DOWN {
		status = http.StatusInternalServerError
	}

	writer.WriteHeader(status)

	// trunk-ignore(golangci-lint/errcheck)
	json.NewEncoder(writer).Encode(payload)
}
