package postgres

import (
	"gorm.io/gorm"

	"backend/api/types"
)

// Paginate applies pagination to the given query.
// Source: https://gorm.io/docs/scopes.html#Pagination (modified)
func Paginate(params *types.PaginationParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (params.Page - 1) * params.PageSize
		return db.Offset(offset).Limit(params.PageSize)
	}
}
