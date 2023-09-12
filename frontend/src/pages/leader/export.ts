import LeaderIndex from './Index'
import Notify from './Notify'
import ClanSettings from './clans/ClanSettings'
import ClanMemberIndex from './clans/members/Index'
import MemberKickpoints from './clans/members/kickpoints/Index'

export default {
  Index: LeaderIndex,
  Notify,
  clans: {
    ClanSettings,
    members: {
      Index: ClanMemberIndex,
      Kickpoints: MemberKickpoints,
    },
  },
} as const
