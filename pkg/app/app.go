package app

import (
	"net/http"

	"github.com/alamin-mahamud/gapi/pkg/api"
)

func Start() {
	router := api.InitRouter()
	handler := api.InitMiddleware(router)
	http.ListenAndServe(":8080", handler)
}
