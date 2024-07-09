package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aaantiii/lostapp/backend/api/middleware"
	"github.com/aaantiii/lostapp/backend/api/services"
	"github.com/aaantiii/lostapp/backend/api/types"
	"github.com/aaantiii/lostapp/backend/api/utils"
)

type PlayersController struct {
	service services.IPlayersService
}

func NewPlayersController(service services.IPlayersService) *PlayersController {
	return &PlayersController{service: service}
}

func (controller *PlayersController) setupWithRouter(router *gin.Engine) {
	const rgName = "players"

	authedRoutes := router.Group(rgName, middleware.AuthMiddleware(false))
	authedRoutes.
		GET("", middleware.PaginationMiddleware(false), controller.GETPlayers).
		GET(":tag", controller.GETPlayerByTag).
		GET("stats/list", controller.GETStatsList)

	authedRoutes.Group("live").
		GET("", middleware.PaginationMiddleware(false), controller.GETLivePlayers).
		GET("stats/leaderboard", middleware.PaginationMiddleware(false), controller.GETLeaderboard)
}

func (controller *PlayersController) GETPlayers(c *gin.Context) {
	var params types.PlayersParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.Set(middleware.ErrorKey, types.ErrValidationFailed)
		return
	}
	params.PaginationParams = utils.PaginationFromContext(c)

	res, err := controller.service.Players(params)
	if err != nil {
		c.Set(middleware.ErrorKey, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (controller *PlayersController) GETLivePlayers(c *gin.Context) {
	var params types.PlayersParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.Set(middleware.ErrorKey, types.ErrValidationFailed)
		return
	}
	params.PaginationParams = utils.PaginationFromContext(c)

	res, err := controller.service.PlayersLive(params)
	if err != nil {
		c.Set(middleware.ErrorKey, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (controller *PlayersController) GETPlayerByTag(c *gin.Context) {
	tag, err := utils.TagFromParams(c, "tag")
	if err != nil {
		c.Set(middleware.ErrorKey, types.ErrBadRequest)
		return
	}

	player, err := controller.service.PlayerByTag(tag)
	if err != nil {
		c.Set(middleware.ErrorKey, err)
		return
	}

	c.JSON(http.StatusOK, player)
}

func (controller *PlayersController) GETLeaderboard(c *gin.Context) {
	var params types.LeaderboardParams
	if err := c.ShouldBindQuery(&params); err != nil {
		c.Set(middleware.ErrorKey, types.ErrValidationFailed)
		return
	}
	params.PaginationParams = utils.PaginationFromContext(c)

	compStat, err := utils.ComparableStatisticByName(params.StatName)
	if err != nil {
		c.Set(middleware.ErrorKey, types.ErrBadRequest)
		return
	}

	stats, err := controller.service.Leaderboard(params, compStat)
	if err != nil {
		c.Set(middleware.ErrorKey, err)
		return
	}

	c.JSON(http.StatusOK, stats)
}

func (controller *PlayersController) GETStatsList(c *gin.Context) {
	c.JSON(http.StatusOK, types.ComparableStats)
}
