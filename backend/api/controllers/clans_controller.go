package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"backend/api/middleware"
	"backend/api/services"
	"backend/api/types"
	"backend/api/util"
	"backend/store/postgres/models"
)

type ClansController struct {
	service           services.IClansService
	kickpointsService services.IKickpointsService
	membersService    services.IMembersService
}

func NewClansController(service services.IClansService, kickpointsService services.IKickpointsService, membersService services.IMembersService) *ClansController {
	return &ClansController{service: service, kickpointsService: kickpointsService, membersService: membersService}
}

func (controller *ClansController) setupWithRouter(router *gin.Engine) {
	const rgName = "clans"

	memberRoutes := router.Group(rgName, middleware.AuthMiddleware(types.AuthRoleMember))
	memberRoutes.GET("", controller.GETClans)
	memberRoutes.Group(":clanTag").
		GET("", controller.GETClanByTag).
		GET("settings", controller.GETClanSettings).
		GET("members/kickpoints", controller.GETActiveClanKickpoints).
		GET("members/:memberTag/kickpoints", controller.GETActiveMemberKickpoints)

	coLeaderRoutes := router.Group(rgName, middleware.AuthMiddleware(types.AuthRoleLeader))
	coLeaderRoutes.GET("leading", controller.GETLeadingClans)

	coLeaderClanRoutes := coLeaderRoutes.Group(":clanTag", middleware.ClanLeaderMiddleware("clanTag", true))
	coLeaderClanRoutes.PUT("settings", controller.PUTClanSettings)

	coLeaderClanRoutes.Group("members/:memberTag/kickpoints").
		POST("", controller.POSTKickpoint).
		PUT(":id", controller.PUTKickpoint).
		DELETE(":id", controller.DELETEKickpoint)
}

func (controller *ClansController) GETClans(c *gin.Context) {
	if clans, err := controller.service.Clans(); err == nil {
		c.JSON(http.StatusOK, clans)
		return
	}

	c.AbortWithStatus(http.StatusInternalServerError)
}

func (controller *ClansController) GETClanByTag(c *gin.Context) {
	clanTag, err := util.TagFromQuery(c, "clanTag")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if clan, err := controller.service.ClanByTag(clanTag); err == nil {
		c.JSON(http.StatusOK, clan)
		return
	}

	c.AbortWithStatus(http.StatusNotFound)
}

func (controller *ClansController) GETActiveClanKickpoints(c *gin.Context) {
	clanTag, err := util.TagFromQuery(c, "clanTag")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if kickpointList, err := controller.kickpointsService.ActiveClanMemberKickpoints(clanTag); err == nil {
		c.JSON(http.StatusOK, kickpointList)
		return
	}

	c.AbortWithStatus(http.StatusNotFound)
}

func (controller *ClansController) GETActiveMemberKickpoints(c *gin.Context) {
	clanTag, err := util.TagFromQuery(c, "clanTag")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	memberTag, err := util.TagFromQuery(c, "memberTag")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	kickpointList, err := controller.kickpointsService.ActivePlayerKickpoints(memberTag, clanTag)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.Status(http.StatusNoContent)
		return
	}
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, kickpointList)
}

func (controller *ClansController) GETClanSettings(c *gin.Context) {
	clanTag, err := util.TagFromQuery(c, "clanTag")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var settings *models.LostClanSettings
	if settings, err = controller.service.ClanSettings(clanTag); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, settings)
	return
}

func (controller *ClansController) PUTClanSettings(c *gin.Context) {
	tag, err := util.TagFromQuery(c, "clanTag")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var settings *types.UpdateClanSettingsPayload
	if err = c.Bind(&settings); err != nil {
		return
	}

	session := util.SessionFromContext(c)
	settings.UpdatedByDiscordID = session.DiscordUser.ID
	if err = controller.service.UpdateClanSettings(tag, settings); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Status(http.StatusNoContent)
}

func (controller *ClansController) GETLeadingClans(c *gin.Context) {
	session := util.SessionFromContext(c)

	var clans []*types.Clan
	var err error
	if session.AuthRole == types.AuthRoleAdmin {
		clans, err = controller.service.Clans()
	} else {
		clans, err = controller.service.ClansWhereMemberIsLeader(session.DiscordUser.ID)
	}

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, clans)
}

func (controller *ClansController) POSTKickpoint(c *gin.Context) {
	clanTag, err := util.TagFromQuery(c, "clanTag")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var playerTag string
	if playerTag, err = util.TagFromQuery(c, "memberTag"); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var payload *types.CreateKickpointPayload
	if err := c.Bind(&payload); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	session := util.SessionFromContext(c)
	payload.PlayerTag = playerTag
	payload.ClanTag = clanTag
	payload.AddedByDiscordID = session.DiscordUser.ID

	if err = controller.kickpointsService.CreateKickpoint(payload); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusCreated)
}

func (controller *ClansController) PUTKickpoint(c *gin.Context) {
	idStr := c.Param("id")
	if idStr == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var payload *types.UpdateKickpointPayload
	if err := c.Bind(&payload); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	session := util.SessionFromContext(c)
	payload.UpdatedByDiscordID = session.DiscordUser.ID
	if err := controller.kickpointsService.UpdateKickpoint(uint(id), payload); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusNoContent)
}

func (controller *ClansController) DELETEKickpoint(c *gin.Context) {
	idParam := c.Param("id")
	if idParam == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := controller.kickpointsService.DeleteKickpoint(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusNoContent)
}
