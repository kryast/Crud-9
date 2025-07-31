package repositories

import (
	"github.com/kryast/Crud-9.git/models"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	Create(article *models.Article) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db}
}

func (a *articleRepository) Create(article *models.Article) error {
	return a.db.Create(article).Error
}
