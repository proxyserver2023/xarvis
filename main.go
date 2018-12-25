package main

import (
	"net/http"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func main() {
	router := initRouter()
	handler := initMiddleware(router)
	http.ListenAndServe(":8080", handler)
}

func initMiddleware(h http.Handler) http.Handler {
	n := negroni.New()
	c := cors.New(
		cors.Options{
			AllowedOrigins: []string{"http://foo.com"},
		},
	)
	n.Use(c)
	n.UseHandler(h)
	return n
}

func initRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})
	return router
}
