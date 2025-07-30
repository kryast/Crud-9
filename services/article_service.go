package services

import "github.com/kryast/Crud-9.git/repositories"

type ArticleService interface {
}

type articleService struct {
	repo repositories.ArticleRepository
}

func NewArticleService(repo repositories.ArticleRepository) ArticleService {
	return &articleService{repo}
}
