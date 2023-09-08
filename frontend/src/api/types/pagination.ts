export interface Pagination {
  page: number
  pageSize: number
  totalItems: number
  totalPages: number
}

export interface PaginatedResponse<T> {
  items: T[]
  pagination: Pagination
}
