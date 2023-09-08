import { QueryClient } from '@tanstack/react-query'
import { GET } from './queryFunctions'

export default new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 1000 * 60 * 3,
      queryFn: GET,
      retry: 2,
      cacheTime: 1000 * 60,
      refetchOnWindowFocus: false,
    },
  },
})
