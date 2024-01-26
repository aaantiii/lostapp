package util

import (
	"github.com/gin-gonic/gin"

	"github.com/aaantiii/lostapp/backend/api/types"
)

const SessionKey = "session" // SessionKey is the key used to store the session in the gin.Context.

func SessionFromContext(c *gin.Context) *types.Session {
	return c.MustGet(SessionKey).(*types.Session)
}
