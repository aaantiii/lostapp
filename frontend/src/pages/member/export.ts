import Index from './Index'
import View from './View'
import Find from './Find'
import Stats from './Stats'
import Leaderboard from './Leaderboard'
import OneVersusOne from './OneVersusOne'
import Kickpoints from './Kickpoints'
import ClansIndex from './clans/Index'
import ClansView from './clans/View'

export default {
  Index,
  View,
  Find,
  Stats,
  Leaderboard,
  OneVersusOne,
  Kickpoints,
  clans: {
    Index: ClansIndex,
    View: ClansView,
  },
} as const
