import { ClanMemberRole } from './clan'

export interface Kickpoint {
  id: number
  amount: number
  date: string
  reason: string
}

export type CreateKickpoint = Omit<Kickpoint, 'id'>

export type UpdateKickpoint = Partial<CreateKickpoint>

export interface ClanMemberKickpoints {
  tag: string
  name: string
  role: ClanMemberRole
  amount: number
}
