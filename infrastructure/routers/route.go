package routers

import "net/http"

// Route - Defines the URL Structure for a typical REST API
// "NAME" => METHOD + PATTERN => HandlerFunc
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
