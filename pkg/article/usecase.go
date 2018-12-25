package article

import "fmt"

type IArticleUsecase interface {
	List()
}

type ArticleUsecase struct {
	repo IArticleRepository
}

func NewArticleUsecase(r IArticleRepository) *ArticleUsecase {
	return &ArticleUsecase{
		repo: r,
	}
}

func (a *ArticleUsecase) List() {
	a.repo.List()
	fmt.Println("Run as Usecase")
}
