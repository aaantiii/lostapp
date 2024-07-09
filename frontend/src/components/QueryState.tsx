import LoadingScreen from './LoadingScreen'
import { ReactNode } from 'react'
import { ApiError } from '@api/types/models'
import { HttpStatusCode } from 'axios'
import { useNavigate } from 'react-router-dom'

type QueryStateProps = {
  loading: boolean
  error: ApiError | null
  loader?: JSX.Element
}

const DEFAULT_ERROR = 'Die Anfrage konnte nicht verarbeitet werden. Bitte versuche es später erneut.'
const ERROR_204 = 'Keine Daten verfügbar.'

export default function QueryState({ loading: isLoading, error, loader }: QueryStateProps) {
  const navigate = useNavigate()

  let child: ReactNode = null
  if (isLoading) {
    child = loader ?? <LoadingScreen />
  } else if (error?.response) {
    if (error.response.status === HttpStatusCode.Unauthorized) {
      navigate('/auth/login')
    } else if (error.response.status === HttpStatusCode.NoContent) {
      child = <ErrorMessage message={ERROR_204} />
    } else {
      child = <ErrorMessage message={error.response.data.message ?? DEFAULT_ERROR} />
    }
  } else if (!error) {
    child = <ErrorMessage message={ERROR_204} />
  } else {
    child = <ErrorMessage message={DEFAULT_ERROR} />
  }

  return <div className="QueryState">{child}</div>
}

function ErrorMessage({ message }: { message: string }) {
  return <p className="error">{message}</p>
}
