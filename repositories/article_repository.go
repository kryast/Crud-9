package repositories

import (
	"github.com/kryast/Crud-9.git/models"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	Create(article *models.Article) error
	GetAll() ([]models.Article, error)
	GetByID(id uint) (*models.Article, error)
	Update(article *models.Article) error
	Delete(id uint) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db}
}

func (ar *articleRepository) Create(article *models.Article) error {
	return ar.db.Create(article).Error
}

func (ar *articleRepository) GetAll() ([]models.Article, error) {
	var articles []models.Article
	err := ar.db.Find(&articles).Error
	return articles, err
}

func (ar *articleRepository) GetByID(id uint) (*models.Article, error) {
	var article models.Article
	err := ar.db.First(&article, id).Error
	return &article, err
}

func (ar *articleRepository) Update(article *models.Article) error {
	return ar.db.Save(article).Error
}

func (ar *articleRepository) Delete(id uint) error {
	return ar.db.Delete(&models.Article{}, id).Error
}
