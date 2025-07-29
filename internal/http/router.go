package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/waelhamzee/devconnector/internal/config"
)

func NewRouter(cfg config.Config) *gin.Engine {
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Recovery())

	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return r
}
