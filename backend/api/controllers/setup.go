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

	// create repos
	usersRepo := repos.NewUsersRepo(db)
	clanSettingsRepo := repos.NewClanSettingsRepo(db)
	clansRepo := repos.NewClansRepo(db)
	guildsRepo := repos.NewGuildsRepo(db)
	kickpointsRepo := repos.NewKickpointsRepo(db)
	membersRepo := repos.NewMembersRepo(db)
	playersRepo := repos.NewPlayersRepo(db)

	// create services
	membersService := services.NewMembersService(membersRepo)
	authService := services.NewAuthService(guildsRepo, usersRepo)
	clansService := services.NewClansService(clansRepo, playersRepo, clanSettingsRepo)
	kickpointsService := services.NewKickpointsService(kickpointsRepo, playersRepo, clanSettingsRepo)
	playersService := services.NewPlayersService(playersRepo, clansRepo)

	// inject services into middleware
	middleware.InjectServices(authService, clansService)

	// create controllers
	controllers := []Controller{
		NewAuthController(authService),
		NewClansController(clansService, kickpointsService, membersService),
		NewPlayersController(playersService),
	}

	// setup Controllers
	for _, controller := range controllers {
		controller.setupWithRouter(router)
	}

	return nil
}
