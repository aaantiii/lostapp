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
	ErrNotSignedIn      = NewApiError(http.StatusUnauthorized, "Du bist nicht angemeldet.")
	ErrUnknown          = NewApiError(http.StatusInternalServerError, "Während deiner Anfrage ist ein unerwarteter Fehler aufgetreten.")
	ErrNotMember        = NewApiError(http.StatusForbidden, "Nur Mitglieder der Lost Family sind berechtigt, auf die Web App zuzugreifen.")
	ErrNoResults        = NewApiError(http.StatusNotFound, "Es wurden keine Ergebnisse gefunden.")
	ErrAdminRequired    = NewApiError(http.StatusForbidden, "Um diese Aktion auszuführen, benötigst du Administratorrechte.")
	ErrNoPermission     = NewApiError(http.StatusForbidden, "Du hast keine Berechtigung, diese Aktion auszuführen.")
	ErrPageOutOfBounds  = NewApiError(http.StatusBadRequest, "Die angegebene Seite ist nicht vorhanden.")
	ErrSignOutFailed    = NewApiError(http.StatusInternalServerError, "Beim Abmelden ist ein Fehler aufgetreten.")
	ErrValidationFailed = NewApiError(http.StatusBadRequest, "Bitte überprüfe deine Eingaben.")
	ErrBadRequest       = NewApiError(http.StatusBadRequest, "Der Server konnte deine Anfrage nicht verstehen.")
	ErrInvalidStatName  = NewApiError(http.StatusBadRequest, "Die angegebene Statistik ist ungültig.")
)
