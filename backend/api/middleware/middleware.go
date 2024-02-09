package middleware

import "github.com/aaantiii/lostapp/backend/api/services"

var authService services.IAuthService

func UseAuthService(auth services.IAuthService) {
	authService = auth
}
