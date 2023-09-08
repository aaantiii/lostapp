package controllers

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"backend/api/middleware"
	"backend/api/services"
	"backend/api/types"
	"backend/api/util"
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
		GET("members/kickpoints", controller.GETActiveClanKickpoints).
		GET("members/:memberTag/kickpoints", controller.GETActiveMemberKickpoints).
		GET("settings", controller.GETClanSettings)

	coLeaderRoutes := router.Group(rgName, middleware.AuthMiddleware(types.AuthRoleLeader))
	coLeaderRoutes.GET("leading", controller.GETLeadingClans)
	coLeaderRoutes.Group(":clanTag").
		POST("members", controller.POSTClanMember).
		PUT("settings", controller.PUTClanSettings).
		PUT("members/:memberTag", controller.PUTClanMember)

	leaderRoutes := router.Group(rgName, middleware.AuthMiddleware(types.AuthRoleLeader), middleware.ClanLeaderMiddleware("clanTag"))
	leaderRoutes.PATCH(":clanTag/members/:memberTag/role", controller.PATCHClanMemberRole)
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

	if kickpointList, err := controller.kickpointsService.ActiveClanKickpoints(clanTag); err == nil {
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

	if settings, err := controller.service.ClanSettings(clanTag); err == nil {
		c.JSON(http.StatusOK, settings)
		return
	}

	c.AbortWithStatus(http.StatusNotFound)
}

func (controller *ClansController) PUTClanSettings(c *gin.Context) {
	tag, err := util.TagFromQuery(c, "clanTag")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var settings *types.UpdateClanSettingsPayload
	if err = c.Bind(&settings); err != nil {
		log.Print(err.Error())
		return
	}

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

func (controller *ClansController) POSTClanMember(c *gin.Context) {
	var payload types.AddMemberPayload
	if err := c.Bind(&payload); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if !payload.Role.IsElder() && !payload.Role.IsMember() {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	session := util.SessionFromContext(c)
	if session.AuthRole != types.AuthRoleAdmin && !controller.membersService.MemberIsLeadingClan(session.DiscordUser.ID, payload.ClanTag) {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	payload.AddedByDiscordID = session.DiscordUser.ID
	err := controller.membersService.CreateMember(payload)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.Status(http.StatusCreated)
}

func (controller *ClansController) PUTClanMember(c *gin.Context) {}

func (controller *ClansController) PATCHClanMemberRole(c *gin.Context) {}
