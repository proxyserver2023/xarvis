package api

import (
	"log"
	"net/http"
)

func LogMiddleware(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nw := NewLogResponseWriter(w)
		f(nw, r)
		log.Println(nw.String(r))
	}
}
