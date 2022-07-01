package internal

import (
	"net/http"
	"os"
)

// HandleShutdown - When a call is made to /shutdown, this response handler
// shuts down the application with the exit code 0
func HandleShudown(writer http.ResponseWriter, request *http.Request) {
	os.Exit(0)
}