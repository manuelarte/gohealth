package env

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

func HandleEnv(writer http.ResponseWriter, request *http.Request) {
	envMap := make(map[string]string)
	for _, env := range os.Environ() {
		pair := strings.Split(env, "=")
		envMap[pair[0]] = pair[1]
	}

	_ = json.NewEncoder(writer).Encode(&envMap)
}
