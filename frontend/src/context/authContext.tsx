import { createContext, useContext, useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import { useQuery, useQueryClient } from '@tanstack/react-query'
import LoadingScreen from '@components/LoadingScreen'
import { AxiosError, HttpStatusCode } from 'axios'
import { AuthContext } from './types'
import { Session } from '@api/types/auth'
import routes from '@api/routes'

const authContext = createContext({} as AuthContext)

export function AuthProvider({ children }: any) {
  const navigate = useNavigate()
  const queryClient = useQueryClient()

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

  const { refetch: fetchLogout } = useQuery({
    queryKey: [routes.auth.logout],
    enabled: false,
    retry: false,
    cacheTime: 0,
  })

  useEffect(() => {
    if (!error) return

    queryClient.setQueryData([routes.auth.session], () => undefined)
    switch (error.response?.status) {
      case HttpStatusCode.Unauthorized:
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
  }, [error])

  async function logout() {
    await fetchLogout()
    queryClient.removeQueries([routes.auth.session])
    navigate('/')
  }

  return <authContext.Provider value={{ ...session, refreshSession, logout }}>{isLoading ? <LoadingScreen /> : children}</authContext.Provider>
}

export function useAuth() {
  return useContext(authContext)
}
