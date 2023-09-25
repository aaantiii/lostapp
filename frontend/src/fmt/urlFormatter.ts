export function uriSafe(str: string): string {
  return str.replaceAll('&', '').replaceAll('#', '').replaceAll('+', '')
}
