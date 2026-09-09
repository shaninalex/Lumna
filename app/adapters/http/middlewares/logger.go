package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// process request
		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()
		method := c.Request.Method
		query := c.Request.URL.RequestURI()
		ip := c.ClientIP()

		log.Printf(
			"%s | %3d | %13v | %-7s %s",
			ip,
			status,
			latency,
			method,
			query,
		)
	}
}
