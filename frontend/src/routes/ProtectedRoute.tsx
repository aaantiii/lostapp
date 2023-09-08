import { useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useAuth } from '../context/authContext'
import { AuthRole } from '../api/types/auth'
import LoadingScreen from '../components/LoadingScreen'
import Spacer from '../components/Spacer'
import Content from '../components/Content'
import useScreenSize, { ScreenSize } from '../hooks/useScreenSize'

interface ProtectedRouteProps {
  requiredRole: AuthRole
  children: JSX.Element[] | JSX.Element
}

export default function ProtectedRoute({ requiredRole, children }: ProtectedRouteProps) {
  const navigate = useNavigate()
  const screenSize = useScreenSize()
  const { userRole } = useAuth()
  const [hasPermission, setHasPermission] = useState(false)

  useEffect(() => {
    if (userRole === undefined) return navigate('/auth/login')
    if (userRole < requiredRole) return navigate('/error/403')

    setHasPermission(true)
  }, [userRole])

  return hasPermission ? (
    <Content>
      <Spacer size={screenSize <= ScreenSize.TabletPortrait ? 'medium' : 'large'} />
      <div className="ProtectedRoute">{children}</div>
    </Content>
  ) : (
    <LoadingScreen />
  )
}
