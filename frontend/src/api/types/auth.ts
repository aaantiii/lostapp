export interface Session {
  discordUser?: DiscordUser
  userRole?: AuthRole
}

export interface DiscordUser {
  id: string
  username: string
  avatarUrl: string
}

export enum AuthRole {
  User = 1,
  Member = 2,
  Leader = 3,
  Admin = 4,
}
