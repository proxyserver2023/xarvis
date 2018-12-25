package api

import "github.com/gorilla/mux"

func Handlers() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/api/v1/polls", MiddlewareHandler(listHandler)).Methods("GET")
	r.HandleFunc("/api/v1/polls/{poll}", MiddlewareHandler(getHandler)).Methods("GET")
	r.HandleFunc("/api/v1/polls", MiddlewareHandler(postHandler)).Methods("POST")
	r.HandleFunc("/api/v1/polls/{poll}", MiddlewareHandler(putHandler)).Methods("PUT")
	r.HandleFunc("/api/v1/polls/{poll}", MiddlewareHandler(deleteHandler)).Methods("DELETE")
	r.HandleFunc("/api/v1/polls/{poll}/answers/{answer}", MiddlewareHandler(voteHandler)).Methods("POST")
	r.NotFoundHandler = LogHandler(notFoundHandler)

	return r
}
