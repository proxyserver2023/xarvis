package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func dummy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "weadskfjksadljfklajs")
}

func dummy2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "13412341234134")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	adminRoutes := router.PathPrefix("/admin").Subrouter()
	adminRoutes.HandleFunc("/", dummy)
	adminRoutes.HandleFunc("/{id}", dummy2)

	router.PathPrefix("/admin").Handler(
		negroni.New(
			// Middleware - 1,
			// Middleware - 2,
			negroni.Wrap(adminRoutes),
		),
	)

	http.ListenAndServe(":8080", router)
}
