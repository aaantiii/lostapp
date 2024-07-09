import { Member } from './models'

export type LivePlayer = {
  discordId: string
  clanMembers?: Member[]
  league: League
  builderBaseLeague: BuilderBaseLeague
  clans: PlayerClan
  attackWins: number
  defenseWins: number
  townHallLevel: number
  tag: string
  name: string
  expLevel: number
  trophies: number
  bestTrophies: number
  donations: number
  donationsReceived: number
  builderHallLevel: number
  builderBaseTrophies: number
  bestBuilderBaseTrophies: number
  warStars: number
  clanCapitalContribution: number
}

export type PlayerClan = {
  tag: string
  name: string
  role: ClanMemberRole
}

export type League = {
  id: number
  name: string
  iconUrls: IconURLs
}

export type IconURLs = {
  tiny: string
  small: string
  medium: string
}

export type BuilderBaseLeague = {
  id: number
  name: string
}

export type LiveClan = {
  memberList: LiveClanMember[]
  warTies?: number
  description?: string
  warLosses?: number
  badgeUrls: IconURLs
  tag: string
  name: string
  level: number
  warWins: number
  members: number
  isWarLogPublic: boolean
}

export enum ClanMemberRole {
  Leader = 'leader',
  CoLeader = 'coLeader',
  Elder = 'admin',
  Member = 'member',
  NotMember = 'notMember',
}

export const ClanMemberRoleTranslated: ReadonlyMap<ClanMemberRole, string> = new Map([
  [ClanMemberRole.Leader, 'Anführer'],
  [ClanMemberRole.CoLeader, 'Vize Anführer'],
  [ClanMemberRole.Elder, 'Ältester'],
  [ClanMemberRole.Member, 'Mitglied'],
])

export type LiveClanMember = {
  tag: string
  name: string
  role: ClanMemberRole
  expLevel: number
}
