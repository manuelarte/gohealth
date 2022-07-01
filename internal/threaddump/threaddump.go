package threaddump

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/pprof"
)

var lookup = pprof.Lookup

// HandleThreaddump - Prints out a plaintext snapshot of the process's
// current-running thread stats
// https://forum.golangbridge.org/t/how-to-take-thread-dump-in-golang/11417/4
func HandleThreaddump(writer http.ResponseWriter, request *http.Request) {
	var buff bytes.Buffer
	profile := lookup("goroutine")
	if profile == nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		errorMsg := `{"error": { "message": "unable to read profile" } }`
		_, _ = writer.Write([]byte(errorMsg))
	}
	if err := profile.WriteTo(&buff, 1); err != nil {
		writer.Header().Add("Content-Type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		errorMsg := fmt.Sprintf(`{"error": { "message": "unable to read profile: %v" } }`, err)
		_, _ = writer.Write([]byte(errorMsg))
	}

	payload := buff.Bytes()
	writer.Header().Add("Content-Type", "text/plain")
	_, _ = writer.Write(payload)
}
