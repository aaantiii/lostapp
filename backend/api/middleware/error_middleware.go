package middleware

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/aaantiii/lostapp/backend/api/types"
	"github.com/aaantiii/lostapp/backend/api/util"
	"github.com/aaantiii/lostapp/backend/env"
)

var errUnknown = types.NewApiError(http.StatusInternalServerError, "Während deiner Anfrage ist ein unerwarteter Fehler aufgetreten.")

// NewErrorMiddleware returns a gin.HandlerFunc that handles types.ApiError, if present in gin.Context.
func NewErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		v, ok := c.Get(util.ErrorKey)
		if !ok {
			return // no error
		}

		err, ok := v.(error)
		if !ok {
			c.JSON(errUnknown.Code, errUnknown)
			return
		}

		if env.MODE.Value() == "DEV" {
			log.Println(err.Error())
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(types.ErrNoResults.Code, types.ErrNoResults)
			return
		}

		var apiErr *types.ApiError
		if !errors.As(err, &apiErr) {
			c.JSON(errUnknown.Code, errUnknown)
			return
		}

		if apiErr.Code == http.StatusNoContent {
			c.Status(apiErr.Code)
			return
		}

		c.JSON(apiErr.Code, apiErr)
	}
}
