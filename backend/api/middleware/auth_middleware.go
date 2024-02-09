package middleware

import (
	"github.com/gin-gonic/gin"

	"github.com/aaantiii/lostapp/backend/api/types"
	"github.com/aaantiii/lostapp/backend/api/utils"
)

// AuthMiddleware is a gin middleware that checks if the user is authenticated.
func AuthMiddleware(requiresAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := utils.SessionCookie.Value(c)
		if err != nil {
			c.Set(ErrorKey, types.ErrNotSignedIn)
			c.Abort()
			return
		}

		session, ok := authService.Session(token)
		if !ok {
			session, err = authService.CreateSession(token)
			if err != nil {
				utils.SessionCookie.Invalidate(c)
				c.Set(ErrorKey, types.ErrNotSignedIn)
				c.Abort()
				return
			}
		}

		if session.NeedsRefresh(authService.Config().MaxSessionDataAge) {
			if session, err = authService.RefreshSession(token); err != nil {
				utils.SessionCookie.Invalidate(c)
				c.Set(ErrorKey, types.ErrNotSignedIn)
				c.Abort()
				return
			}
		}

		if requiresAdmin && !session.User.IsAdmin {
			c.Set(ErrorKey, types.ErrAdminRequired)
			c.Abort()
			return
		}

		c.Set(utils.SessionKey, session)
	}
}
