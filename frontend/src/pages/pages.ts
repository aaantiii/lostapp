import { lazy } from 'react'

export default {
  Dashboard: lazy(() => import('./Dashboard')),
  Error: lazy(() => import('./Error')),
  Account: lazy(() => import('./Account')),
  Leaderboard: lazy(() => import('./Leaderboard')),
  clans: {
    Index: lazy(() => import('./clans/Index')),
    ByTag: lazy(() => import('./clans/ByTag')),
    members: {
      Index: lazy(() => import('./clans/members/Index')),
      ByTag: lazy(() => import('./clans/members/ByTag')),
      Kickpoints: lazy(() => import('./clans/members/kickpoints/Index')),
    },
  },
  auth: {
    Login: lazy(() => import('./auth/Login')),
  },
}
