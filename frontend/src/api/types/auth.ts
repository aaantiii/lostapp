import { DiscordUser } from './user'

export interface Session {
  discordUser?: DiscordUser
  userRole?: AuthRole
}

export enum AuthRole {
  User = 1,
  Member = 2,
  Leader = 3,
  Admin = 4,
}
