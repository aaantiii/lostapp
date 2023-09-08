import { useEffect } from 'react'
import useDocumentTitle from '../../../hooks/useDocumentTitle'
import useDashboardNavigate from '../../../hooks/useDashboardNavigate'
import LoadingScreen from '../../../components/LoadingScreen'
import { useMessage } from '../../../context/messageContext'
import { useAuth } from '../../../context/authContext'
import { useNavigate } from 'react-router-dom'

export default function LoginSuccess() {
  useDocumentTitle('Anmeldung erfolgreich')

  const navigate = useNavigate()
  const navigateToDashboard = useDashboardNavigate()
  const { discordUser } = useAuth()
  const { sendMessage } = useMessage()

  useEffect(() => {
    if (!discordUser) return

    sendMessage({
      message: `Hey ${discordUser.username}, deine Anmeldung war erfolgreich!`,
      type: 'success',
    })
    navigateToDashboard()
  }, [discordUser])

  useEffect(() => {
    const timeoutId = setTimeout(() => navigate('/'), 3000)
    return () => clearTimeout(timeoutId)
  }, [])

  return <LoadingScreen></LoadingScreen>
}
