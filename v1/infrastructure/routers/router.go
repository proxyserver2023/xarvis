package routers

import (
	"github.com/alamin-mahamud/go-bootstrap/v1/infrastructure/middleware"
	"github.com/gorilla/mux"
)

// Init - returns a new instance of router after
// registering all the routes.
func Init() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		var handler = middleware.Logger(route.HandlerFunc, route.Name)
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
