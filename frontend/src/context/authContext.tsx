import { createContext, useContext, useEffect } from 'react'
import { useNavigate } from 'react-router-dom'
import { useQuery, useQueryClient } from '@tanstack/react-query'
import { AxiosError, HttpStatusCode } from 'axios'
import { AuthRole, Session } from '@api/types/auth'
import routes from '@api/routes'
import client from '@api/client'
import { useMessages } from './messagesContext'
import LoadingScreen from '@components/LoadingScreen'
import { urlDecodeTag } from '@/utils/cocFormatter'

type AuthContext = Session & {
  refreshSession: () => void
  logout: () => Promise<void>
  getAuthRoles: (clanTag: string | undefined) => AuthRole[] | null
}

const authContext = createContext({} as AuthContext)

export function AuthProvider({ children }: any) {
  const navigate = useNavigate()
  const queryClient = useQueryClient()
  const { sendMessage } = useMessages()

  const {
    refetch: refreshSession,
    isLoading,
    data: session,
    error,
  } = useQuery<Session, AxiosError>({
    queryKey: [routes.auth.session],
    staleTime: 1000 * 60 * 2,
    retry: false,
  })

  useEffect(() => {
    if (error instanceof AxiosError && error.response?.status === HttpStatusCode.Unauthorized) {
      sendMessage({ message: 'Melde dich an, um die App nutzen zu können.', type: 'warning' })
      navigate('/auth/login')
    }
  }, [error])

  async function logout() {
    const res = await client.get(routes.auth.logout)
    if (res.status !== HttpStatusCode.Ok) sendMessage({ message: 'Bei der Abmeldung ist ein Fehler aufgetreten.', type: 'error' })
    else sendMessage({ message: 'Du wurdest erfolgreich abgemeldet.', type: 'success' })

    queryClient.removeQueries([routes.auth.session])
    navigate('/')
  }

  function getAuthRoles(clanTag: string | undefined) {
    if (!session?.user) return null

    const roles: AuthRole[] = []
    if (session.user.isAdmin) roles.push(AuthRole.Admin)

    if (session.user.leaderOf !== undefined && session.user.leaderOf.length > 0) {
      roles.push(AuthRole.AnyLeader)
      if (session.user.leaderOf.includes(urlDecodeTag(clanTag))) roles.push(AuthRole.ClanLeader)
    }

    if (session.user.coLeaderOf !== undefined && session.user.coLeaderOf.length > 0) {
      roles.push(AuthRole.AnyCoLeader)
      if (session.user.coLeaderOf.includes(urlDecodeTag(clanTag))) roles.push(AuthRole.ClanCoLeader)
    }

    if (session.user.memberOf !== undefined && session.user.memberOf.length > 0) {
      roles.push(AuthRole.AnyMember)
      if (session.user.memberOf.includes(urlDecodeTag(clanTag))) roles.push(AuthRole.ClanMember)
    }

    return roles.length > 0 ? roles : null
  }

  return (
    <authContext.Provider value={{ ...session, refreshSession, logout, getAuthRoles }}>
      {isLoading ? <LoadingScreen /> : children}
    </authContext.Provider>
  )
}

export function useAuth() {
  return useContext(authContext)
}
