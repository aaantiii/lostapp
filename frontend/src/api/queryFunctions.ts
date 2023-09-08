import client from './client'

export async function GET({ queryKey: [path, ids, params] }: any) {
  if (typeof path !== 'string') throw new Error('invalid path in query: path must be string')

  const uri = addQueryParams(addRouteIds(path, ids), params)

  const { data, status } = await client.get(uri)

  if (status >= 400) throw new Error('get request failed')

  return data
}

export async function PUT<T>({ queryKey: [path, ids, data] }: { queryKey: [string, any, T] }) {
  const uri = addRouteIds(path, ids)

  const { status } = await client.put<T>(uri, data)

  if (status >= 300) throw new Error('put request failed')

  return data
}

function addRouteIds(path: string, ids: any): string {
  let uri = path
  if (ids) {
    for (const [prop, value] of Object.entries<string | number | undefined>(ids)) {
      if (!value) continue
      uri = uri.replace(`:${prop}`, value.toString())
    }
  }

  return encodeURI(uri)
}

function addQueryParams(uri: string, params: any): string {
  if (params) {
    uri += '?'
    for (const [prop, value] of Object.entries(params)) {
      if (Array.isArray(value)) {
        value.forEach((v) => (uri += `${prop}=${v}&`))
      } else {
        uri += `${prop}=${value}&`
      }
    }
    uri = uri.slice(0, -1) // remove trailing '&'
  }

  return encodeURI(uri)
}
