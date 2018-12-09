package articles

import (
	"encoding/json"
	"net/http"

	"github.com/alamin-mahamud/go-bootstrap/repository"
	articleUseCase "github.com/alamin-mahamud/go-bootstrap/usecase/articles"
)

// List - passes to the actual usecase for getting all the list with the repository.
func List(w http.ResponseWriter, r *http.Request) {
	// get Repo
	repo := &repository.MongoDB{}

	// call UseCase with Repo
	articleList := articleUseCase.List(repo)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(articleList); err != nil {
		panic(err)
	}
}
