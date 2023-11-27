package util

import (
	"fmt"
	"strings"
	"time"
)

const (
	dateFormat     = "02.01.2006"
	dateTimeFormat = "02.01.2006, 15:04"
)

func FormatDate(t time.Time) string {
	return t.Format(dateFormat)
}

func FormatDateTime(t time.Time) string {
	return t.Format(dateTimeFormat)
}

func ParseDateString(s string) (time.Time, error) {
	date, err := time.Parse(dateFormat, s)
	if err != nil {
		return date, err
	}

	return TruncateToDay(date), nil
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
