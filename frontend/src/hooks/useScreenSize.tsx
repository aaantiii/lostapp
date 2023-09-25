import { useEffect, useState } from 'react'

export enum ScreenSize {
  Mobile = 480,
  TabletPortrait = 768,
  TabletLandscape = 1024,
  Desktop = 1025,
}

export default function useScreenSize() {
  const [screenSize, setScreenSize] = useState(getCurrentSize())

  useEffect(() => {
    function handleResize() {
      setScreenSize(getCurrentSize())
    }

    window.addEventListener('resize', handleResize)
    return () => removeEventListener('resize', handleResize)
  }, [])

  function getCurrentSize(): ScreenSize {
    if (window.innerWidth <= ScreenSize.Mobile) return ScreenSize.Mobile
    if (window.innerWidth <= ScreenSize.TabletPortrait) return ScreenSize.TabletPortrait
    if (window.innerWidth <= ScreenSize.TabletLandscape) return ScreenSize.TabletLandscape
    return ScreenSize.Desktop
  }

  return screenSize
}
