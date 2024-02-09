package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aaantiii/lostapp/backend/api/middleware"
	"github.com/aaantiii/lostapp/backend/api/services"
	"github.com/aaantiii/lostapp/backend/api/types"
	"github.com/aaantiii/lostapp/backend/api/utils"
)

type ClansController struct {
	service services.IClansService
}

func NewClansController(service services.IClansService) *ClansController {
	return &ClansController{service: service}
}

func (controller *ClansController) setupWithRouter(router *gin.Engine) {
	const rgName = "clans"

	authedRoutes := router.Group(rgName, middleware.AuthMiddleware(false))
	authedRoutes.
		GET("", middleware.PaginationMiddleware(true), controller.GETClans).
		GET(":clanTag/settings", controller.GETClanSettings)

	memberRoutes := authedRoutes.Group(":clanTag", middleware.ClanAuthMiddleware(types.AuthRoleMember))
	memberRoutes.
		GET("", controller.GETClanByTag).
		GET("members/kickpoints", controller.GETActiveClanKickpoints).
		GET("members/:memberTag/kickpoints", controller.GETActiveMemberKickpoints)

	coLeaderRoutes := authedRoutes.Group(":clanTag", middleware.ClanAuthMiddleware(types.AuthRoleCoLeader))
	coLeaderRoutes.
		PUT("settings", controller.PUTClanSettings)
	coLeaderRoutes.Group("members/:memberTag/kickpoints").
		POST("", controller.POSTKickpoint).
		PUT(":id", controller.PUTKickpoint).
		DELETE(":id", controller.DELETEKickpoint)
}

func (controller *ClansController) GETClans(c *gin.Context) {
	var params types.ClansParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.Set(middleware.ErrorKey, err)
		return
	}
	params.PaginationParams = utils.PaginationFromContext(c)

	clans, err := controller.service.Clans(params)
	if err != nil {
		c.Set(middleware.ErrorKey, err)
		return
	}

	c.JSON(http.StatusOK, clans)
}

func (controller *ClansController) GETClanByTag(c *gin.Context) {
	clanTag, err := utils.TagFromParams(c, "clanTag")
	if err != nil {
		c.Set(middleware.ErrorKey, types.ErrBadRequest)
		return
	}

	clan, err := controller.service.ClanByTag(clanTag)
	if err != nil {
		c.Set(middleware.ErrorKey, err)
		return
	}

	c.JSON(http.StatusOK, clan)
}

func (controller *ClansController) GETActiveClanKickpoints(c *gin.Context) {
	clanTag, err := utils.TagFromParams(c, "clanTag")
	if err != nil {
		c.Set(middleware.ErrorKey, types.ErrBadRequest)
		return
	}

	kickpoints, err := controller.service.ActiveClanKickpoints(clanTag)
	if err != nil {
		c.Set(middleware.ErrorKey, err)
		return
	}

	c.JSON(http.StatusOK, kickpoints)
}

func (controller *ClansController) GETActiveMemberKickpoints(c *gin.Context) {
	clanTag, err := utils.TagFromParams(c, "clanTag")
	if err != nil {
		c.Set(middleware.ErrorKey, types.ErrBadRequest)
		return
	}

	memberTag, err := utils.TagFromParams(c, "memberTag")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	kickpointList, err := controller.service.ActiveClanMemberKickpoints(memberTag, clanTag)
	if err != nil {
		c.Set(middleware.ErrorKey, err)
		return
	}

	c.JSON(http.StatusOK, kickpointList)
}

func (controller *ClansController) GETClanSettings(c *gin.Context) {
	clanTag, err := utils.TagFromParams(c, "clanTag")
	if err != nil {
		c.Set(middleware.ErrorKey, types.ErrBadRequest)
		return
	}

	settings, err := controller.service.ClanSettings(clanTag)
	if err != nil {
		c.Set(middleware.ErrorKey, err)
		return
	}

	c.JSON(http.StatusOK, settings)
}

func (controller *ClansController) PUTClanSettings(c *gin.Context) {
	tag, err := utils.TagFromParams(c, "clanTag")
	if err != nil {
		c.Set(middleware.ErrorKey, types.ErrBadRequest)
		return
	}

	var payload *types.UpdateClanSettingsPayload
	if err = c.ShouldBind(&payload); err != nil {
		c.Set(middleware.ErrorKey, types.ErrValidationFailed)
		return
	}

	session := utils.SessionFromContext(c)
	if err = controller.service.UpdateClanSettings(tag, session.User.ID, payload); err != nil {
		c.Set(middleware.ErrorKey, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (controller *ClansController) POSTKickpoint(c *gin.Context) {
	clanTag, err := utils.TagFromParams(c, "clanTag")
	if err != nil {
		c.Set(middleware.ErrorKey, types.ErrBadRequest)
		return
	}

	playerTag, err := utils.TagFromParams(c, "memberTag")
	if err != nil {
		c.Set(middleware.ErrorKey, types.ErrBadRequest)
		return
	}

	var payload *types.CreateKickpointPayload
	if err = c.ShouldBind(&payload); err != nil {
		c.Set(middleware.ErrorKey, types.ErrValidationFailed)
		return
	}

	session := utils.SessionFromContext(c)
	payload.PlayerTag = playerTag
	payload.ClanTag = clanTag
	payload.CreatedByDiscordID = session.User.ID

	if err = controller.service.CreateKickpoint(payload); err != nil {
		c.Set(middleware.ErrorKey, err)
		return
	}

	c.Status(http.StatusCreated)
}

func (controller *ClansController) PUTKickpoint(c *gin.Context) {
	id, err := utils.UintFromParams(c, "id")
	if err != nil {
		c.Set(middleware.ErrorKey, types.ErrBadRequest)
		return
	}

	var payload *types.UpdateKickpointPayload
	if err = c.ShouldBind(&payload); err != nil {
		c.Set(middleware.ErrorKey, types.ErrValidationFailed)
		return
	}

	session := utils.SessionFromContext(c)
	payload.UpdatedByDiscordID = session.User.ID
	if err = controller.service.UpdateKickpoint(id, payload); err != nil {
		c.Set(middleware.ErrorKey, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (controller *ClansController) DELETEKickpoint(c *gin.Context) {
	id, err := utils.UintFromParams(c, "id")
	if err != nil {
		c.Set(middleware.ErrorKey, types.ErrBadRequest)
		return
	}

	if err = controller.service.DeleteKickpoint(id); err != nil {
		c.Set(middleware.ErrorKey, err)
		return
	}

	c.Status(http.StatusNoContent)
}
