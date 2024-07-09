package types

import (
	"math"
	"slices"
	"sort"
)

// PaginationParams are the parameters used to fetch paginated data.
type PaginationParams struct {
	Page        int `form:"page" binding:"omitempty,min=1"`
	Limit       int `form:"limit"`
	ExtraOffset int `form:"extraOffset"` // ExtraOffset is used to offset the pagination by a given amount. E.g. by leaderboard to skip top 3.
}

// Pagination is the pagination metadata returned by the API.
type Pagination struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
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
	totalPages := int(math.Ceil(float64(count) / float64(params.Limit)))
	pagination := Pagination{
		Page:       params.Page,
		Limit:      params.Limit,
		TotalItems: count,
		TotalPages: totalPages,
	}
	pagination.setNavigation()

	return &PaginatedResponse[T]{
		Items:      items,
		Pagination: pagination,
	}
}

const pageNavigationSize = 7

func (p *Pagination) setNavigation() {
	nav := make([]int, 0, pageNavigationSize)
	if p.TotalPages <= pageNavigationSize {
		for i := 1; i <= p.TotalPages; i++ {
			nav = append(nav, i)
		}
	} else {
		quarter := int(math.Ceil(float64(p.TotalPages) / 4))
		mid := int(math.Ceil(float64(p.TotalPages) / 2))
		threeQuarter := p.TotalPages - quarter + 1
		nav = append(nav, 1, quarter, mid, threeQuarter, p.TotalPages)
	}

	if !slices.Contains(nav, p.Page) {
		nav = append(nav, p.Page)
	}

	if oneBelow := p.Page - 1; p.Page > 1 && !slices.Contains(nav, oneBelow) {
		nav = append(nav, oneBelow)
	}

	if oneAbove := p.Page + 1; p.Page < p.TotalPages && !slices.Contains(nav, oneAbove) {
		nav = append(nav, oneAbove)
	}

	sort.Slice(nav, func(i, j int) bool {
		return nav[i] < nav[j]
	})

	p.Navigation = nav
}
