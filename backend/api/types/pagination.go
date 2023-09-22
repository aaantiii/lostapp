package types

type Pagination struct {
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
	TotalItems int `json:"totalItems"`
	TotalPages int `json:"totalPages"`
}

type PaginationParams struct {
	Page     int `form:"page" binding:"required,min=1"`
	PageSize int `form:"pageSize" binding:"required,min=1,max=50"`
}

type PaginatedResponse[T any] struct {
	Pagination Pagination `json:"pagination"`
	Items      []T        `json:"items"`
}

func NewPaginatedResponse[T any](data []T, params PaginationParams) *PaginatedResponse[T] {
	totalItems := len(data)
	totalPages := totalItems / params.PageSize
	if totalItems%params.PageSize > 0 {
		totalPages++
	}

	res := &PaginatedResponse[T]{
		Pagination: Pagination{
			Page:       params.Page,
			PageSize:   params.PageSize,
			TotalItems: totalItems,
			TotalPages: totalPages,
		},
	}

	start := (params.Page - 1) * params.PageSize
	if start >= totalItems {
		return res
	}

	end := start + params.PageSize
	if end > totalItems {
		end = totalItems
	}

	res.Items = data[start:end]
	return res
}
