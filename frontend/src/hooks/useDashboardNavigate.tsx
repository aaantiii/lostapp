import { useCallback } from 'react'
import { useNavigate } from 'react-router-dom'
import { useAuth } from '@context/authContext'
import { AuthRole } from '@api/types/auth'

export default function useDashboardNavigate() {
  const navigate = useNavigate()
  const { userRole } = useAuth()

  const redirectFunc = useCallback(() => {
    switch (userRole) {
      case undefined:
        navigate('/auth/login')
        break
      case AuthRole.User:
        navigate('/apply')
        break
      case AuthRole.Member:
        navigate('/member')
        break
      case AuthRole.Leader:
      case AuthRole.Admin:
        navigate('/leader')
        break
    }
  }, [userRole])

  return redirectFunc
}
