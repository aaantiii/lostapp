import { lazy } from 'react'

export default {
  Index: lazy(() => import('./Index')),
  Error: lazy(() => import('./Error')),
  Account: lazy(() => import('./Account')),
  members: {
    View: lazy(() => import('./members/ViewMember')),
  },
  auth: {
    Login: lazy(() => import('./auth/Login')),
  },
}
