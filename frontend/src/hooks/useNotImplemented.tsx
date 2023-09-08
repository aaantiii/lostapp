import { useEffect } from 'react'
import { useNavigate } from 'react-router-dom'

// Todo: create dialog instead of error navigate
export default function useNotImplemented() {
  const navigate = useNavigate()

  useEffect(() => {
    navigate('/error/not-implemented')
  }, [])
}
