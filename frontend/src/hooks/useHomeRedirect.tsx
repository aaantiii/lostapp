import { useAuth } from '@context/authContext'
import { useEffect } from 'react'
import { useLocation, useNavigate } from 'react-router-dom'

// useHomeRedirect is a hook that redirects the user to their corresponding home page.
// user not logged in: redirect to /auth/login
// user logged in and is admin: redirect to /admin/dashboard
// user logged in and is not admin: redirect to /user/leaderboard
export default function useHomeRedirect() {
  const navigate = useNavigate()
  const { pathname } = useLocation()
  const { user } = useAuth()

  useEffect(() => {
    if (!user && pathname !== '/auth/login') return navigate('/auth/login', { replace: true })
    if (!user) return
    if (user.isAdmin) navigate('/admin/dashboard')
    else navigate('/user/leaderboard')
  })
}
