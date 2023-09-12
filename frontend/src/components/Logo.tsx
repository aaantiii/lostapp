import '@styles/components/Logo.scss'
import lostLogo from '@assets/img/lost_logo.webp'
import { useNavigate } from 'react-router-dom'

export default function Logo() {
  const navigate = useNavigate()
  return (
    <a onClick={() => navigate('/')} title="Startseite" className="Logo">
      <span className="title">Lost Clans</span>
      <img src={lostLogo} alt="Lost Clan Logo" loading="eager" />
    </a>
  )
}
