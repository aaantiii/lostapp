import '../scss/components/ExpandableList.scss'
import { useLocation, useNavigate } from 'react-router-dom'
import { useCallback, useRef } from 'react'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faAngleUp } from '@fortawesome/free-solid-svg-icons'
import { IconProp } from '@fortawesome/fontawesome-svg-core'

interface ExpandableListProps {
  title: string
  children: JSX.Element[] | JSX.Element
  expanded?: boolean
}

interface ExpandableListItemProps {
  title: string
  href: string
  icon?: IconProp
  disabled?: boolean
}

export default function ExpandableList({ title, children, expanded }: ExpandableListProps) {
  const listRef = useRef<HTMLDivElement>(null)

  const toggleList = useCallback(() => {
    listRef.current?.classList.toggle('expanded')
  }, [listRef])

  return (
    <div ref={listRef} className={`ExpandableList ${expanded ? 'expanded' : ''}`}>
      <a onClick={toggleList} className={`toggle ${expanded ? 'expanded' : ''}`}>
        <span>{title}</span>
        <FontAwesomeIcon icon={faAngleUp} />
      </a>
      <ul className={expanded ? 'expanded' : ''}>{children}</ul>
    </div>
  )
}

export function ExpandableListItem({ title, href, icon, disabled }: ExpandableListItemProps) {
  const navigate = useNavigate()
  const location = useLocation()

  const isActive = location.pathname.replaceAll('/', '') === href.replaceAll('/', '')

  const handleNavigate = useCallback(() => {
    if (isActive) return
    navigate(href)
  }, [location, href])

  return (
    <li className={`item${isActive ? ' active' : ''}${disabled ? ' disabled' : ''}`}>
      <a onClick={handleNavigate}>
        {icon && (
          <div className="icon">
            <FontAwesomeIcon icon={icon} />
          </div>
        )}
        <span>{title}</span>
      </a>
    </li>
  )
}
