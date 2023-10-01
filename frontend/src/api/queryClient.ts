import { QueryClient } from '@tanstack/react-query'
import { buildURI } from './urlBuilder'
import client from './client'

export default new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 1000 * 60 * 3,
      retry: 1,
      cacheTime: 1000 * 60 * 2,
      queryFn: async ({ queryKey: [path, ids, params] }: any) => {
        if (typeof path !== 'string') throw new Error('invalid path in query: path must be string')

        const { data, status } = await client.get(buildURI(path, ids), { params })

        if (status >= 400) throw new Error('get request failed')

        return data
      },
    },
  },
})
