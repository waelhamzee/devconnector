package http

import (
	"github.com/gin-gonic/gin"
)

// ErrorResponse represents a standard error response
func ErrorResponse(ctx *gin.Context, status int, err error) {
	ctx.JSON(status, gin.H{
		"error": err.Error(),
	})
}
