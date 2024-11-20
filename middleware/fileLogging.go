package middleware

import (
	"log"
	"os"
	"time"
	"github.com/gin-gonic/gin"
)

func FileLoggerMiddleware() gin.HandlerFunc {
	// os.O_APPEND: Appends logs to the file if it already exists.
	// os.O_CREATE: Creates the file if it doesn't exist.
	// os.O_WRONLY: Opens the file in write-only mode.
	// 0666: Sets the file permissions (readable and writable by all users).
	file, err := os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}

	// log.LstdFlags option includes standard logging flags like the date and time.
	logger := log.New(file, "", log.LstdFlags)

	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		endTime := time.Now()

		latency := endTime.Sub(startTime)

		logger.Printf(
			"[%s] %s %s %d %s\n",
			c.ClientIP(),
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			latency,
		)
	}
}