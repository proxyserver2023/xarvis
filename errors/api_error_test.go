package errors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAPIError(t *testing.T) {
	t.Run("Error() should return the APIError Message Field", func(t *testing.T) {
		e := APIError{
			Message: "abc",
		}
		assert.Equal(t, "abc", e.Error())
	})

	t.Run("StatusCode() returns status", func(t *testing.T) {
		e := APIError{
			Status: 400,
		}
		assert.Equal(t, 400, e.StatusCode())
	})
}
