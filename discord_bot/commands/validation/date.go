package validation

import "time"

func ValidateEventDates(startsAt, endsAt time.Time) (string, bool) {
	if startsAt.Before(time.Now()) {
		return "Das Startdatum muss in der Zukunft liegen.", false
	}
	if !endsAt.After(startsAt) {
		return "Das Enddatum muss nach dem Startdatum liegen.", false
	}
	return "", true
}
