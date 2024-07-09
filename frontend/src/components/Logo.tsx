import '@styles/components/Logo.scss'
import lostLogo from '@assets/img/lost_logo.webp'
import { useNavigate } from 'react-router-dom'

type LogoProps = {
  className?: string
}

export default function Logo({ className }: LogoProps) {
  const navigate = useNavigate()

  function handleClick() {
    navigate('/#')
  }

  return (
    <a onClick={handleClick} title="Dashboard" className={`Logo ${className}`}>
      <span className="title">Lost Clans</span>
    </a>
  )
}

// export default function Logo({ className }: LogoProps) {
//   const navigate = useNavigate()

//   function handleClick() {
//     navigate('/#')
//     scrollTo(0, 0)
//   }

//   return (
//     <a onClick={handleClick} title="Startseite" className={`Logo ${className}`}>
//       <span className="title">Lost Clans</span>
//       <img src={lostLogo} alt="Lost Clan Logo" loading="eager" />
//     </a>
//   )
// }
