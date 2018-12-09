package middleware

import (
	"log"
	"net/http"
	"time"
)

// Logger Logs into StdOut like the following
// 2014/11/19 12:41:39 GET /todos  TodoIndex 148.324us
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)

	})
}
