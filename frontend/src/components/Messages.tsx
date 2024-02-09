import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { useMessages } from '@context/messagesContext'
import '@styles/components/Messages.scss'
import { useEffect, useRef, useState } from 'react'
import { faWarning, faXmark } from '@fortawesome/free-solid-svg-icons'

export type MessageProps = {
  message: string
  type: 'error' | 'success' | 'warning'
  id: string
}

export default function Messages() {
  const { messages } = useMessages()

  return (
    <div className="Messages">
      {messages.map((props) => (
        <Message {...props} key={props.id} />
      ))}
    </div>
  )
}

function Message({ message, type, id }: MessageProps) {
  const messageRef = useRef<HTMLDivElement>(null)
  const indicatorRef = useRef<HTMLDivElement>(null)
  const [open, setOpen] = useState(true)
  const { removeMessage } = useMessages()

  useEffect(() => {
    function handleTransitionEnd() {
      removeMessage(id)
    }

    if (!messageRef.current || !indicatorRef.current) return

    indicatorRef.current.addEventListener('animationend', closeMessage) // indicator animation
    messageRef.current.addEventListener('transitionend', handleTransitionEnd) // opacity transition

    return () => {
      indicatorRef.current?.removeEventListener('animationend', closeMessage)
      messageRef.current?.removeEventListener('transitionend', handleTransitionEnd)
    }
  }, [])

  function closeMessage() {
    setOpen(false)
  }

  return (
    <div ref={messageRef} className={`Message ${type}${open ? ' open' : ''}`}>
      <div className="body">
        <div className="content">
          {(type === 'warning' || type === 'error') && <FontAwesomeIcon icon={faWarning} className="icon" />}
          <span>{message}</span>
        </div>
        <a className="close-button" onClick={closeMessage}>
          <FontAwesomeIcon icon={faXmark} />
        </a>
      </div>
      <div className="close-indicator">
        <div ref={indicatorRef} className="indicator"></div>
      </div>
    </div>
  )
}
