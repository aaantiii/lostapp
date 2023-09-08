// import { useState } from 'react'

// export default function useLocalStorage<T>(key: string, defaultValue?: T) {
//   const [storedValue, setStoredValue] = useState<T | undefined>(() => {
//     const item = localStorage.getItem(key)
//     return item ? JSON.parse(item) : defaultValue
//   })

//   function setValue(value?: T) {
//     setStoredValue(value)
//     localStorage.setItem(key, JSON.stringify(value))
//   }

//   return [storedValue, setValue]
// }
