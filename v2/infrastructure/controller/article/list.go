package article

import (
	"encoding/json"
	"net/http"
)

// List all articles
func List(w http.ResponseWriter, r *http.Request) {
	// repo := &repository.MongoDB{}
	// articleList := articleUseCase.List(repo)
	articleList := []string{"alamin", "anika", "anamika"}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(articleList); err != nil {
		panic(err)
	}
}
