import { AxiosError } from 'axios'
import { User } from './auth'
import { ClanMemberRole } from './coc'
import { ModifiedBy } from './timestamp'

export type ApiError = AxiosError<{
  message: string
  code: number
}>

export type Member = {
  playerTag: string
  clanTag: string
  clanRole: ClanMemberRole
  addedByUser?: User
  player?: Player
  clan?: Clan
}

export type Clan = {
  tag: string
  name: string
  settings?: ClanSettings
  clanMembers: Member[]
}

export type Kickpoint = ModifiedBy & {
  id: number
  amount: number
  date: string
  description: string
}

export type CreateKickpointPayload = Omit<Kickpoint, 'id'>

export type UpdateKickpointPayload = CreateKickpointPayload

export type Player = {
  cocTag: string
  name: string
  discordId: string
  members: Member[]
}

export type ClanSettings = ModifiedBy & {
  clanTag: string
  minSeasonWins: number
  maxKickpoints: number
  kickpointsExpireAfterDays: number
}

export type UpdateClanSettingsPayload = Omit<ClanSettings, 'clanTag' | keyof ModifiedBy>
