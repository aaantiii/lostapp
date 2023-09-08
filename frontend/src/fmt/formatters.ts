export const dateTimeFormatter = new Intl.DateTimeFormat('de-DE', {
  year: 'numeric',
  month: '2-digit',
  day: '2-digit',
})

export const numberFormatter = new Intl.NumberFormat('de-DE', {})

export function capitalizeFirstLetter(str: string): string {
  return str.charAt(0).toUpperCase() + str.slice(1)
}

export function addDaysToDate(date: Date, days: number): Date {
  date.setDate(date.getDate() + days)
  return date
}
