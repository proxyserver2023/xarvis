package routers

import (
	articleHandler "github.com/alamin-mahamud/go-bootstrap/infrastructure/handlers/articles"
	userHandler "github.com/alamin-mahamud/go-bootstrap/infrastructure/handlers/users"
)

// Routes are grouped by according to their usage.
type Routes []Route

var articleRoutes = Routes{
	Route{"List", "GET", "/articles/list", articleHandler.List},

	/*
		Route{},
		Route{},
		Route{},
	*/
}

var userRoutes = Routes{
	Route{"List", "GET", "/users/list", userHandler.List},
}

var routes = append(articleRoutes, userRoutes...)
