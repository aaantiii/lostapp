import { ClanMemberRole } from './clan'

export interface Player {
  discordId: string
  clans: PlayerClan[]
  name: string
  tag: string
  warPreference: string
  trophies: number
  expLevel: number
  townHallLevel: number
  attackWins: number
  defenseWins: number
}

export interface PlayerClan {
  tag: string
  name: string
  role: ClanMemberRole
}
