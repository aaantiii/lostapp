export type Pagination = {
  page: number
  pageSize: number
  totalPages: number
  totalItems: number
  navigation: number[]
}

export type PaginatedResponse<T> = {
  items: T[]
  pagination: Pagination
}
