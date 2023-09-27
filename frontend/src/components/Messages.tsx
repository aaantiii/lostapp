import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { useMessage } from '@context/messageContext'
import '@styles/components/Messages.scss'
import { useEffect, useRef, useState } from 'react'
import { faWarning, faXmark } from '@fortawesome/free-solid-svg-icons'

export interface MessageProps {
  message: string
  type: 'error' | 'success' | 'warning'
  id: number
}

export default function Messages() {
  const { messages } = useMessage()

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
  const [open, setOpen] = useState(true)
  const { removeMessage } = useMessage()

  useEffect(() => {
    if (!messageRef.current) return

    messageRef.current.addEventListener('animationend', closeMessage) // indicator animation
    messageRef.current.addEventListener('transitionend', handleTransitionEnd) // opacity transition

    return () => {
      messageRef.current?.removeEventListener('animationend', closeMessage)
      messageRef.current?.removeEventListener('transitionend', handleTransitionEnd)
    }
  }, [])

  function closeMessage() {
    setOpen(false)
  }

  function handleTransitionEnd() {
    removeMessage(id)
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
        <div className="indicator"></div>
      </div>
    </div>
  )
}
