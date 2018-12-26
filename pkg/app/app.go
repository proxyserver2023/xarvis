package app

import (
	"net/http"

	"github.com/alamin-mahamud/gapi/pkg/api"
)

func Start() {
	loadApplicationConfiguration()
	loadErrorMessages()
	connectDatabase()
	apiRouting()
	startTheServer()
	router := api.InitRouter()
	handler := api.InitMiddleware(router)
	http.Handle("/", handler)
	http.ListenAndServe(":8080", nil)
}
