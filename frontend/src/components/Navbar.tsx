import { useLocation } from 'react-router-dom'
import '../scss/components/Navbar.scss'
import Logo from './Logo'
import { useCallback, useEffect, useRef, useState } from 'react'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faBars, faXmark } from '@fortawesome/free-solid-svg-icons'
import UserAvatar from './UserAvatar'
import useNavItems from '../hooks/useNavItems'
import Notifications from './Notifications'
import { useAuth } from '../context/authContext'

export default function Navbar() {
  const { discordUser, logout } = useAuth()
  const location = useLocation()
  const [isOpen, setIsOpen] = useState(false)
  const vNavRef = useRef<HTMLElement>(null)
  const vNavItemsRef = useRef<HTMLDivElement>(null)

  const navItems = useNavItems()

  const toggleNav = useCallback(() => {
    if (vNavRef.current) setIsOpen(vNavRef.current.classList.toggle('open'))
  }, [vNavRef])

  // for extend collapse buttons
  const setNavItemsExpanded = useCallback(
    (expanded: boolean) => {
      if (!vNavItemsRef.current) return
      for (const item of vNavItemsRef.current.children) {
        expanded ? item.classList.add('expanded') : item.classList.remove('expanded')
      }
    },
    [vNavItemsRef]
  )

  useEffect(() => {
    if (isOpen) toggleNav()
  }, [location])

  return (
    <div className="Navbar">
      <nav className="hnav">
        <a onClick={toggleNav} className="toggle-vnav">
          {isOpen ? <FontAwesomeIcon icon={faXmark} /> : <FontAwesomeIcon icon={faBars} />}
        </a>
        <Logo />
      </nav>
      <nav className="vnav" ref={vNavRef}>
        <div className="header">
          {discordUser && (
            <div className="top">
              <a onClick={logout} className="logout">
                abmelden
              </a>
              <Notifications />
            </div>
          )}
          <UserAvatar />
        </div>
        <div className="functions">
          <a onClick={() => setNavItemsExpanded(false)} className="function">
            einklappen
          </a>
          <a onClick={() => setNavItemsExpanded(true)} className="function">
            ausklappen
          </a>
        </div>

        <div className="items" ref={vNavItemsRef}>
          {navItems}
        </div>
      </nav>
    </div>
  )
}
