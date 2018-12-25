package main

import (
	"net/http"

	"github.com/rs/cors"

	"github.com/gorilla/mux"
	"github.com/meatballhat/negroni-logrus"
	"github.com/urfave/negroni"
)

func main() {
	router := initRouter()
	handler := initMiddleware(router)
	http.ListenAndServe(":8080", handler)
}

func initMiddleware(h http.Handler) http.Handler {
	n := negroni.New()
	initCORSMiddleware(n)
	initLoggerMiddleware(n)
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

func initCORSMiddleware(n *negroni.Negroni) {
	c := cors.New(
		cors.Options{
			AllowedOrigins: []string{"http://foo.com"},
		},
	)
	n.Use(c)
}

func initLoggerMiddleware(n *negroni.Negroni) {
	// l := negroni.NewLogger()
	// l.SetFormat("[{{.Status}} {{.Duration}}] - {{.Request.UserAgent}}")

	l := negronilogrus.NewMiddleware()
	// l := negronilogrus.NewCustomMiddleware(logrus.DebugLevel, &logrus.JSONFormatter{}, "web")
	n.Use(l)
}
