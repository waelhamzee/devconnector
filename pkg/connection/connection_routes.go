package connection

import (
	"github.com/gin-gonic/gin"
	httphelper "github.com/waelhamzee/devconnector/internal/http"
	"gorm.io/gorm"
)

func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	handler := NewController(svc)

	connections := rg.Group("/connections")
	connections.Use(httphelper.JWTAuthMiddleware())
	{
		connections.POST("/", handler.CreateConnection)
		connections.GET("/user/:user_id", handler.ListConnections)
		connections.DELETE("/user/:user_id/target/:target_id", handler.DeleteConnection)
	}
}
