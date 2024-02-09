package utils

import (
	"time"

	"github.com/jinzhu/now"
)

func KickpointMinDate(expiresAfterDays int) time.Time {
	return now.BeginningOfDay().AddDate(0, 0, -(expiresAfterDays - 1))
}
