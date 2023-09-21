export interface DiscordUser {
  id: string
  name: string
  avatarUrl: string
}

export interface CreatedByUser {
  createdAt: string
  createdByUser: DiscordUser
}

export interface UpdatedByUser {
  updatedAt?: string
  updatedByUser?: DiscordUser
}
