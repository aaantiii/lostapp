import { Navigate, Route, Routes } from 'react-router-dom'
import pages from '@pages/pages'
import AuthorizedRoute from './AuthorizedRoute'
import { AuthRole } from '@api/types/auth'

export default function Routing() {
  return (
    <Routes>
      <Route path="error/:code" element={<pages.Error />} />

      <Route path="" element={<AuthorizedRoute requiredRole={AuthRole.AnyMember} />}>
        <Route index element={<Navigate to="/dashboard" replace />} />
        <Route path="dashboard" element={<pages.Dashboard />} />
        <Route path="account" element={<pages.Account />} />
        <Route path="leaderboard" element={<pages.Leaderboard />} />
      </Route>

      <Route path="auth">
        <Route path="login" element={<pages.auth.Login />} />
      </Route>

      <Route path="clans" element={<AuthorizedRoute requiredRole={AuthRole.AnyMember} />}>
        <Route index element={<pages.clans.Index />} />
        <Route path=":clanTag">
          <Route index element={<pages.clans.ByTag />} />
          <Route path="members" element={<pages.clans.members.Index />} />
          <Route path="members/:memberTag">
            <Route index element={<pages.clans.members.ByTag />} />
            <Route path="kickpoints" element={<pages.clans.members.Kickpoints />} />
          </Route>
        </Route>
      </Route>

      <Route path="*" element={<Navigate to="/error/404" />} />
    </Routes>
  )
}
