package middleware

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"

	"github.com/aaantiii/lostapp/backend/api/types"
	"github.com/aaantiii/lostapp/backend/api/util"
)

// NewClanAuthMiddleware returns a gin.HandlerFunc that checks if user has requiredRole within clan.
// Clan tag is taken from the :clanTag param in gin.Context.
func NewClanAuthMiddleware(requiredRole types.AuthRole) func(c *gin.Context) {
	return func(c *gin.Context) {
		session := util.SessionFromContext(c)
		clanTag := c.Param("clanTag")

		switch requiredRole {
		case types.AuthRoleLeader:
			if !slices.Contains(session.User.LeaderOf, clanTag) {
				c.Set(util.ErrorKey, types.NewApiError(http.StatusForbidden, "Um diese Aktion ausführen zu können, musst du Anführer dieses Clans sein."))
				c.Abort()
				return
			}
		case types.AuthRoleCoLeader:
			if !slices.Contains(session.User.CoLeaderOf, clanTag) {
				c.Set(util.ErrorKey, types.NewApiError(http.StatusForbidden, "Um diese Aktion ausführen zu können, musst du Vize-Anführer dieses Clans sein."))
				c.Abort()
				return
			}
		case types.AuthRoleMember:
			if !slices.Contains(session.User.MemberOf, clanTag) {
				c.Set(util.ErrorKey, types.NewApiError(http.StatusForbidden, "Um diese Aktion ausführen zu können, musst du Mitglied dieses Clans sein."))
				c.Abort()
				return
			}
		default:
			c.Set(util.ErrorKey, types.ErrNotMember)
			c.Abort()
			return
		}
	}
}
