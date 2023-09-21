package middleware

import (
	"backend/api/services"
)

var (
	authService  services.IAuthService
	clansService services.IClansService
)

func InjectServices(auth services.IAuthService, clans services.IClansService) {
	authService = auth
	clansService = clans
}
