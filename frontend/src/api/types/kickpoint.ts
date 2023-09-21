import { ClanMemberRole } from './clan'
import { CreatedByUser, DiscordUser, UpdatedByUser } from './user'

export interface Kickpoint extends CreatedByUser, UpdatedByUser {
  id: number
  amount: number
  date: string
  description: string
}

export type CreateKickpoint = Omit<Kickpoint, 'id'>

export type UpdateKickpoint = Partial<CreateKickpoint>

export interface ClanMemberKickpoints {
  tag: string
  name: string
  role: ClanMemberRole
  amount: number
}
