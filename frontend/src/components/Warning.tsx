import { faWarning } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'

type WarningProps = {
  children?: React.ReactNode
}

export default function Warning({ children }: WarningProps) {
  if (!children) return null

  return (
    <div className="Warning">
      <FontAwesomeIcon icon={faWarning} className="icon" />
      <div className="content">{children}</div>
    </div>
  )
}
