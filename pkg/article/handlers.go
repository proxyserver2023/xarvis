package article

import (
	"net/http"
)

var uc IArticleUsecase

// List -
func List(w http.ResponseWriter, r *http.Request) {
	rp := NewMongoDB()
	uc = NewArticleUsecase(rp)
	uc.List()
}
