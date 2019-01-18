package router

import (
	"net/http"

	"github.com/alamin-mahamud/xarvis/pkg/middleware"
	"github.com/gorilla/mux"
)

// Init - initializes routes
func Init() http.Handler {
	routerInstance := newRouter()
	initAuthRouter(routerInstance)
	finalHandler := middleware.Init(routerInstance)
	return finalHandler
}

func newRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	versionedRouter := r.PathPrefix("/" + "v1").Subrouter()
	return versionedRouter
}

// install routes with appropriate handlers and database connection
func initAuthRouter(r *mux.Router) {
	authRouter := r.PathPrefix("/authentication").Subrouter()
	registerAuthRoutes(authRouter)
}

func registerAuthRoutes(r *mux.Router) {
	routes := getAuthRoutes()
	for _, route := range routes {
		r.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
}
