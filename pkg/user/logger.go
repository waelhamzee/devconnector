package user

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		latency := time.Since(t)
		log.Printf("%s %s %d %s", c.Request.Method, c.Request.URL.Path, c.Writer.Status(), latency)
	}
}
