import '@styles/components/UserAvatar.scss'
import { faUser } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { useNavigate } from 'react-router-dom'
import { User } from '@api/types/user'

type UserAvatarProps = {
  user: User | undefined
  to?: string
  className?: string
}

export default function UserAvatar({ user, to, className }: UserAvatarProps) {
  const navigate = useNavigate()

  return (
    <div className={`UserAvatar ${className}`} onClick={to ? () => navigate(to) : undefined}>
      {user?.avatarUrl ? <img src={user.avatarUrl} alt="User Avatar" loading="lazy" /> : <FontAwesomeIcon icon={faUser} />}
    </div>
  )
}
