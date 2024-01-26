package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/aaantiii/lostapp/backend/api/types"
	"github.com/aaantiii/lostapp/backend/api/util"
)

const (
	minPage     = 1
	minPageSize = 3
	maxPageSize = 50
)

func NewPaginationMiddleware(optional bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var pagination types.PaginationParams
		if err := c.ShouldBindQuery(&pagination); err != nil {
			if optional {
				return
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid pagination: " + err.Error()})
			return
		}

		if pagination.Page < minPage {
			pagination.Page = minPage
		}

		switch {
		case pagination.PageSize < minPageSize:
			pagination.PageSize = minPageSize
		case pagination.PageSize > maxPageSize:
			pagination.PageSize = maxPageSize
		}

		c.Set(util.PaginationKey, pagination)
	}
}
