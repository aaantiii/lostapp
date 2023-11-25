package util

import (
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
	return d.Round(time.Second).String()
}
