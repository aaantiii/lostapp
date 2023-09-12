import { Outlet } from 'react-router-dom'
import { AuthRole } from '@api/types/auth'
import ProtectedRoute from './ProtectedRoute'
import useNotImplemented from '@hooks/useNotImplemented'

export default function UserRoute() {
  useNotImplemented()
  return (
    <ProtectedRoute requiredRole={AuthRole.User}>
      <Outlet />
    </ProtectedRoute>
  )
}
