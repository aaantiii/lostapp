package postgres

import (
	"fmt"
	"strings"

	"gorm.io/gorm"

	"bot/types"
)

func WithPaging(params types.PaginationParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (params.Page - 1) * params.PageSize
		return db.Offset(offset).Limit(params.PageSize)
	}
}

// WithSearchQuery returns a scope that filters records based on the given search query. It looks for find in the given dbColumns.
func WithSearchQuery(find string, dbColumns ...string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if find == "" || len(dbColumns) == 0 {
			return db
		}

		textCols := make([]string, len(dbColumns))
		for i, col := range dbColumns {
			textCols[i] = fmt.Sprintf("%s::text", col)
		}

		query := fmt.Sprintf("POSITION(? in LOWER(%s)) > 0", strings.Join(textCols, " || ' ' || "))
		return db.Where(query, strings.ToLower(find))
	}
}

func WithLimit(limit int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if limit <= 0 {
			return db
		}
		return db.Limit(limit)
	}
}

func WithPreloading(fields ...string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(fields) == 0 {
			return db
		}
		for _, field := range fields {
			db = db.Preload(field)
		}
		return db
	}
}
