package middleware

import "github.com/aaantiii/lostapp/backend/api/services"

var authService services.IAuthService

func InjectServices(auth services.IAuthService) {
	authService = auth
}
