package api

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"backend/api/controllers"
	"backend/env"
)

// router is the HTTP router used to serve the API.
var router *gin.Engine

// Init initializes router with a new gin.Engine instance.
func Init() error {
	if env.DEBUG.Value() != "true" {
		gin.SetMode(gin.ReleaseMode)
	}

	router = gin.Default()
	if err := router.SetTrustedProxies(nil); err != nil {
		return err
	}

	corsConfig := cors.Config{
		AllowOrigins:     []string{env.FRONTEND_URL.Value()},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type: application/xml", "Origin"},
		AllowHeaders:     []string{"Content-Length", "Content-Type", "Origin"},
		AllowCredentials: true,
	}
	if err := corsConfig.Validate(); err != nil {
		return err
	}
	router.Use(cors.New(corsConfig))

	if err := controllers.Setup(router); err != nil {
		return err
	}

	return nil
}

// ListenAndServe starts the HTTP server.
func ListenAndServe() error {
	port := env.PORT.Value()
	log.Printf("HTTP server listening on port %s", port)
	return router.Run(":" + port)
}
