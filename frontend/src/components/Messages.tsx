import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { useMessage } from '@context/messageContext'
import '@styles/components/Messages.scss'
import { useCallback, useLayoutEffect, useRef, useState } from 'react'
import { faXmark } from '@fortawesome/free-solid-svg-icons'

export interface MessageProps {
  message: string
  type: 'error' | 'success'
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

  useLayoutEffect(() => {
    if (!messageRef.current) return

    messageRef.current.addEventListener('animationend', closeMessage)
    messageRef.current.addEventListener('transitionend', handleTransitionEnd)

    return () => {
      messageRef.current?.removeEventListener('animationend', closeMessage)
      messageRef.current?.removeEventListener('transitionend', handleTransitionEnd)
    }
  }, [])

  const closeMessage = useCallback(() => {
    setOpen(false)
  }, [setOpen])

  const handleTransitionEnd = useCallback(() => {
    removeMessage(id)
  }, [removeMessage, id])

  return (
    <div ref={messageRef} className={`Message ${type}${open ? ' open' : ''}`}>
      <div className="content">
        <span>{message}</span>
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
