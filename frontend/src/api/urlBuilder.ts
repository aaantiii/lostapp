export function replaceRouteIds(path: string, ids: any): string {
  let uri = path
  if (ids) {
    for (const [prop, value] of Object.entries<string | number | undefined>(ids)) {
      if (!value) throw new Error(`invalid id: ${prop} is undefined`)
      uri = uri.replace(`:${prop}`, value.toString())
    }
  }

  return encodeURI(uri)
}

export function replaceQueryParams(uri: string, params: any): string {
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
