package middleware

import (
	"errors"
	"log/slog"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/aaantiii/lostapp/backend/api/types"
)

const ErrorKey = "error"

// ErrorMiddleware returns a gin.HandlerFunc that handles types.ApiError, if present in gin.Context.
func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		v, ok := c.Get(ErrorKey)
		if !ok {
			return // no error
		}

		err, ok := v.(error)
		if !ok {
			c.JSON(types.ErrUnknown.Code, types.ErrUnknown)
			return
		}

		slog.Debug("Error from ErrorMiddleware.", slog.Any("err", err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(types.ErrNoResults.Code, types.ErrNoResults)
			return
		}

		var apiErr *types.ApiError
		if !errors.As(err, &apiErr) {
			c.JSON(types.ErrUnknown.Code, types.ErrUnknown)
			return
		}

		c.JSON(apiErr.Code, apiErr)
	}
}
