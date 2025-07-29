package comment

import (
	"github.com/gin-gonic/gin"
	httphelper "github.com/waelhamzee/devconnector/internal/http"
	"gorm.io/gorm"
)

func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	handler := NewController(svc)

	comments := rg.Group("/comments")
	comments.Use(httphelper.JWTAuthMiddleware())
	{
		comments.POST("/", handler.CreateComment)
		comments.GET("/post/:post_id", handler.ListCommentsByPost)
		comments.DELETE(":id", handler.DeleteComment)
	}
}
