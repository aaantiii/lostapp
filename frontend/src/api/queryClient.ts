import { QueryClient } from '@tanstack/react-query'
import { replaceQueryParams, replaceRouteIds } from './urlBuilder'
import client from './client'

export default new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 1000 * 60 * 3,
      retry: 2,
      cacheTime: 1000 * 60,
      refetchOnWindowFocus: false,
      queryFn: async ({ queryKey: [path, ids, params] }: any) => {
        if (typeof path !== 'string') throw new Error('invalid path in query: path must be string')

        const uri = replaceQueryParams(replaceRouteIds(path, ids), params)

        const { data, status } = await client.get(uri)

        if (status >= 400) throw new Error('get request failed')

        return data
      },
    },
  },
})
