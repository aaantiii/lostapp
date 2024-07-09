import { ClanMemberRole } from './coc'

export type ClanMemberKickpoints = {
  tag: string
  name: string
  role: ClanMemberRole
  amount: number
}

export type ComparableStatistic = {
  id: number
  name: string
  displayName: string
}

export type PlayerStatistic = {
  name: string
  tag: string
  clanName: string
  value: number
  placement: number
}
