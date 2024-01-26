package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/aaantiii/lostapp/backend/api/middleware"
	"github.com/aaantiii/lostapp/backend/api/services"
	"github.com/aaantiii/lostapp/backend/api/types"
	"github.com/aaantiii/lostapp/backend/api/util"
)

type AuthController struct {
	service services.IAuthService
}

func NewAuthController(service services.IAuthService) *AuthController {
	return &AuthController{service}
}

func (controller *AuthController) setupWithRouter(router *gin.Engine) {
	const rgName = "auth"

	publicRoutes := router.Group(rgName)
	publicRoutes.GET("discord/login", controller.GETDiscordLogin)
	publicRoutes.GET("discord/callback", controller.GETDiscordCallback)

	userRoutes := router.Group(rgName, middleware.AuthMiddleware(types.AuthRoleUser))
	userRoutes.GET("discord/logout", controller.GETDiscordLogout)
	userRoutes.GET("session", controller.GETSession)
}

func (*AuthController) GETSession(c *gin.Context) {
	c.JSON(http.StatusOK, util.SessionFromContext(c))
}

func (controller *AuthController) GETDiscordLogin(c *gin.Context) {
	randomState := uuid.New().String()
	util.StateCookie.Set(c, randomState, 120)

	c.Redirect(http.StatusSeeOther, controller.service.AuthCodeURL(randomState))
}

func (controller *AuthController) GETDiscordCallback(c *gin.Context) {
	targetState, err := util.StateCookie.Value(c)
	if err != nil {
		util.FERouteLoginFailed.RedirectWithStatus(c, http.StatusSeeOther)
		return
	}

	util.StateCookie.Invalidate(c)
	currentState := c.Query("state")
	if currentState != targetState {
		util.FERouteLoginFailed.RedirectWithStatus(c, http.StatusSeeOther)
		return
	}

	token, err := controller.service.ExchangeCode(c.Query("code"))
	if err != nil {
		util.FERouteLoginFailed.RedirectWithStatus(c, http.StatusSeeOther)
		return
	}

	util.SessionCookie.Set(c, token.AccessToken, 600000)
	util.FERouteLoginSuccess.RedirectWithStatus(c, http.StatusSeeOther)
}

func (controller *AuthController) GETDiscordLogout(c *gin.Context) {
	token, err := util.SessionCookie.Value(c)
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	util.SessionCookie.Invalidate(c)
	if err = controller.service.DeleteSession(token); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.Status(http.StatusOK)
}
