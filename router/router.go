package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kryast/Crud-9.git/handlers"
	"github.com/kryast/Crud-9.git/repositories"
	"github.com/kryast/Crud-9.git/services"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	repo := repositories.NewArticleRepository(db)
	svc := services.NewArticleService(repo)
	h := handlers.NewArticleHandler(svc)

	r.POST("/article", h.Create)
	return r
}
