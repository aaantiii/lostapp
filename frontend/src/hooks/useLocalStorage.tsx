import { useEffect, useState } from 'react'

// useLocalStorage is a hook that allows storing primitives, maps and objects in localStorage.
export default function useLocalStorage<T>(key: string, initialValue: T) {
  const [storedValue, setStoredValue] = useState<T>(() => {
    const item = localStorage.getItem(key)
    if (!item) return initialValue
    if (typeof initialValue === 'string') return item
    return JSON.parse(item)
  })

  useEffect(() => {
    let stringifiedValue = ''

    if (typeof storedValue === 'string') stringifiedValue = storedValue
    else if (storedValue instanceof Map) stringifiedValue = JSON.stringify(Array.from(storedValue))
    else stringifiedValue = JSON.stringify(storedValue)

    localStorage.setItem(key, stringifiedValue)
  }, [storedValue])

  return [storedValue, setStoredValue] as const
}
