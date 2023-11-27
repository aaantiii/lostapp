package postgres

import (
	"fmt"
	"strings"

	"gorm.io/gorm"

	"bot/types"
)

// ScopePaginate applies pagination to the given query.
//
// Example: db.Scopes(ScopePaginate(&types.PaginationParams{Page: 1, PageSize: 20})).Find(&users)
func ScopePaginate(params types.PaginationParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (params.Page - 1) * params.PageSize
		return db.Offset(offset).Limit(params.PageSize)
	}
}

func ScopeContains(value string, fields ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if value == "" || len(fields) == 0 {
			return db
		}

		for i, field := range fields {
			fields[i] = fmt.Sprintf("%s::text", field)
		}

		concatenatedFields := strings.Join(fields, " || ' ' || ")
		query := fmt.Sprintf("POSITION(? in LOWER(%s)) > 0", concatenatedFields)
		return db.Where(query, strings.ToLower(value))
	}
}

func ScopeLimit(limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if limit <= 0 {
			return db
		}

		return db.Limit(limit)
	}
}
