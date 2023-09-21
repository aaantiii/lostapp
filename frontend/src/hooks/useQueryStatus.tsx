import { useEffect } from 'react'

interface QueryStatusProps {
  status: 'success' | 'error' | 'loading'
  fetchStatus: 'fetching' | 'idle' | 'paused'
  onSuccess?: () => void
  onError?: () => void
}

export default function useQueryStatus({ status, fetchStatus, onSuccess, onError }: QueryStatusProps) {
  useEffect(() => {
    if (fetchStatus !== 'idle') return

    if (status === 'success') onSuccess?.()
    else if (status === 'error') onError?.()
  }, [fetchStatus])
}
