package postgres

import (
	"fmt"
	"strings"

	"gorm.io/gorm"

	"github.com/aaantiii/lostapp/backend/api/types"
)

func WithPagination(params types.PaginationParams) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if params.Page == 0 {
			return db // for optional pagination
		}
		offset := (params.Page - 1) * params.Limit
		return db.Offset(offset).Limit(params.Limit)
	}
}

func WithContains(value string, fields ...string) func(db *gorm.DB) *gorm.DB {
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

type Preloader struct {
	Field string
	Args  []any
}

func WithPreloading(preload ...Preloader) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(preload) == 0 {
			return db
		}
		for _, p := range preload {
			if p.Args == nil {
				db = db.Preload(p.Field)
				continue
			}
			db = db.Preload(p.Field, p.Args...)
		}
		return db
	}
}
