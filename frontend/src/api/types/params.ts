export type QueryParam = {
  q?: string
}

export type PaginationParams = {
  page: string
  limit: string
}

export type DiscordIdParam = {
  discordId?: string
}

export type PlayersParams = QueryParam & PaginationParams & DiscordIdParam & { isMember?: 'true' | 'false' }

export type ClansParams = QueryParam & PaginationParams
