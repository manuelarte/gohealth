package main

import (
	"github.com/manuelarte/gohealth/pkg/info"
	"net/http/httptest"
)

type custominfo struct {
	Author string `json:"author,omitempty"`
}

func main() {
	info.CommitAuthor = "manuelarte"

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/info", nil)
	var payloadFunc info.Payload[custominfo] = func(info info.Info) custominfo {
		return custominfo{
			Author: info.Git.CommitAuthor,
		}
	}

	info.HandleCustomInfo[custominfo](writer, request, payloadFunc)
	println(writer.Body.String())
}
