package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/aaantiii/lostapp/backend/api/middleware"
	"github.com/aaantiii/lostapp/backend/api/services"
	"github.com/aaantiii/lostapp/backend/api/types"
	"github.com/aaantiii/lostapp/backend/api/utils"
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
	publicRoutes.
		GET("discord/login", controller.GETDiscordLogin).
		GET("discord/callback", controller.GETDiscordCallback)

	userRoutes := router.Group(rgName, middleware.AuthMiddleware(false))
	userRoutes.
		GET("discord/logout", controller.GETDiscordLogout).
		GET("session", controller.GETSession)
}

func (*AuthController) GETSession(c *gin.Context) {
	c.JSON(http.StatusOK, utils.SessionFromContext(c))
}

func (controller *AuthController) GETDiscordLogin(c *gin.Context) {
	randomState := uuid.NewString()
	utils.StateCookie.Set(c, randomState, 600)
	c.Redirect(http.StatusSeeOther, controller.service.AuthCodeURL(randomState))
}

func (controller *AuthController) GETDiscordCallback(c *gin.Context) {
	targetState, err := utils.StateCookie.Value(c)
	if err != nil {
		utils.FERouteLoginFailed.RedirectWithStatus(c, http.StatusSeeOther)
		return
	}

	utils.StateCookie.Invalidate(c)
	currentState := c.Query("state")
	if currentState != targetState {
		utils.FERouteLoginFailed.RedirectWithStatus(c, http.StatusSeeOther)
		return
	}

	token, err := controller.service.ExchangeCode(c.Query("code"))
	if err != nil {
		utils.FERouteLoginFailed.RedirectWithStatus(c, http.StatusSeeOther)
		return
	}

	utils.SessionCookie.Set(c, token.AccessToken, 600000)
	utils.FERouteIndex.RedirectWithStatus(c, http.StatusSeeOther)
}

func (controller *AuthController) GETDiscordLogout(c *gin.Context) {
	token, err := utils.SessionCookie.Value(c)
	if err != nil {
		c.Set(middleware.ErrorKey, types.ErrNotSignedIn)
		return
	}

	utils.SessionCookie.Invalidate(c)
	if err = controller.service.RevokeSession(token); err != nil {
		c.Set(middleware.ErrorKey, types.ErrSignOutFailed)
		return
	}

	c.Status(http.StatusOK)
}
