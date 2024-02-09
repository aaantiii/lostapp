import { createContext, useContext, useEffect, useState } from 'react'
import { MessageProps } from '@components/Messages'
import useScreenSize, { ScreenSize } from '@hooks/useScreenSize'

type MessagesContext = {
  messages: MessageProps[]
  sendMessage: (message: Omit<MessageProps, 'id'>) => void
  removeMessage: (id: string) => void
}

const messagesContext = createContext({} as MessagesContext)

export function MessagesProvider({ children }: { children: React.ReactNode }) {
  const [messages, setMessages] = useState<MessageProps[]>([])
  const [messageQueue, setMessageQueue] = useState(messages)
  const screenSize = useScreenSize()

  const maxVisible = screenSize <= ScreenSize.TabletPortrait ? 2 : 3

  useEffect(() => {
    if (messages.length >= maxVisible) return

    const nextMessage = messageQueue.shift()
    if (!nextMessage) return

    setMessages((messages) => [...messages, nextMessage])
  }, [messages, messageQueue])

  function sendMessage(props: Omit<MessageProps, 'id'>) {
    if (messageQueue.length >= 6) return
    const message: MessageProps = { ...props, id: crypto.randomUUID() }
    setMessageQueue((prev) => [...prev, message])
  }

  function removeMessage(id: string) {
    setMessages((messages) => messages.filter((message) => message.id !== id))
  }

  return <messagesContext.Provider value={{ messages, sendMessage, removeMessage }}>{children}</messagesContext.Provider>
}

export function useMessages() {
  return useContext(messagesContext)
}
