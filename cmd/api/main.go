package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello there")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":3000", n)
}
