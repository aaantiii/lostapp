import { createContext, useContext, useEffect } from 'react'
import useLocalStorage from '@hooks/useLocalStorage'

export type Theme = 'light' | 'dark'

type ThemeContext = {
  theme: Theme
  toggleTheme: () => void
}

const defaultTheme = 'dark'

const themeContext = createContext({ theme: defaultTheme } as ThemeContext)

export function ThemeProvider({ children }: { children: JSX.Element | JSX.Element[] }) {
  const [theme, setTheme] = useLocalStorage<Theme>('theme', defaultTheme)

  function toggleTheme() {
    setTheme((prev) => (prev === 'dark' ? 'light' : 'dark'))
  }

  useEffect(() => {
    document.body.setAttribute('data-theme', theme)
  }, [theme])

  return <themeContext.Provider value={{ theme, toggleTheme }}>{children}</themeContext.Provider>
}

export function useTheme() {
  return useContext(themeContext)
}
