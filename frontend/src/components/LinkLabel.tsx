import '@styles/components/LinkLabel.scss'
import { faUpRightFromSquare } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { AnchorHTMLAttributes, forwardRef } from 'react'
import { useNavigate } from 'react-router-dom'
import { IconProp } from '@fortawesome/fontawesome-svg-core'

type LinkLabelProps = AnchorHTMLAttributes<HTMLAnchorElement> & {
  icon?: IconProp
  children: string
}

const LinkLabel = forwardRef<HTMLAnchorElement, LinkLabelProps>(({ children, icon, ...anchorProps }, fRef) => {
  const navigate = useNavigate()

  if (anchorProps.target === '_self') {
    return (
      <a {...anchorProps} className="LinkLabel" ref={fRef} href={undefined} onClick={() => navigate(anchorProps.href!)}>
        <span>{children}</span>
        <FontAwesomeIcon icon={icon ?? faUpRightFromSquare} />
      </a>
    )
  }

  return (
    <a {...anchorProps} className="LinkLabel" ref={fRef}>
      <span>{children}</span>
      <FontAwesomeIcon icon={icon ?? faUpRightFromSquare} />
    </a>
  )
})

export default LinkLabel
