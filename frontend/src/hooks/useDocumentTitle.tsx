import { useEffect } from 'react'

const titlePrefix = 'Lost Clans - '

export default function useDocumentTitle(title: string) {
  useEffect(() => {
    document.title = `${titlePrefix}${title}`
  }, [title])

  return <h1>{title}</h1>
}
