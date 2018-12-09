package router

import (
	articleController "github.com/alamin-mahamud/go-bootstrap/v2/infrastructure/controller/article"
)

// Routes are grouped by according to their usage.
type Routes []Route

var articleRoutes = Routes{
	Route{"List", "GET", "/articles/list", articleController.List},
}

var userRoutes = Routes{
	// Route{"List", "GET", "/users/list", userController.List},
}

var routes = append(articleRoutes, userRoutes...)
