package util

import "golang.org/x/text/message"

var printer = message.NewPrinter(message.MatchLanguage("de"))

func FormatNumber(n int) string {
	return printer.Sprint(n)
}
