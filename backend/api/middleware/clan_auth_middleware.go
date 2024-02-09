package middleware

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"

	"github.com/aaantiii/lostapp/backend/api/types"
	"github.com/aaantiii/lostapp/backend/api/utils"
)

// ClanAuthMiddleware returns a gin.HandlerFunc that checks if user has requiredRole within clan.
// Clan tag is taken from the :clanTag param of the request.
func ClanAuthMiddleware(requiredRole types.AuthRole) func(c *gin.Context) {
	return func(c *gin.Context) {
		session := utils.SessionFromContext(c)
		clanTag := c.Param("clanTag")

		switch requiredRole {
		case types.AuthRoleLeader:
			if !slices.Contains(session.User.LeaderOf, clanTag) {
				c.Set(ErrorKey, types.NewApiError(http.StatusForbidden, "Um diese Aktion ausführen zu können, musst du Anführer dieses Clans sein."))
				c.Abort()
				return
			}
		case types.AuthRoleCoLeader:
			if !slices.Contains(session.User.CoLeaderOf, clanTag) {
				c.Set(ErrorKey, types.NewApiError(http.StatusForbidden, "Um diese Aktion ausführen zu können, musst du Vize-Anführer dieses Clans sein."))
				c.Abort()
				return
			}
		case types.AuthRoleMember:
			if !slices.Contains(session.User.MemberOf, clanTag) {
				c.Set(ErrorKey, types.NewApiError(http.StatusForbidden, "Um diese Aktion ausführen zu können, musst du Mitglied dieses Clans sein."))
				c.Abort()
				return
			}
		default:
			c.Set(ErrorKey, types.ErrNotMember)
			c.Abort()
			return
		}
	}
}
