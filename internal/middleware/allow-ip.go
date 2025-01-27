package middleware

import (
	"net/http"

	"github.com/fluffy-melli/MoliDB/internal/runtime"
	"github.com/gin-gonic/gin"
)

func Allow(ip string) bool {
	if len(runtime.Configs.Allow_IP) == 0 {
		return true
	}
	for _, allowip := range runtime.Configs.Allow_IP {
		if allowip == ip {
			return true
		}
	}
	return false
}

func AllowIPMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		if !Allow(ip) {
			c.JSON(http.StatusTooManyRequests, gin.H{"error": "Unallowed connection IP"})
			c.Abort()
			return
		}
		c.Next()
	}
}
