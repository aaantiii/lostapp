export function buildURI(path: string, ids?: any): string {
  if (ids) {
    for (const [prop, value] of Object.entries<string | number | undefined>(ids)) {
      if (value === undefined) throw new Error(`invalid id: ${prop} is ${value}`)
      path = path.replace(`:${prop}`, value.toString())
    }
  }

  return encodeURI(path)
}
