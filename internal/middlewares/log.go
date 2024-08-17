package middlewares

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger middleware
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()

		// Process request
		c.Next()

		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		// Get request status code
		status := c.Writer.Status()

		log.Printf("[GIN] %v | %3d | %13v | %15s | %-7s %#v\n%s",
			end.Format("2006/01/02 - 15:04:05"),
			status,
			latency,
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL,
			c.Errors.String(),
		)
	}
}
