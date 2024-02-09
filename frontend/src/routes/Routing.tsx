import { Navigate, Route, Routes } from 'react-router-dom'
import pages from '@pages/pages'
import AuthorizedRoute from './AuthorizedRoute'
import { AuthRole } from '@api/types/auth'

export default function Routing() {
  return (
    <Routes>
      <Route index element={<pages.Index />} />
      <Route path="error/:code" element={<pages.Error />} />
      <Route path="" element={<AuthorizedRoute requiredRole={AuthRole.AnyMember} />}>
        <Route path="account" element={<pages.Account />} />
      </Route>

      <Route path="auth">
        <Route path="login" element={<pages.auth.Login />} />
      </Route>

      <Route path="members" element={<AuthorizedRoute requiredRole={AuthRole.AnyMember} />}>
        <Route index element={<Navigate to="/members/@me" replace />} />
        <Route path=":discordId" element={<pages.members.View />} />
      </Route>

      <Route path="*" element={<Navigate to="/error/404" />} />
    </Routes>
  )
}
