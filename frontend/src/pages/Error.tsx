import { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import errors from '@assets/content/errors.json'
import Center from '@components/Center'
import Spacer from '@components/Spacer'

export default function Error() {
  const { code } = useParams()
  const [error, setError] = useState(errors['500'])

  useEffect(() => {
    const err = errors[code as keyof typeof errors]
    if (err) setError(err)
    else setError(errors.unknown)
  }, [code])

  return (
    <main>
      <h1>{error.title}</h1>
      <p>{error.description}</p>
      <Spacer size="medium" />
      <Center>{error.image && <img src={`/gifs/${error.image}`} alt={error.title} />}</Center>
    </main>
  )
}
