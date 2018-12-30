package errors

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

const MESSAGE_FILE = "../config/errors.yaml"

func TestNewAPIError(t *testing.T) {
	defer func() {
		templates = nil
	}()
	assert.Nil(t, LoadMessages(MESSAGE_FILE))

	e := NewAPIError(http.StatusContinue, "xyz", nil)
	assert.Equal(t, http.StatusContinue, e.Status)
	assert.Equal(t, "xyz", e.Message)

	e = NewAPIError(http.StatusNotFound, "NOT_FOUND", nil)
	assert.Equal(t, http.StatusNotFound, e.Status)
}
