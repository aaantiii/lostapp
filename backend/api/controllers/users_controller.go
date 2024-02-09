package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aaantiii/lostapp/backend/api/middleware"
	"github.com/aaantiii/lostapp/backend/api/services"
)

type UsersController struct {
	service services.IUsersService
}

func NewUsersController(service services.IUsersService) *UsersController {
	return &UsersController{service: service}
}

func (controller *UsersController) setupWithRouter(router *gin.Engine) {
	const rgName = "users"

	userRoutes := router.Group(rgName, middleware.AuthMiddleware(false))
	userRoutes.
		GET(":id", controller.GETUserByID)
}

func (controller *UsersController) GETUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := controller.service.UserByID(id)
	if err != nil {
		c.Set(middleware.ErrorKey, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
