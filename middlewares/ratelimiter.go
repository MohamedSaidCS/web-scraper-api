package middlewares

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

func RateLimiter() gin.HandlerFunc {
	limiter := rate.NewLimiter(1, 4)
	return func(c *gin.Context) {
		if !limiter.Allow() {
			c.JSON(429, gin.H{"message": "Too many requests"})
			c.Abort()
			return
		}
		c.Next()
	}
}
