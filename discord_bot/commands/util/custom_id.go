package util

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

func BuildCustomID(cmdName, userID, otherID string) string {
	return fmt.Sprintf("%s$%s$%s$%s", cmdName, userID, otherID, uuid.NewString())
}

func ParseCustomID(customID string) (cmdName, userID, otherID string) {
	split := strings.Split(customID, "$")
	if len(split) != 4 {
		return "", "", ""
	}

	return split[0], split[1], split[2]
}
