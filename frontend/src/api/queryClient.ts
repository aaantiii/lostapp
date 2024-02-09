import { QueryClient } from '@tanstack/react-query'
import client from './client'
import { replaceIds } from './routes'

export default new QueryClient({
  defaultOptions: {
    queries: {
      staleTime: 1000 * 60 * 3,
      retry: 1,
      cacheTime: 1000 * 60 * 2,
      queryFn: defaultQueryFn,
    },
  },
})

// defaultQueryFn makes a get request to the api
async function defaultQueryFn({ queryKey: [path, ids, params] }: any) {
  if (typeof path !== 'string') throw new Error('invalid path in query: path must be string')

  const { data, status, statusText } = await client.get(replaceIds(path, ids), { params })

  if (import.meta.env.DEV) console.log(`GET ${path} (${status} ${statusText})\n${JSON.stringify(data, null, 2)}`)

  if (status >= 400) throw new Error('get request failed')

  return data
}
