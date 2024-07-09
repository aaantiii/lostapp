package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aaantiii/lostapp/backend/api/types"
	"github.com/aaantiii/lostapp/backend/api/utils"
)

const (
	minPage     = 1
	minPageSize = 3
	maxPageSize = 50
)

func PaginationMiddleware(optional bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pagination types.PaginationParams
		if err := c.ShouldBindQuery(&pagination); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid pagination: " + err.Error()})
			return
		}
		if optional && pagination.Page == 0 {
			c.Set(utils.PaginationKey, types.PaginationParams{})
			return
		}

		if pagination.Page < minPage {
			pagination.Page = minPage
		}

		switch {
		case pagination.Limit < minPageSize:
			pagination.Limit = minPageSize
		case pagination.Limit > maxPageSize:
			pagination.Limit = maxPageSize
		}

		c.Set(utils.PaginationKey, pagination)
	}
}
