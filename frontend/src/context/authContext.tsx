import { createContext, useContext, useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useQuery } from '@tanstack/react-query'
import LoadingScreen from '../components/LoadingScreen'
import { AxiosError, HttpStatusCode } from 'axios'
import { AuthContext } from '../types/context'
import { Session } from '../api/types/auth'
import routes from '../api/routes'

const authContext = createContext({} as AuthContext)

export function AuthProvider({ children }: any) {
  const navigate = useNavigate()
  const [sessionData, setSessionData] = useState<Session>()

  const { refetch: refreshSession, isLoading } = useQuery<Session, AxiosError>({
    queryKey: [routes.auth.session],
    staleTime: 1000 * 60 * 2,
    retry: false,
    onSuccess: setSessionData,
    onError: (error) => {
      setSessionData(undefined)
      switch (error.response?.status) {
        case HttpStatusCode.Unauthorized:
        case HttpStatusCode.Ok:
          break
        case HttpStatusCode.RequestTimeout:
          navigate('/error/408')
          break
        case HttpStatusCode.InternalServerError:
          navigate('/error/500')
          break
        case HttpStatusCode.ServiceUnavailable:
          navigate('/error/503')
          break
        default:
          navigate('/error/unknown')
          break
      }
    },
  })

  const { refetch: fetchLogout } = useQuery({
    queryKey: [routes.auth.logout],
    enabled: false,
    retry: false,
  })

  async function logout() {
    setSessionData(undefined)
    await fetchLogout()
    await refreshSession()
    navigate('/')
  }

  return <authContext.Provider value={{ ...sessionData, refreshSession, logout }}>{isLoading ? <LoadingScreen /> : children}</authContext.Provider>
}

export function useAuth() {
  return useContext(authContext)
}
