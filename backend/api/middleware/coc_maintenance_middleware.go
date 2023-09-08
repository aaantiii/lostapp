package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CocMaintenanceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if clansService.IsMaintenance() {
			c.AbortWithStatus(http.StatusServiceUnavailable)
			return
		}
		defer c.Next()
	}
}
