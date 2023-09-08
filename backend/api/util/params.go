package util

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
)

// TagFromQuery returns the tag from the query parameter "tag" and adds a "#" to the beginning.
// If the tag is empty, an empty string is returned.
func TagFromQuery(c *gin.Context, paramName string) (string, error) {
	if tag := strings.Trim(c.Param(paramName), " "); tag != "" {
		return "#" + tag, nil
	}
	return "", errors.New("tag is empty")
}
