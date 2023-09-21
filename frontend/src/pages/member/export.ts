import Index from './Index'
import Find from './Find'
import Stats from './Stats'
import Leaderboard from './Leaderboard'
import OneVersusOne from './OneVersusOne'
import ClansIndex from './clans/Index'
import ClansView from './clans/View'
import ViewMember from './clans/ViewClanMember'

export default {
  Index,
  Find,
  Stats,
  Leaderboard,
  OneVersusOne,
  clans: {
    Index: ClansIndex,
    View: ClansView,
    members: {
      View: ViewMember,
    },
  },
} as const
