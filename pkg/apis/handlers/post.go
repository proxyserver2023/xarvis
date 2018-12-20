package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/alamin-mahamud/gapi/pkg/daos"
	"github.com/go-chi/chi"
)

type PostHandler struct {
	Db daos.DB
}

func NewPostHandler(databaseRepo daos.DB) *PostHandler {
	return &PostHandler{
		Db: databaseRepo,
	}
}

func (p *PostHandler) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := p.Db.Fetch(r.Context(), 5)
	respondwithJSON(w, http.StatusOK, payload)
}

// GetByID returns a post details
func (p *PostHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := p.Db.GetByID(r.Context(), int64(id))

	if err != nil {
		responseWithError(w, http.StatusNoContent, "Content not found")
	}

	respondwithJSON(w, http.StatusOK, payload)
}

// respondwithJSON write json response format
func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func responseWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
