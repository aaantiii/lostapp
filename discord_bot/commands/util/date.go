package util

import (
	"fmt"
	"time"
)

func FormatMonthYear(month, year int) string {
	return fmt.Sprintf("%02d-%d", month, year)
}

func FormatDateString(t time.Time) string {
	return t.Format("02.01.2006")
}
