# goctuator

Monitoring is an important function for monitoring and observing
Application Program Interfaces (APIs). This project creates a series
of endpoints which may be used to instrument and monitor a Go-based
API.

This project is inspired by
[spring-boot-actuator](https://docs.spring.io/spring-boot/docs/2.7.1/reference/html/actuator.html)
and [go-actuator](https://github.com/sinhashubham95/go-actuator).

## Quickstart

To use this library, install the go module:

```bash
go get gitlab.com/mikeyGlitz/gohealth
```

After the module has been installed, import the handler into your framework.
The actuator endpoint implements the standard Go handler function as prescribed
by the `http.HandlerFunc` in the `net/http` library.

### HTTP Package

```go
handler := actuator.GetHandler(&actuator.Config{})
http.Handle("/actuator", handler)
```

### Gin Handler

```go
handler := actuator.GetHandler(&actuator.Config{})
router := gin.Default()
r.Get("/actuator", func(ctx *gin.Context) {
    handler(ctx.Writer, ctx.Request)
})
```

### Additional Frameworks

For additional frameworks, please fill out a documentation request
in the repo's issues section.
I'm only familiar with `net/http` and [gin-gonic](https://gin-gonic.com/)
at the moment.

## Configuration

This package exposes a configuration in the form of the `actuator.Configuration`
structure. The configuration enables developers to change the endpoints that
goctuator exposes. A developer may also supply custom health checkers to the API
to enable the `/health` endpoint to report the status on more services (i.e. a database connection)

```go
config := &actuator.Configuration{
    Endpoints: []actuator.Endpoint{
        actuator.ENV,
        actuator.INFO,
        actuator.THREADDUMP,
        actuator.HEALTH,
        actuator.METRICS,
        actuator.SHUTDOWN,
    },
    HealthCheckers: []health.HealthChecker{
        health.PingChecker{},
        health.DiskStatusChecker{},
        // Add additional custom health checkers
    }
}
```

If no `Endpoints` or `HealthCheckers` are provided, the actuator will be configured
to use all of the default built-in endpoints and health status checkers.

## Endpoints

### `/Env`

This endpoint exposes the runtime environment variables as a series of key-value pairs

```json
{
  "key_1": "value1",
  "key_2": "value2",
  "key_3": "value3",
  "key_4": "value4"
}
```

### `/shutdown`

When a call is made to `/actuator/shutdown` the current running process
(in this case, your API) will terminate with an exit status of 0

### `/metrics`

Metrics displays all a JSON-encoded snapshot of the memory statistics at the time
that the request was received. This endpoint relies on the `runtime.ReadMemStats()`
function to gather the requested information

```JSON
{
    "alloc": 969240,
    "total_alloc": 969240,
    "sys": 10700041,
    "lookups": 0,
    "mallocs": 6645,
    "frees": 218,
    "heap_alloc": 969240,
    "heap_sys": 3702784,
    "heap_idle": 1490944,
    "heap_inuse": 2211840,
    "heap_released": 1392640,
    "heap_objects": 6427,
    "stack_inuse": 491520,
    "stack_sys": 491520,
    "m_span_inuse": 46512,
    "m_span_sys": 48960,
    "m_cache_in_use": 19200,
    "m_cache_sys": 31200,
    "buck_hash_sys": 1445238,
    "gc_sys": 3992984,
    "other_sys": 987354,
    "next_gc": 4194304,
    "last_gc": 0,
    "pause_total_ns": 0,
    "pause_ns": [0],
    "pause_end": [0],
    "num_gc": 0,
    "num_forced_gc": 0,
    "gc_cpu_fraction": 0,
    "enable_gc": true,
    "debug_gc": false,
    "BySize": [
        {
            "size": 0,
            "mallocs": 0,
            "frees": 0
        }
    ]
}
```

### `/threaddump`

Thread dump displays a snapshot of the current thread at the time the request
was received. This endpoint leverages `pprof.Lookup` to gather information from the
`goroutine` profile.

```text
oroutine profile: total 1
1 @ 0x10ad368 0x10ad170 0x10a9bd4 0x10b3ee5 0x102a407 0x1052dc1
#       0x10ad367       runtime/pprof.writeRuntimeProfile+0x97  /usr/local/go/src/runtime/pprof/pprof.go:707
#       0x10ad16f       runtime/pprof.writeGoroutine+0x9f       /usr/local/go/src/runtime/pprof/pprof.go:669
#       0x10a9bd3       runtime/pprof.(*Profile).WriteTo+0x3e3  /usr/local/go/src/runtime/pprof/pprof.go:328
#       0x10b3ee4       main.main+0x64                          /Users/inanc/go/src/github.com/inancgumus/main.go:9
#       0x102a406       runtime.main+0x206                      /usr/local/go/src/runtime/proc.go:201
```

### `/info`

This endpoint exposes compile-time information about the API binary

```json
{
  "app": {
    "name": "my-app",
    "description": "A test API application which uses actuators",
    "version": "v1.0.0"
  },
  "git": {
    "commit_author": "John Doe <jdoe@example.com>",
    "commit_id": "",
    "commit_time": "",
    "build_time": "",
    "branch": "main",
    "url": "https://gitlab.com/jdoe/example-api"
  },
  "runtime": {
    "arch": "x86_64",
    "os": "darwin",
    "version": "go1.17"
  }
}
```

#### Updating the `/info` Endpoint

The variables for the `/info` endpoint are expected to be set at compile-time
when you build your application.
Variables are set with [ldflags](https://www.digitalocean.com/community/tutorials/using-ldflags-to-set-version-information-for-go-applications).

```bash
go build -ldflags="-X gitlab.com/mikeyGlitz/gohealth/pkg/info.AppName=${appName} <other flags>"
```

The `/info` endpoint utilizes the following variables:

| Variable                    | Description                                                |
| --------------------------- | ---------------------------------------------------------- |
| AppName                     | The name of your application                               |
| AppDescription              | A brief description of what your application does          |
| AppVersion                  | An application version i.e. v1.0.0                         |
| **Git Variables**           |                                                            |
| CommitID                    | The SHA1 of the commit                                     |
| CommitTime                  | A timestamp of when the commit took place                  |
| BuildTime                   | A timestamp of when the application build took place       |
| RepositoryUrl               | A URL of where the Git repository is located               |
| Branch                      | The branch name the build was executed on                  |
| **Application Environment** |                                                            |
| OS                          | Operating system the application runs - darwin, linux      |
| RuntimeVersion              | The version of go the application was built for            |
| Arch                        | The CPU architecture the application was built for - arm64 |

### `/health`

The health endpoint displays the running health of the application and any supporting services
(i.e. database connections)

This endpoint was inspired by the spring-boot-actuator project. Responses to `/health` are in the
following shape:

```json
{
  "status": "UP",
  "components": {
    "diskspace": {
      "status": "UP",
      "details": {
        "all": "",
        "used": "",
        "free": "",
        "available": ""
      }
    },
    "ping": {
      "status": "UP"
    }
  }
}
```

If any services are `DOWN` the top-level status will display `DOWN` and the endpoint
will respond with a status code `503`.

#### Writing Your Own Health Checks

The `pkg/health` package is designed in a way that health checks can be extended.
To create your own health check implement the `HealthChecker` interface.
The following example below demonstrates how to implement an OpenSearch health checker:

```go
type OpensearchHealthChecker struct {
    Client *opensearch.Client
}

func (checker *OpensearchHealthChecker) CheckHealth() health.HealthCheckerResponse {
    res, err := checker.Client.Cluster.Health()
    // prepare the healthcheck response
}
```

The health checker would then be added in when you want to generate the actuator handler

```go
config := &actuator.Config{
    HealthCheckers: []health.HealthChecker{
        &health.PingChecker{},
        &health.DiskChecker{}
        &OpensearchHealthChecker{},
    }
    handler := actuator.GetHandler(&config)
}
```

## References

- Spring Boot Actuator Reference - https://www.baeldung.com/spring-boot-actuators
- go-actuator - https://github.com/sinhashubham95/go-actuator
- threadump - https://forum.golangbridge.org/t/how-to-take-thread-dump-in-golang/11417/4
- MemStats - https://pkg.go.dev/runtime#MemStats
- Read Mem Stats - https://pkg.go.dev/runtime#ReadMemStats
- Disk Usage - https://stackoverflow.com/questions/20108520/get-amount-of-free-disk-space-using-go
- Using ldflags - https://www.stackovercloud.com/2019/10/25/using-ldflags-to-set-version-information-for-go-applications/#:~:text=As%20mentioned%20before%2C%20ldflags%20stands%20for%20linker%20flags%2C,that%20runs%20as%20a%20part%20of%20go%20build.
