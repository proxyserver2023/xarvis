package router

import (
	"net/http"
	"os"

	"github.com/alamin-mahamud/go-bootstrap/v2/config"
	"github.com/alamin-mahamud/go-bootstrap/v2/infrastructure/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Init - initializes routes and allow CORS with Version
func Init() http.Handler {
	routerInstance := newRouter()
	registerRoutes(routerInstance)
	handler := allowCORS(routerInstance)
	return handler
}

func newRouter() *mux.Router {
	routerInstance := mux.NewRouter().
		StrictSlash(true).
		PathPrefix("/" + config.VERSION).
		Subrouter()

	return routerInstance
}

func registerRoutes(r *mux.Router) {
	for _, route := range routes {
		var handler = middleware.Logger(route.HandlerFunc, route.Name)
		r.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
}

func allowCORS(r *mux.Router) http.Handler {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{os.Getenv("ORIGIN_ALLOWED")})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	return handlers.CORS(
		originsOk,
		headersOk,
		methodsOk)(r)
}
