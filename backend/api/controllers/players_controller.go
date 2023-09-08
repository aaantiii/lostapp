package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"backend/api/middleware"
	"backend/api/services"
	"backend/api/types"
	"backend/api/util"
)

type PlayersController struct {
	service services.IPlayersService
}

func NewPlayersController(service services.IPlayersService) *PlayersController {
	return &PlayersController{service: service}
}

func (controller *PlayersController) setupWithRouter(router *gin.Engine) {
	const rgName = "players"

	memberRoutes := router.Group(rgName, middleware.AuthMiddleware(types.AuthRoleMember))
	memberRoutes.GET("", controller.GETPlayers)
	memberRoutes.GET(":tag", controller.GETPlayerByTag)
	memberRoutes.GET("leaderboard/:statsId", controller.GETLeaderboard)
	memberRoutes.GET("comparable-stats", controller.GETComparableStats)
}

// GETPlayers responds with a slice of coc.Player.
// Filter query parameters: discordId, name, clanTag, page (+itemsPerPage).
func (controller *PlayersController) GETPlayers(c *gin.Context) {
	var params types.PlayersParams
	if err := c.Bind(&params); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	players, err := controller.service.Players(params)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	if params.Page != 0 {
		c.JSON(http.StatusOK, types.NewPaginatedResponse(players, params.PaginationParams))
		return
	}

	c.JSON(http.StatusOK, players)
}

func (controller *PlayersController) GETPlayerByTag(c *gin.Context) {
	tag, err := util.TagFromQuery(c, "tag")
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if player, err := controller.service.PlayerByTag(tag); err == nil {
		c.JSON(http.StatusOK, player)
		return
	}

	c.AbortWithStatus(http.StatusNotFound)
}

func (controller *PlayersController) GETLeaderboard(c *gin.Context) {
	var params types.LeaderboardPlayersParams
	if err := c.Bind(&params); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var err error
	if params.StatsID, err = strconv.Atoi(c.Param("statsId")); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if achievements, err := controller.service.PlayersLeaderboard(params); err == nil {
		c.JSON(http.StatusOK, types.NewPaginatedResponse(achievements, params.PaginationParams))
		return
	}

	c.AbortWithStatus(http.StatusNotFound)
}

func (controller *PlayersController) GETComparableStats(c *gin.Context) {
	c.JSON(http.StatusOK, types.ComparableStats)
}
