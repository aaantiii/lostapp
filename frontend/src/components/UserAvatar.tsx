import '@styles/components/UserAvatar.scss'
import { useAuth } from '@context/authContext'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faUser } from '@fortawesome/free-solid-svg-icons'
import { useNavigate } from 'react-router-dom'
import useDashboardNavigate from '@hooks/useDashboardNavigate'

interface UserAvatarProps {
  onClick?: () => void
  title?: string
}

export default function UserAvatar({ onClick, title }: UserAvatarProps) {
  const navigate = useNavigate()
  const overviewRedirect = useDashboardNavigate()
  const { discordUser } = useAuth()

  return discordUser ? (
    <a title={title ?? 'Dashboard'} onClick={onClick ?? overviewRedirect} className="UserAvatar">
      <img src={discordUser.avatarUrl} alt="Discord Avatar" />
      <span>{discordUser.username}</span>
    </a>
  ) : (
    <a title="Anmelden" onClick={() => navigate('/auth/login')} className="UserAvatar noauth">
      <FontAwesomeIcon icon={faUser} />
      <span>Login</span>
    </a>
  )
}
