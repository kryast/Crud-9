package repositories

import "gorm.io/gorm"

type ArticleRepository interface {
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db}
}
