package logging

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		path := context.Request.URL.Path

		if context.Request.URL.RawQuery != "" {
			path = path + "?" + context.Request.URL.RawQuery
		}
		// Process next request
		context.Next()

		// Results
		timestamp := time.Now()
		latency := timestamp.Sub(start)
		clientIP := context.ClientIP()
		method := context.Request.Method
		statusCode := context.Writer.Status()

		fmt.Printf("[HTTP] %v | %3d | %13v | %15s | %-7s %#v\n",
			timestamp.Format("2006/01/02 - 15:04:05"),
			statusCode,
			latency,
			clientIP,
			method,
			path,
		)
	}
}
