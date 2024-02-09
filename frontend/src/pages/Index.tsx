import LoadingScreen from '@components/LoadingScreen'
import { useAuth } from '@context/authContext'
import { useEffect } from 'react'
import { useNavigate } from 'react-router-dom'

export default function Index() {
  const { user } = useAuth()
  const navigate = useNavigate()

  useEffect(() => {
    if (!user) navigate('/auth/login')
    else navigate('/members/@me')
  }, [user])

  return (
    <main>
      <LoadingScreen />
    </main>
  )
}
