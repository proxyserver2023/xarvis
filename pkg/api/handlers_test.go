package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	server   *httptest.Server
	pollsURL string
)

func init() {
	server = httptest.NewServer(Handlers())
	pollsURL = fmt.Sprintf("%s/api/v1/polls", server.URL)
}

func TestNotFound(t *testing.T) {
	r, _ := http.NewRequest(
		"GET",
		server.URL+"notFound",
		nil,
	)
	w := httptest.NewRecorder()
	notFoundHandler(w, r)
	assert := assert.New(t)
	assert.Equal(
		http.StatusNotFound,
		w.Code,
		"Status must be 404",
	)
}
