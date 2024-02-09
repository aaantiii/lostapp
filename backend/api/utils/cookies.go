package utils

import (
	"github.com/gin-gonic/gin"

	"github.com/aaantiii/lostapp/backend/env"
)

type Cookie string

const (
	SessionCookie Cookie = "session"
	StateCookie   Cookie = "state"
)

func (cookie Cookie) Name() string {
	return string(cookie)
}

func (cookie Cookie) Set(c *gin.Context, value string, maxAge int) {
	c.SetCookie(cookie.Name(), value, maxAge, "/", env.DOMAIN.Value(), true, true)
}

func (cookie Cookie) Value(c *gin.Context) (string, error) {
	return c.Cookie(cookie.Name())
}

func (cookie Cookie) Invalidate(c *gin.Context) {
	cookie.Set(c, "", -1)
}
