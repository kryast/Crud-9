package services

import (
	"github.com/kryast/Crud-9.git/models"
	"github.com/kryast/Crud-9.git/repositories"
)

type ArticleService interface {
	Create(article *models.Article) error
	GetAll() ([]models.Article, error)
	GetByID(id uint) (*models.Article, error)
}

type articleService struct {
	repo repositories.ArticleRepository
}

func NewArticleService(repo repositories.ArticleRepository) ArticleService {
	return &articleService{repo}
}

func (as *articleService) Create(article *models.Article) error {
	return as.repo.Create(article)
}

func (as *articleService) GetAll() ([]models.Article, error) {
	return as.repo.GetAll()
}

func (as *articleService) GetByID(id uint) (*models.Article, error) {
	return as.repo.GetByID(id)
}
