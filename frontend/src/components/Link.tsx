import '@styles/components/Link.scss'
import { useNavigate } from 'react-router-dom'

interface LinkProps {
  children: string
  to?: string
  href?: string
  newWindow?: boolean
}

export default function Link({ children, to, href, newWindow }: LinkProps) {
  const navigate = useNavigate()
  return (
    <a className="Link" onClick={to ? () => navigate(to) : undefined} href={href} target={newWindow ? '_blank' : '_self'}>
      {children}
    </a>
  )
}
