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

func Middleware1(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("1")
	next(w, r)
	fmt.Println("6")
}

func Middleware2(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("2")
	next(w, r)
	fmt.Println("5")
}

func Middleware3(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("3")
	next(w, r)
	fmt.Println("4")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	apiRoutes := mux.NewRouter()
	common := negroni.New(
		negroni.HandlerFunc(Middleware1),
		negroni.HandlerFunc(Middleware2),
	)
	router.PathPrefix("/").Handler(common.With(
		negroni.HandlerFunc(Middleware3),
		negroni.Wrap(apiRoutes),
	))
	http.ListenAndServe(":8080", router)
}
