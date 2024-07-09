import { AuthRole } from '@api/types/auth'
import { useAuth } from '@context/authContext'
import { ReactNode } from 'react'

type RoleRenderProps = {
  role: AuthRole
  clanTag?: string
  children: ReactNode
}

// Renders its children if the user is authorized, otherwise renders nothing.
export default function RoleRender({ role, clanTag, children }: RoleRenderProps) {
  const { getAuthRoles, user } = useAuth()
  if (user && user.isAdmin) return children

  const userRoles = getAuthRoles(clanTag)
  if (userRoles && userRoles.includes(role)) return children

  return null
}
