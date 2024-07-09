package controllers

import (
	"net/http"
	"slices"

	"github.com/aaantiii/goclash"
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
		GET("list", controller.GETClansList)
	authedRoutes.Group(":clanTag").
		GET("", controller.GETClanByTag).
		GET("settings", controller.GETClanSettings).
		GET("events", controller.GETClanEvents)
	authedRoutes.Group("live").
		GET("", middleware.PaginationMiddleware(false), controller.GETLiveClans)

	memberRoutes := authedRoutes.Group(":clanTag", middleware.ClanAuthMiddleware(types.AuthRoleMember))
	memberRoutes.Group("members").
		GET("", controller.GETClanMembers).
		GET("kickpoints", controller.GETActiveClanKickpoints).
		GET(":memberTag/kickpoints", middleware.PaginationMiddleware(true), controller.GETActiveMemberKickpoints)

	coLeaderRoutes := authedRoutes.Group(":clanTag", middleware.ClanAuthMiddleware(types.AuthRoleCoLeader))
	coLeaderRoutes.
		PUT("settings", controller.PUTClanSettings)
	coLeaderRoutes.Group("members").
		POST("", controller.POSTMember).
		DELETE(":memberTag", controller.DELETEMember).
		PATCH(":memberTag", controller.PATCHMember)
	coLeaderRoutes.Group("members/:memberTag/kickpoints").
		POST("", controller.POSTKickpoint).
		PUT(":id", controller.PUTKickpoint).
		DELETE(":id", controller.DELETEKickpoint)
}

func (controller *ClansController) GETClans(c *gin.Context) {
	var params types.ClansParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.Set(middleware.ErrorKey, types.ErrBadRequest)
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

func (controller *ClansController) GETClansList(c *gin.Context) {
	clans, err := controller.service.ClansList(types.ClansParams{})
	if err != nil {
		c.Set(middleware.ErrorKey, err)
		return
	}

	c.JSON(http.StatusOK, clans)
}

func (controller *ClansController) GETLiveClans(c *gin.Context) {
	var params types.ClansParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.Set(middleware.ErrorKey, types.ErrBadRequest)
		return
	}
	params.PaginationParams = utils.PaginationFromContext(c)

	clans, err := controller.service.LiveClans(params)
	if err != nil {
		c.Set(middleware.ErrorKey, err)
		return
	}

	c.JSON(http.StatusOK, clans)
}

func (controller *ClansController) GETClanEvents(c *gin.Context) {
	clanTag, err := utils.TagFromParams(c, "clanTag")
	if err != nil {
		c.Set(middleware.ErrorKey, types.ErrBadRequest)
		return
	}

	pagination := utils.PaginationFromContext(c)
	events, err := controller.service.ClanEvents(clanTag, pagination)
	if err != nil {
		c.Set(middleware.ErrorKey, err)
		return
	}

	c.JSON(http.StatusOK, events)
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

func (controller *ClansController) GETClanMembers(c *gin.Context) {
	clanTag, err := utils.TagFromParams(c, "clanTag")
	if err != nil {
		c.Set(middleware.ErrorKey, types.ErrBadRequest)
		return
	}
	var params types.MembersParams
	if err = c.ShouldBindQuery(&params); err != nil {
		c.Set(middleware.ErrorKey, types.ErrBadRequest)
		return
	}
	params.ClanTag = clanTag

	members, err := controller.service.ClanMembers(params)
	if err != nil {
		c.Set(middleware.ErrorKey, err)
		return
	}

	c.JSON(http.StatusOK, members)
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

	var params types.KickpointParams
	params.PlayerTag = memberTag
	params.ClanTag = clanTag
	params.PaginationParams = utils.PaginationFromContext(c)

	kickpoints, err := controller.service.ActiveClanMemberKickpoints(params)
	if err != nil {
		c.Set(middleware.ErrorKey, err)
		return
	}

	c.JSON(http.StatusOK, kickpoints)
}

func (controller *ClansController) POSTMember(c *gin.Context) {
	var payload *types.CreateMemberPayload
	if err := c.ShouldBind(&payload); err != nil {
		c.Set(middleware.ErrorKey, types.ErrValidationFailed)
		return
	}

	session := utils.SessionFromContext(c)
	if !checkLeaderPerms(session.User, payload.ClanTag, goclash.ClanRoleLeader) {
		c.Set(middleware.ErrorKey, types.ErrNoPermission)
		return
	}

	if err := controller.service.CreateMember(payload, session.User.ID); err != nil {
		c.Set(middleware.ErrorKey, err)
		return
	}

	c.Status(http.StatusCreated)
}

func (controller *ClansController) DELETEMember(c *gin.Context) {
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

	session := utils.SessionFromContext(c)
	if !checkLeaderPerms(session.User, clanTag, goclash.ClanRoleLeader) {
		c.Set(middleware.ErrorKey, types.ErrNoPermission)
		return
	}

	if err = controller.service.DeleteMember(playerTag, clanTag); err != nil {
		c.Set(middleware.ErrorKey, err)
		return
	}

	c.Status(http.StatusNoContent)
}

func (controller *ClansController) PATCHMember(c *gin.Context) {
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

	session := utils.SessionFromContext(c)
	if !checkLeaderPerms(session.User, clanTag, goclash.ClanRoleLeader) {
		c.Set(middleware.ErrorKey, types.ErrNoPermission)
		return
	}

	var payload *types.UpdateMemberPayload
	if err = c.ShouldBind(&payload); err != nil {
		c.Set(middleware.ErrorKey, types.ErrValidationFailed)
		return
	}

	payload.PlayerTag = playerTag
	payload.ClanTag = clanTag
	if err = controller.service.UpdateMember(payload); err != nil {
		c.Set(middleware.ErrorKey, err)
		return
	}

	c.Status(http.StatusNoContent)
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

// checkLeaderPerms checks if the user has the required permissions to perform the action, e.g. to add a coLeader the user must be leader.
func checkLeaderPerms(user *types.AuthUser, clanTag string, role goclash.ClanRole) bool {
	if user.IsAdmin {
		return true
	}

	switch role {
	case goclash.ClanRoleMember, goclash.ClanRoleAdmin:
		return true
	case goclash.ClanRoleCoLeader:
		return slices.Contains(user.LeaderOf, clanTag)
	case goclash.ClanRoleLeader:
		return user.IsAdmin
	default:
		return false
	}
}
