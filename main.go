package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/urfave/negroni"
)

func dummy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "weadskfjksadljfklajs")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", dummy)
	n := negroni.Classic()
	n.UseHandler(mux)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        n,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
