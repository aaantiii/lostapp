import { useCallback } from 'react'
import { useNavigate } from 'react-router-dom'
import { useAuth } from '../context/authContext'
import { AuthRole } from '../api/types/auth'

export default function useDashboardNavigate() {
  const navigate = useNavigate()
  const { userRole } = useAuth()

  const redirectFunc = useCallback(() => {
    switch (userRole) {
      case undefined:
        navigate('/auth/login')
        break
      case AuthRole.User:
        navigate('/user')
        break
      case AuthRole.Member:
        navigate('/member')
        break
      case AuthRole.Leader:
      case AuthRole.Admin: // admin not implemented
        navigate('/leader')
        break
    }
  }, [userRole])

  return redirectFunc
}
