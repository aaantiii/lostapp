import { useEffect, useState } from 'react'

export default function useDebouncedValue<T>(initialValue: T, delay = 200): [T, React.Dispatch<React.SetStateAction<T>>] {
  const [actualValue, setActualValue] = useState(initialValue)
  const [debounceValue, setDebounceValue] = useState(initialValue)

  useEffect(() => {
    const timeoutId = setTimeout(() => setDebounceValue(actualValue), delay)
    return () => clearTimeout(timeoutId)
  }, [actualValue, delay])

  return [debounceValue, setActualValue]
}
