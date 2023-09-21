import { useEffect } from 'react'

export default function useDocumentTitle(title: string) {
  useEffect(() => {
    document.title = title
  }, [title])

  return <h1>{title}</h1>
}
