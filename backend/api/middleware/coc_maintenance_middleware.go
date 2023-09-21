package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CocMaintenanceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if clansService.IsMaintenance() {
			c.String(http.StatusServiceUnavailable, "Die Clash of Clans API ist momentan unter Wartungsarbeiten. Bitte versuche es später erneut.")
			return
		}
		defer c.Next()
	}
}
