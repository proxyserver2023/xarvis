package routers

import (
	"net/http"

	ph "github.com/alamin-mahamud/gapi/pkg/apis/handlers"
	"github.com/alamin-mahamud/gapi/pkg/daos"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func BuildRouter(db daos.DB) http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	pHandler := ph.NewPostHandler(db)
	r.Route("/", func(rt chi.Router) {
		rt.Mount("/posts", postRouter(pHandler))
	})

	return r
}

// A completely separate router for posts routes
func postRouter(pHandler *ph.PostHandler) http.Handler {
	r := chi.NewRouter()
	r.Get("/", pHandler.Fetch)
	r.Get("/{id:[0-9]+}", pHandler.GetByID)

	// r.Post("/", pHandler.Create)
	// r.Put("/{id:[0-9]+}", pHandler.Update)
	// r.Delete("/{id:[0-9]+}", pHandler.Delete)

	return r
}
