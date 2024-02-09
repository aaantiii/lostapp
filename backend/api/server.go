package api

import (
	"log"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/aaantiii/lostapp/backend/api/controllers"
	"github.com/aaantiii/lostapp/backend/api/middleware"
	"github.com/aaantiii/lostapp/backend/env"
)

// NewRouter returns a new gin.Engine with everything set up.
func NewRouter() (*gin.Engine, error) {
	if env.MODE.Value() == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	if err := router.SetTrustedProxies(nil); err != nil {
		return nil, err
	}

	corsConfig := cors.Config{
		AllowOrigins:     []string{env.FRONTEND_URL.Value()},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Origin"},
		AllowHeaders:     []string{"Content-Length", "Content-Type", "Origin"},
		AllowCredentials: true,
	}
	if err := corsConfig.Validate(); err != nil {
		return nil, err
	}
	router.Use(cors.New(corsConfig))
	router.Use(middleware.ErrorMiddleware())

	if err := controllers.SetupWithRouter(router); err != nil {
		return nil, err
	}

	return router, nil
}

// ListenAndServe starts the HTTP server.
func ListenAndServe(router *gin.Engine) error {
	port := env.PORT.Value()
	if env.CERT_DIR.Value() == "" {
		log.Printf("HTTP server listening on port %s.", port)
		return router.Run(":" + port)
	}

	log.Printf("HTTPS server listening on port %s.", port)
	cert := filepath.Join(env.CERT_DIR.Value(), "cert.pem")
	key := filepath.Join(env.CERT_DIR.Value(), "privkey.pem")
	return router.RunTLS(":"+port, cert, key)
}
