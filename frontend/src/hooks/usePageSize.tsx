import useScreenSize, { ScreenSize } from './useScreenSize'

export default function usePageSize(mobile: number, tablet: number, desktop?: number): number {
  const screenSize = useScreenSize()

  if (screenSize <= ScreenSize.Mobile) return mobile
  if (screenSize <= ScreenSize.TabletPortrait) return tablet
  if (desktop) return desktop

  return tablet
}
