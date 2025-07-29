package user

import (
	"github.com/gin-gonic/gin"
	httphelper "github.com/waelhamzee/devconnector/internal/http"
	"gorm.io/gorm"
)

func RegisterRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	repo := NewUserRepository(db)
	svc := NewUserService(repo)
	handler := NewUserController(svc)

	users := rg.Group("/users")
	users.Use(httphelper.JWTAuthMiddleware())
	{
		users.GET(":id", handler.GetUser)
	}
}
