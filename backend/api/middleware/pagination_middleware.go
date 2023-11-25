package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"backend/api/types"
	"backend/api/util"
)

func NewPaginationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var pagination types.PaginationParams
		if err := c.Bind(&pagination); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid pagination"})
			return
		}

		defer c.Next()
		c.Set(util.PaginationKey, pagination)
	}
}
