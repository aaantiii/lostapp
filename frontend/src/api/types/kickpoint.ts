import { ClanMemberRole } from './clan'

export interface Kickpoint {
  id: number
  amount: number
  date: string
  reason: string
  description: string
}

export interface ClanMemberKickpoints {
  tag: string
  name: string
  role: ClanMemberRole
  amount: number
}
