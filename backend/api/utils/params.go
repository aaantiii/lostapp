package utils

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// TagFromParams returns the a tag defined as paramName from gin.Context and adds a "#" to the beginning.
// If the tag is empty, an empty string is returned.
func TagFromParams(c *gin.Context, paramName string) (string, error) {
	if tag := strings.ReplaceAll(c.Param(paramName), " ", ""); tag != "" {
		return "#" + tag, nil
	}
	return "", errors.New("tag is empty")
}

func UintFromParams(c *gin.Context, paramName string) (uint, error) {
	if param := c.Param(paramName); param != "" {
		v, err := strconv.ParseUint(param, 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(v), nil
	}
	return 0, errors.New("param is empty")
}
