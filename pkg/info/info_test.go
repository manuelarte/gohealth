package info

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInfoHandler(t *testing.T) {
	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/info", nil)

	HandleInfo(writer, request)
	expected := `{
		"app": {},
		"git": {},
		"runtime": {}
	}`
	assert.Equal(t, 200, writer.Result().StatusCode)
	assert.JSONEq(t, expected, writer.Body.String())
}

func TestInfoHandler_CustomPayload(t *testing.T) {
	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/info", nil)

	type custominfo struct {
		Author string `json:"author,omitempty"`
	}
	var payloadFunc Payload[custominfo] = func(info Info) custominfo {
		return custominfo{
			Author: info.Git.CommitAuthor,
		}
	}

	HandleCustomInfo[custominfo](writer, request, payloadFunc)
	expected := "{}"

	assert.Equal(t, 200, writer.Result().StatusCode)
	assert.JSONEq(t, expected, writer.Body.String())
}
