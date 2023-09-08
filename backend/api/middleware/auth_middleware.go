package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/api/types"
	"backend/api/util"
)

// AuthMiddleware stellt die Authentication und Authorization middleware für eingehende Requests zur Verfügung.
// AuthMiddleware stellt sicher, dass der Session per Discord authentifiziert ist und die benötigte
// types.AuthRole für die angeforderte Resource hat.
func AuthMiddleware(requiredRole types.AuthRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := util.SessionCookie.Value(c)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		session, exists := authService.Session(token)
		if !exists {
			session, err = authService.CreateSession(token)
			if err != nil {
				util.SessionCookie.Invalidate(c)
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}

		if session.LastRefreshed() > authService.Config().MaxSessionDataAge {
			if session, err = authService.RefreshSession(token); err != nil {
				util.SessionCookie.Invalidate(c)
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		}

		if !session.HasPermission(requiredRole) {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}

		defer c.Next()
		c.Set(util.SessionKey, session)
	}
}
