package articles

import (
	articleEntity "github.com/alamin-mahamud/go-bootstrap/entity/article"
	"github.com/alamin-mahamud/go-bootstrap/repository"
)

// List Usecase for Articles
func List(repo repository.IArticleRepo) []articleEntity.Article {
	articlesList, _ := repo.List()
	return articlesList
}
