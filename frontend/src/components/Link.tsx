import { useNavigate } from 'react-router-dom'

interface LinkProps {
  children: string
  to: string
}

export default function Link({ children, to }: LinkProps) {
  const navigate = useNavigate()
  return (
    <a className="Link" onClick={() => navigate(to)}>
      {children}
    </a>
  )
}
