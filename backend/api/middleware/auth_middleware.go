package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aaantiii/lostapp/backend/api/types"
	"github.com/aaantiii/lostapp/backend/api/util"
)

// AuthMiddleware is a gin middleware that checks if the user is authenticated.
func AuthMiddleware(requiredRole types.AuthRole) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := util.SessionCookie.Value(c)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		session, ok := authService.Session(token)
		if !ok {
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

		if session.User.IsAdmin {
			c.Set(util.SessionKey, session)
			return
		}

		if requiredRole == types.AuthRoleAdmin && !session.User.IsAdmin {
			c.Set(util.ErrorKey, types.ErrAdminRequired)
			c.Abort()
			return
		}

		c.Set(util.SessionKey, session)
	}
}
