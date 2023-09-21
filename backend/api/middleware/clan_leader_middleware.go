package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/api/types"
	"backend/api/util"
)

// ClanLeaderMiddleware is a middleware that checks if the user is a leader of the clan.
func ClanLeaderMiddleware(clanTagKey string, allowCoLeaders bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := util.SessionFromContext(c)
		if session.AuthRole == types.AuthRoleAdmin {
			defer c.Next()
			return
		}

		clanTag, err := util.TagFromQuery(c, clanTagKey)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		guild, err := authService.Guild(clanTag)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		if guild.IsLeader(session.DiscordUser.Roles) || (allowCoLeaders && guild.IsCoLeader(session.DiscordUser.Roles)) {
			defer c.Next()
			return
		}

		c.AbortWithStatus(http.StatusForbidden)
	}
}
