package middleware

import (
	"net/http"

	"github.com/amaanq/coc.go"
	"github.com/gin-gonic/gin"

	"backend/api/types"
	"backend/api/util"
)

// ClanLeaderMiddleware is a middleware that checks if the user is the leader of the clan.
func ClanLeaderMiddleware(clanTagKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := util.SessionFromContext(c)
		if session.AuthRole == types.AuthRoleAdmin {
			c.Next()
			return
		}

		clanTag, err := util.TagFromQuery(c, clanTagKey)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}

		clanMembers, err := clansService.ClanMembers(clanTag)
		if err != nil {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		players, err := playersService.PlayersByDiscordID(session.DiscordUser.ID)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		for _, member := range clanMembers {
			for _, player := range players {
				if member.Tag != player.Tag {
					continue
				}
				if member.Role != coc.Leader {
					break
				}
				c.Next()
				return
			}
		}
	}
}
