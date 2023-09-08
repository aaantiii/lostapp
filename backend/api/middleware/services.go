package middleware

import "backend/api/services"

var (
	authService    services.IAuthService
	clansService   services.IClansService
	playersService services.IPlayersService
)

func InjectServices(auth services.IAuthService, clans services.IClansService, players services.IPlayersService) {
	authService = auth
	clansService = clans
	playersService = players
}
