import { AuthRole } from '@api/types/auth'
import { useAuth } from '@context/authContext'
import { Outlet, useParams } from 'react-router-dom'

type AuthorizedRouteProps = {
  requiredRole: AuthRole
}

// a component that renders its children if the user is authorized, otherwise redirects to the error/403 page.
export default function AuthorizedRoute({ requiredRole }: AuthorizedRouteProps) {
  const { clanTag } = useParams()
  const { getAuthRoles } = useAuth()
  return getAuthRoles(clanTag)?.includes(requiredRole) ? <Outlet /> : <p>You are not authorized to see this page.</p>
}
