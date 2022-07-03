package threaddump

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleThreaddump(t *testing.T) {
	writer := httptest.NewRecorder()
	reader := httptest.NewRequest("GET", "/threaddump", nil)

	HandleThreaddump(writer, reader)
	assert.Equal(t, 200, writer.Result().StatusCode)
}
