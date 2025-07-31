package services

import (
	"github.com/kryast/Crud-9.git/models"
	"github.com/kryast/Crud-9.git/repositories"
)

type ArticleService interface {
	Create(article *models.Article) error
}

type articleService struct {
	repo repositories.ArticleRepository
}

func NewArticleService(repo repositories.ArticleRepository) ArticleService {
	return &articleService{repo}
}

func (a *articleService) Create(article *models.Article) error {
	return a.repo.Create(article)
}
