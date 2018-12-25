package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func dummy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "weadskfjksadljfklajs")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	adminRoutes := router.PathPrefix("/admin").Subrouter()
	adminRoutes.HandleFunc("/", dummy)
	http.ListenAndServe(":8080", router)
}
