package post

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	httphelper "github.com/waelhamzee/devconnector/internal/http"
)

func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	handler := NewController(svc)

	posts := rg.Group("/posts")
	posts.Use(httphelper.JWTAuthMiddleware())
	{
		posts.POST("/", handler.CreatePost)
		posts.GET("/", handler.ListPosts)
		posts.GET(":id", handler.GetPost)
		posts.DELETE(":id", handler.DeletePost)
	}
}
