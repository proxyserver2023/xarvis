package article

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/list", List)
	return r
}
