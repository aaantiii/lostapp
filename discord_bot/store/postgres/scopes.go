package postgres

import (
	"database/sql"
	"fmt"

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
		if value == "" {
			return db
		}

		query := "("
		for i, field := range fields {
			query += fmt.Sprintf("POSITION(LOWER(@value) in LOWER(%s)) > 0", field)
			if len(fields)-1 > i {
				query += " OR "
			}
		}
		query += ")"

		return db.Where(query, sql.Named("value", value))
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
