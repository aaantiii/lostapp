package util

import (
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/now"

	"bot/store/postgres/models"
)

const (
	dateFormat        = "02.01.2006"
	dateTimeFormat    = "02.01.2006, 15:04"
	dateParsingFormat = "2.1.2006"
)

func FormatDate(t time.Time) string {
	return t.Format(dateFormat)
}

func FormatDateTime(t time.Time) string {
	return t.Format(dateTimeFormat)
}

func ParseDateString(s string) (time.Time, error) {
	date, err := time.Parse(dateParsingFormat, s)
	if err != nil {
		return date, err
	}

	return TruncateToDay(date), nil
}

func FormatFromAt(from *models.User, at time.Time) string {
	msg := ""
	if from != nil {
		msg += fmt.Sprintf("von %s ", from.Name)
	}
	if !at.IsZero() {
		msg += fmt.Sprintf("am %s", FormatDateTime(at))
	}

	return strings.Trim(msg, " ")
}

func TruncateToDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.Local)
}

func FormatDuration(d time.Duration) string {
	days := int(d.Hours() / 24)
	hours := int(d.Hours()) % 24
	minutes := int(d.Minutes()) % 60
	seconds := int(d.Seconds()) % 60

	var res string
	if days > 0 {
		res += fmt.Sprintf("%dd ", days)
	}
	if hours > 0 {
		res += fmt.Sprintf("%dh ", hours)
	}
	if minutes > 0 {
		res += fmt.Sprintf("%dm ", minutes)
	}
	if seconds > 0 {
		res += fmt.Sprintf("%ds ", seconds)
	}

	return strings.Trim(res, " ")
}

func KickpointMinDate(expiresAfterDays int) time.Time {
	return now.BeginningOfDay().AddDate(0, 0, -(expiresAfterDays - 1))
}
