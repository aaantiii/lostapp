export interface PaginationParams {
  page?: number
  pageSize?: number
}

export interface PlayersParams extends PaginationParams {
  name?: string
  tag?: string
  clanName?: string
  clanTag?: string
  discordID?: string
}
