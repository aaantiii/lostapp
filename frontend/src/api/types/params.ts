export type QueryParam = {
  q?: string
}

export type PaginationParams = {
  page: string
  pageSize: string
}

export type DiscordIdParam = {
  discordId?: string
}

export type PlayersParams = QueryParam & PaginationParams & DiscordIdParam
