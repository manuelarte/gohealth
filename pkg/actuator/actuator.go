package actuator

import (
	"encoding/json"
	"net/http"
	"strings"

	"gitlab.com/mikeyGlitz/gohealth/internal"
	"gitlab.com/mikeyGlitz/gohealth/internal/env"
	"gitlab.com/mikeyGlitz/gohealth/internal/metrics"
	"gitlab.com/mikeyGlitz/gohealth/internal/threaddump"
	"gitlab.com/mikeyGlitz/gohealth/pkg/health"
	"gitlab.com/mikeyGlitz/gohealth/pkg/info"
)

type Endpoint string

const (
	INFO       Endpoint = "info"
	ENV        Endpoint = "env"
	THREADDUMP Endpoint = "threaddump"
	SHUTDOWN   Endpoint = "shutdown"
	METRICS    Endpoint = "metrics"
	HEALTH     Endpoint = "health"
)

// All endpoints supported by the actuator
var allEndpoints = []Endpoint{
	INFO,
	ENV,
	THREADDUMP,
	METRICS,
	HEALTH,
	SHUTDOWN,
}

type Config struct {
	Endpoints      []Endpoint
	HealthCheckers []health.HealthChecker
}

// setConfigDefaults - Set the default values if fields in the config are empty
func setConfigDefaults(config *Config) {
	// If the endpoints are null or empty, set them to the defaults
	if config.Endpoints == nil {
		config.Endpoints = allEndpoints
	}

	if config.HealthCheckers == nil {
		config.HealthCheckers = health.DefaultHealthCheckers
	}
}

// GetHandler - Prepares the handler function to be attached to an endpoint
func GetHandler(config *Config) http.HandlerFunc {
	setConfigDefaults(config)

	// Add all the health checkers set by the configuration
	for _, checker := range config.HealthCheckers {
		health.Add(checker)
	}

	// Prepare the list of endpoints based on what is configured
	handlerMap := make(map[Endpoint]http.HandlerFunc)
	for _, ep := range config.Endpoints {
		switch ep {
		case INFO:
			handlerMap[INFO] = info.HandleInfo
		case ENV:
			handlerMap[ENV] = env.HandleEnv
		case METRICS:
			handlerMap[METRICS] = metrics.HandleMetrics
		case HEALTH:
			handlerMap[HEALTH] = health.HandleHealthchecks
		case THREADDUMP:
			handlerMap[THREADDUMP] = threaddump.HandleThreaddump
		case SHUTDOWN:
			handlerMap[SHUTDOWN] = internal.HandleShudown
		}
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		// Validate the request. We only accept GET protocol here
		if request.Method != "GET" {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			_, _ = writer.Write([]byte("Method Not Allowed"))
			return
		}

		// Extract the path from the URL
		paths := strings.Split(request.URL.Path, "/")
		entrypoint := ""
		pathLength := len(paths)
		if pathLength > 0 {
			entrypoint = paths[pathLength-1]
		}

		// Assumption: if we have an empty string, assume we're requesting the root url
		if len(entrypoint) == 0 {
			links := make([]string, 0)
			for k := range handlerMap {
				links = append(links, "/"+string(k))
			}

			payload := map[string][]string{"_links": links}
			_ = json.NewEncoder(writer).Encode(payload)
			return
		}

		// Match the handler with the appropriate endpoint
		if handler, ok := handlerMap[Endpoint(entrypoint)]; ok {
			handler(writer, request)
			return
		}

		writer.WriteHeader(http.StatusNotFound)
		_, _ = writer.Write([]byte("Not found"))
	}
}
