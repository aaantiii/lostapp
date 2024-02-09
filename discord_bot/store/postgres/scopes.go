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

func WithSearchQuery(value string, fields ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if value == "" || len(fields) == 0 {
			return db
		}

		strFields := make([]string, len(fields))
		for i, field := range fields {
			strFields[i] = fmt.Sprintf("%s::text", field)
		}

		find := strings.Join(strFields, " || ' ' || ")
		query := fmt.Sprintf("POSITION(? in LOWER(%s)) > 0", find)
		return db.Where(query, strings.ToLower(value))
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
