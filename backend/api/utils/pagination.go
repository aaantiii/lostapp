package utils

import (
	"github.com/gin-gonic/gin"

	"github.com/aaantiii/lostapp/backend/api/types"
)

const PaginationKey = "pagination"

func PaginationFromContext(c *gin.Context) types.PaginationParams {
	return c.MustGet(PaginationKey).(types.PaginationParams)
}

func ValidatePagination(params types.PaginationParams, count int64) error {
	if count <= 0 {
		return types.ErrNoResults
	}
	if int64(params.Page*params.Limit-params.Limit) > count {
		return types.ErrPageOutOfBounds
	}

	return nil
}

func PaginateSlice[T any](params types.PaginationParams, slice []T) []T {
	start := params.Page*params.Limit - params.Limit
	end := params.Page * params.Limit
	if end > len(slice) {
		end = len(slice)
	}

	return slice[start:end]
}
