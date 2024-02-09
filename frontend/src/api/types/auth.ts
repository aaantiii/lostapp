export type User = {
  id: string
  name: string
  avatarUrl: string
  memberOf?: string[]
  coLeaderOf?: string[]
  leaderOf?: string[]
  isAdmin: boolean
}

export type Session = {
  user?: User
}

export enum AuthRole {
  AnyMember = 'anyMember',
  ClanMember = 'clanMember',
  AnyCoLeader = 'anyCoLeader',
  ClanCoLeader = 'clanCoLeader',
  AnyLeader = 'anyLeader',
  ClanLeader = 'clanLeader',
  Admin = '~~~admin~~~',
}
