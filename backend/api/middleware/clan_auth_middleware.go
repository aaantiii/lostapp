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
func ClanAuthMiddleware(requiredRole types.AuthRole) func(*gin.Context) {
	return func(c *gin.Context) {
		session := utils.SessionFromContext(c)
		if session.User.IsAdmin {
			return
		}

		clanTag := c.Param("clanTag")
		if err := checkClanRole(requiredRole, session.User, clanTag); err != nil {
			c.Set(ErrorKey, err)
			c.Abort()
		}
	}
}

func checkClanRole(requiredRole types.AuthRole, user *types.AuthUser, clanTag string) *types.ApiError {
	switch requiredRole {
	case types.AuthRoleLeader:
		if slices.Contains(user.LeaderOf, clanTag) {
			return nil
		}
		return types.NewApiError(http.StatusForbidden, "Um diese Aktion ausführen zu können, musst du Anführer dieses Clans sein.")
	case types.AuthRoleCoLeader:
		if slices.Contains(user.CoLeaderOf, clanTag) {
			return nil
		}
		return types.NewApiError(http.StatusForbidden, "Um diese Aktion ausführen zu können, musst du Vize-Anführer dieses Clans sein.")
	case types.AuthRoleMember:
		if slices.Contains(user.MemberOf, clanTag) {
			return nil
		}
		return types.NewApiError(http.StatusForbidden, "Um diese Aktion ausführen zu können, musst du Mitglied dieses Clans sein.")
	}
	return types.ErrNotMember
}
