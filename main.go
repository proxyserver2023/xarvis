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

func Middleware1(next http.Handler) http.Handler {
	// do some stuff
	// call the next middleware
	// do some stuff
	return next
}

func Middleware2(next http.Handler) http.Handler {
	// do some stuff
	// call the next middleware
	// do some stuff
	return next
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	apiRoutes := mux.NewRouter()
	webRoutes := mux.NewRouter()
	common := negroni.New(
		Middleware1,
		Middleware2,
	)
	router.PathPrefix("/").Handler(common.With(
		APIMiddleware1,
		negroni.Wrap(apiRoutes),
	))
	http.ListenAndServe(":8080", router)
}
