import '@styles/pages/Error.scss'
import errors from '@assets/content/errors.json'
import { useParams } from 'react-router-dom'
import { useEffect, useRef, useState } from 'react'
import useDocumentTitle from '@hooks/useDocumentTitle'
import Button from '@components/Button'
import Spacer from '@components/Spacer'
import Content from '@components/Content'
import Center from '@components/Center'

interface Error {
  title: string
  content: string[]
  gif?: string
}

export default function ErrorPage() {
  const { code } = useParams()
  const [error, setError] = useState<Error>(errors['unknown'])
  const heading = useDocumentTitle(`Fehler${error ? ` - ${error.title}` : ''}`)

  const contentRef = useRef<HTMLDivElement>(null)

  useEffect(() => {
    if (!code) return

    const err = Object.entries(errors).find((prop) => prop[0] == code)
    if (err) setError(err[1])
  }, [code])

  useEffect(() => {
    if (!contentRef.current) return

    contentRef.current.innerHTML = error.content.join('')
  }, [error])

  return (
    <main className="p-error">
      <Content>
        <Spacer size="medium" />
        {heading}
        <Center>
          <div ref={contentRef}></div>
          {error.gif && <img src={error.gif} loading="eager" alt="Fehler" />}
          <Button to={-2}>Zurück zur vorherigen Seite</Button>
        </Center>
      </Content>
    </main>
  )
}
