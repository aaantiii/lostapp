import { createContext, useContext, useState } from 'react'
import { MessageProps } from '@components/Messages'
import { MessageContext } from './types'

const messageContext = createContext({} as MessageContext)

export function MessageProvider({ children }: { children: JSX.Element | JSX.Element[] }) {
  const [messages, setMessages] = useState<MessageProps[]>([])
  const [currentId, setCurrentId] = useState(0)

  function sendMessage(message: Omit<MessageProps, 'id'>) {
    const props: MessageProps = { ...message, id: currentId }
    setCurrentId((id) => id + 1)
    setMessages((messages) => [...messages, props])
  }

  function removeMessage(id: number) {
    setMessages((messages) => messages.filter((message) => message.id != id))
  }

  return <messageContext.Provider value={{ messages, sendMessage, removeMessage } satisfies MessageContext}>{children}</messageContext.Provider>
}

export function useMessage() {
  return useContext(messageContext)
}
