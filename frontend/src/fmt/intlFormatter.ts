export const dateTimeFormatter = new Intl.DateTimeFormat('de-DE', {
  dateStyle: 'medium',
  timeStyle: 'short',
})

export const numberFormatter = new Intl.NumberFormat('de-DE', {})

export function capitalizeFirstLetter(str: string): string {
  return str.charAt(0).toUpperCase() + str.slice(1)
}

export function addDaysToDate(date: Date, days: number): Date {
  date.setDate(date.getDate() + days)
  return date
}

export function timeUntil(date: Date): string {
  const diff = date.getTime() - new Date().getTime()
  const diffDays = Math.floor(diff / (1000 * 3600 * 24))
  const diffHours = Math.floor((diff % (1000 * 3600 * 24)) / (1000 * 3600))
  const diffMinutes = Math.floor((diff % (1000 * 3600)) / (1000 * 60))

  let res = ''
  if (diffDays > 0) res += `${diffDays}d `
  if (diffHours > 0) res += `${diffHours}h `
  if (diffMinutes > 0) res += `${diffMinutes}m`
  return res.trim()
}
