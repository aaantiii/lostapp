import '@styles/components/UserAvatar.scss'
import { DiscordUser } from '@api/types/user'

interface UserAvatarProps {
  user: DiscordUser
  title?: string
  onClick?: () => void
  noHover?: boolean
  nameFirst?: boolean
}

export default function UserAvatar({ onClick, title, user, noHover, nameFirst }: UserAvatarProps) {
  return (
    <a title={title} onClick={onClick} className={`UserAvatar${nameFirst ? ' reverse' : ''}`} style={noHover ? { pointerEvents: 'none' } : undefined}>
      <img src={user.avatarUrl} alt="Discord Avatar" />
      <span>{user.name}</span>
    </a>
  )
}
