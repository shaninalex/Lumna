package httpx

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		query := c.Request.URL.RequestURI()
		ip := c.ClientIP()
		log.Printf(
			"%s | %3d | %13v | %-7s %s\n",
			ip,
			status,
			latency,
			method,
			query,
		)
	}
}
