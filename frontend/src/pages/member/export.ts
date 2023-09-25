import Index from './Index'
import View from './View'
import Find from './Find'
import Stats from './Stats'
import Leaderboard from './Leaderboard'
import OneVersusOne from './OneVersusOne'
import ClansIndex from './clans/Index'
import ClansView from './clans/View'
import ViewClanMember from './clans/ViewClanMember'

export default {
  Index,
  View,
  Find,
  Stats,
  Leaderboard,
  OneVersusOne,
  clans: {
    Index: ClansIndex,
    View: ClansView,
    members: {
      View: ViewClanMember,
    },
  },
} as const
