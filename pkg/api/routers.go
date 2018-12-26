package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})
	return router
}
