import '@styles/components/Notifications.scss'
import { useCallback, useRef } from 'react'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faBell, faXmark } from '@fortawesome/free-solid-svg-icons'
import { useQuery } from '@tanstack/react-query'
import routes from '@api/routes'

export default function Notifications() {
  const contentRef = useRef<HTMLDivElement>(null)

  const { data: notifications, isLoading } = useQuery({
    queryKey: [routes.notifications],
    enabled: false,
    retry: false,
    staleTime: 1000 * 60,
  })

  const openNotifications = useCallback(() => {
    if (!contentRef.current) return
    contentRef.current?.classList.add('open')
  }, [contentRef])

  const closeNotifications = useCallback(() => {
    if (!contentRef.current) return
    contentRef.current?.classList.remove('open')
  }, [contentRef])

  return (
    <div className="Notifications">
      <a className="button" onClick={openNotifications}>
        <FontAwesomeIcon icon={faBell} />
        {true && <span className="amount">9+</span>}
      </a>
      <div className="content" ref={contentRef}>
        <div className="header">
          <a className="close">
            <FontAwesomeIcon icon={faXmark} onClick={closeNotifications} />
          </a>
          <span className="title">Benachrichtigungen</span>
        </div>
        <div className="notification"></div>
      </div>
    </div>
  )
}
