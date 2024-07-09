import '@styles/components/Navbar.scss'
import Logo from './Logo'
import UserAvatar from './UserAvatar'
import { useEffect, useState } from 'react'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { useAuth } from '@context/authContext'
import { useLocation, useNavigate } from 'react-router-dom'
import { IconProp } from '@fortawesome/fontawesome-svg-core'
import useScreenSize, { ScreenSize } from '@hooks/useScreenSize'
import { faBars, faDashboard, faRankingStar, faRightToBracket, faShieldHalved, faUserShield, faXmark } from '@fortawesome/free-solid-svg-icons'
import { AuthRole } from '@api/types/auth'

type VNavItemProps = {
  children?: any
  to?: string
  icon?: JSX.Element | IconProp
  title?: string
  hideCollapsed?: boolean
}

export default function Navbar() {
  const [vNavOpen, setVNavOpen] = useState(false)
  const vNavItems = useNavItems()
  const { pathname } = useLocation()
  const screenSize = useScreenSize()

  useEffect(() => {
    if (screenSize <= ScreenSize.TabletLandscape) setVNavOpen(false)
  }, [pathname])

  useEffect(() => {
    document.body.setAttribute('data-nav-open', vNavOpen.toString())
  }, [vNavOpen])

  function handleToggleKeyDown(e: React.KeyboardEvent<HTMLAnchorElement>) {
    if (e.key === 'Enter') setVNavOpen((prev) => !prev)
  }

  return (
    <div className="Navbar">
      <div className="hnav" role="navigation">
        <a className="vnav-toggler" onClick={() => setVNavOpen((prev) => !prev)} onKeyDown={handleToggleKeyDown} tabIndex={0} title="Toggle Navbar">
          {vNavOpen ? <FontAwesomeIcon icon={faXmark} /> : <FontAwesomeIcon icon={faBars} />}
        </a>
        <Logo />
      </div>
      <nav className="vnav">{vNavItems}</nav>
    </div>
  )
}

function VNavItem({ children, to, icon, title, hideCollapsed }: VNavItemProps) {
  const navigate = useNavigate()
  const { pathname } = useLocation()
  const iconElement = icon && Object.hasOwn(icon as object, 'icon') ? <FontAwesomeIcon icon={icon as IconProp} /> : (icon as JSX.Element | undefined)

  const isLink = to !== undefined
  const isActive = to === pathname
  let className = 'item'
  if (isLink) className += ' link'
  if (isActive) className += ' active'
  if (hideCollapsed) className += ' hide-collapsed'

  function handleClick() {
    if (isActive || to === undefined) return
    navigate(to)
  }

  function handleKeyDown(e: React.KeyboardEvent<HTMLAnchorElement>) {
    switch (e.key) {
      case 'Enter':
        handleClick()
        break
      case 'ArrowUp':
        const prev = e.currentTarget.previousElementSibling
        if (prev) (prev as HTMLAnchorElement).focus()
        break
      case 'ArrowDown':
        const next = e.currentTarget.nextElementSibling
        if (next) (next as HTMLAnchorElement).focus()
        break
    }
  }

  return (
    <a className={className} title={title} onClick={isLink ? handleClick : undefined} tabIndex={isLink ? 0 : undefined} onKeyDown={handleKeyDown}>
      {iconElement && <div className="icon">{iconElement}</div>}
      {children}
    </a>
  )
}

function VNavItemGroup({ children }: { children: JSX.Element | JSX.Element[] }) {
  return <div className="group">{children}</div>
}

function useNavItems() {
  const { user, getAuthRoles } = useAuth()
  const userRoles = getAuthRoles(undefined)

  function navItems() {
    const items: JSX.Element[] = []

    if (!user) {
      items.push(
        <VNavItemGroup key="auth">
          <VNavItem to="/auth/login" title="Login" icon={faRightToBracket}>
            <span>Login</span>
          </VNavItem>
        </VNavItemGroup>
      )
      return items
    }

    items.push(
      <VNavItemGroup key="avatar">
        <VNavItem to="/account" icon={<UserAvatar user={user} />} title="Mein Account">
          <span>{user.name}</span>
        </VNavItem>
      </VNavItemGroup>
    )

    // admin items
    if (userRoles?.includes(AuthRole.Admin)) {
    }

    // member items
    items.push(
      <VNavItemGroup key="member1">
        <VNavItem to="/dashboard" icon={faUserShield} title="Dashboard">
          <span>Dashboard</span>
        </VNavItem>
        <VNavItem to="/clans" icon={faShieldHalved} title="LOST Clans">
          <span>LOST Clans</span>
        </VNavItem>
        <VNavItem to="/leaderboard" icon={faRankingStar} title="Leaderboard">
          <span>Leaderboard</span>
        </VNavItem>
      </VNavItemGroup>
    )

    return items
  }

  return navItems()
}
