package controllers

import (
	"github.com/gin-gonic/gin"

	"github.com/aaantiii/lostapp/backend/api/middleware"
	"github.com/aaantiii/lostapp/backend/api/repos"
	"github.com/aaantiii/lostapp/backend/api/services"
	"github.com/aaantiii/lostapp/backend/client"
	"github.com/aaantiii/lostapp/backend/store/postgres"
)

type Controller interface {
	setupWithRouter(*gin.Engine)
}

// SetupWithRouter initializes all controllers and runs Controller.setupWithRouter for every controller.
func SetupWithRouter(router *gin.Engine) error {
	db, err := postgres.NewClient()
	if err != nil {
		return err
	}
	clashClient, err := client.NewClashClient()
	if err != nil {
		return err
	}

	authService := services.NewAuthService(
		repos.NewGuildsRepo(db),
		repos.NewUsersRepo(db),
	)
	middleware.UseAuthService(authService)

	controllers := []Controller{
		NewAuthController(authService),
		NewUsersController(services.NewUsersService(
			repos.NewUsersRepo(db),
		)),
		NewClansController(services.NewClansService(
			repos.NewClansRepo(db),
			repos.NewKickpointsRepo(db),
			repos.NewClanSettingsRepo(db),
			repos.NewClanEventsRepo(db),
			repos.NewMembersRepo(db),
			clashClient,
		)),
		NewPlayersController(services.NewPlayersService(
			repos.NewPlayersRepo(db),
			repos.NewMembersRepo(db),
			clashClient,
		)),
	}

	// setup Controllers
	for _, controller := range controllers {
		controller.setupWithRouter(router)
	}

	return nil
}
