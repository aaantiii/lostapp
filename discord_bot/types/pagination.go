package types

import (
	"math"
)

// PaginationParams are the parameters used to fetch paginated data.
type PaginationParams struct {
	Page     int `form:"page"`
	PageSize int `form:"pageSize"`
}

// Pagination is the pagination metadata returned by the API.
type Pagination struct {
	Page       int   `json:"page"`
	PageSize   int   `json:"pageSize"`
	TotalItems int64 `json:"totalItems"`
	TotalPages int   `json:"totalPages"`
	Navigation []int `json:"navigation,omitempty"`
}

// PaginatedResponse is the response returned by the API when fetching paginated data.
type PaginatedResponse[T any] struct {
	Pagination Pagination `json:"pagination"`
	Items      []T        `json:"items,omitempty"`
}

// NewPaginatedResponse creates a new PaginatedResponse.
func NewPaginatedResponse[T any](items []T, params PaginationParams, count int64) *PaginatedResponse[T] {
	totalPages := int(math.Ceil(float64(count) / float64(params.PageSize)))
	pagination := Pagination{
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalItems: count,
		TotalPages: totalPages,
	}
	pagination.setNavigation()

	return &PaginatedResponse[T]{
		Items:      items,
		Pagination: pagination,
	}
}

const pageNavigationSize = 5

func (p *Pagination) setNavigation() {
	nav := make([]int, 0, pageNavigationSize)
	if p.TotalPages <= pageNavigationSize+1 {
		for i := 1; i <= p.TotalPages; i++ {
			nav = append(nav, i)
		}
	} else {
		afterFirst := int(math.Ceil(float64(p.TotalPages) / 4))
		mid := int(math.Ceil(float64(p.TotalPages) / 2))
		beforeLast := p.TotalPages - afterFirst + 1
		nav = append(nav, 1, afterFirst, mid, beforeLast, p.TotalPages)
	}

	p.Navigation = nav
}
