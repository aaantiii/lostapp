import LeaderIndex from './Index'
import Notify from './Notify'
import ClanSettings from './clans/ClanSettings'
import ClanMemberIndex from './clans/members/Index'
import ManageClanMembers from './clans/members/Manage'
import MemberKickpoints from './clans/members/Kickpoints'

export default {
  Index: LeaderIndex,
  Notify,
  clans: {
    ClanSettings,
    members: {
      Index: ClanMemberIndex,
      Manage: ManageClanMembers,
      Kickpoints: MemberKickpoints,
    },
  },
} as const
