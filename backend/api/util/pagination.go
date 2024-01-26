package util

import (
	"github.com/gin-gonic/gin"

	"github.com/aaantiii/lostapp/backend/api/types"
)

const PaginationKey = "pagination"

func PaginationFromContext(c *gin.Context) *types.PaginationParams {
	return c.MustGet(PaginationKey).(*types.PaginationParams)
}
