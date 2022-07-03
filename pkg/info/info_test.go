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
	assert.Equal(t, 200, writer.Result().StatusCode)
}
