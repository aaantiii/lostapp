package controllers

import (
	"time"

	"github.com/gin-gonic/gin"

	"backend/api/middleware"
	"backend/api/repos"
	"backend/api/services"
	"backend/store/cache"
	"backend/store/postgres"
)

type Controller interface {
	setupWithRouter(*gin.Engine)
}

// Setup initializes all controllers and runs Controller.setupWithRouter for every controller.
func Setup(router *gin.Engine) error {
	db, err := postgres.NewClient()
	if err != nil {
		return err
	}

	cocCache, err := cache.NewCocCache(db, time.Minute*3)
	if err != nil {
		return err
	}

	// Create repos
	clanSettingsRepo := repos.NewClanSettingsRepo(db)
	membersRepo := repos.NewMembersRepo(db)
	clansRepo := repos.NewClansRepo(db, cocCache)
	kickpointsRepo := repos.NewKickpointsRepo(db, cocCache)
	playersRepo := repos.NewPlayersRepo(db, cocCache)

	// Create services
	membersService := services.NewMembersService(membersRepo)
	authService := services.NewAuthService(membersRepo)
	clansService := services.NewClansService(clansRepo, playersRepo, clanSettingsRepo)
	kickpointsService := services.NewKickpointsService(kickpointsRepo, playersRepo, clanSettingsRepo)
	playersService := services.NewPlayersService(playersRepo, clansRepo)

	// Inject services into middleware and add global middleware to router
	middleware.InjectServices(authService, clansService, playersService)
	router.Use(middleware.CocMaintenanceMiddleware())

	// Create controllers
	controllers := []Controller{
		NewAuthController(authService),
		NewClansController(clansService, kickpointsService, membersService),
		NewPlayersController(playersService),
	}

	// Setup Controllers
	for _, controller := range controllers {
		controller.setupWithRouter(router)
	}

	return nil
}
