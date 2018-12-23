package app

import (
	"net/http"

	"github.com/alamin-mahamud/gapi/pkg/article"
)

func Start() {
	r := article.InitRouter()
	http.ListenAndServe(":8089", r)
}
