package router

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// repo := repositories.NewArticleRepository(db)
	// svc := services.NewArticleService(repo)
	// h := handlers.NewArticleHandler(svc)

	return r
}
