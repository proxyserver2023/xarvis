package errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestInternalServerError(t *testing.T) {
	assert.Equal(
		t,
		http.StatusInternalServerError,
		InternalServerError(errors.New("")).Status,
	)
}

func TestUnauthorized(t *testing.T) {
	assert.Equal(
		t,
		http.StatusUnauthorized,
		Unauthorized("t").Status,
	)

}
