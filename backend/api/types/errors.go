package types

import "net/http"

type ApiError struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

// implement error interface
func (e ApiError) Error() string {
	return e.Message
}

func NewApiError(code int, message string) *ApiError {
	return &ApiError{
		Message: message,
		Code:    code,
	}
}

var (
	ErrNotMember     = NewApiError(http.StatusForbidden, "Nur Mitglieder der Lost Family sind berechtigt, auf die Web App zuzugreifen.")
	ErrNoResults     = NewApiError(http.StatusNotFound, "Es wurden keine Ergebnisse gefunden.")
	ErrAdminRequired = NewApiError(http.StatusForbidden, "Um diese Aktion auszuführen, benötigst du Administratorrechte.")
)
