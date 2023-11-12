package util

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func BuildCustomID(cmdName, userID string) string {
	return fmt.Sprintf("%s*%s*%s", cmdName, userID, uuid.NewString())
}

func ParseCustomID(customID string) (cmdName, userID string) {
	split := strings.Split(customID, "*")
	return split[0], split[1]
}
