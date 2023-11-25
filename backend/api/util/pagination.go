package util

import (
	"github.com/gin-gonic/gin"

	"backend/api/types"
)

const PaginationKey = "pagination"

func PaginationFromContext(c *gin.Context) *types.PaginationParams {
	return c.MustGet(PaginationKey).(*types.PaginationParams)
}
