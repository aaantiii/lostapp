import '@styles/components/TabView.scss'
import { IconProp } from '@fortawesome/fontawesome-svg-core'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { KeyboardEvent, ReactNode, useEffect, useLayoutEffect, useRef, useState } from 'react'
import { useSearchParams } from 'react-router-dom'
import { Helmet } from 'react-helmet-async'

export type TabViewProps = {
  defaultTab: string
  tabs: Tab[]
}

type Tab = {
  id: string
  title: string
  icon: IconProp
  element: JSX.Element
}

export default function TabView({ defaultTab, tabs }: TabViewProps) {
  const [searchParams, setSearchParams] = useSearchParams({ tab: defaultTab })
  const [activeTab, setActiveTab] = useState<Tab | undefined>(undefined)
  const switchRef = useRef<HTMLDivElement>(null)

  useLayoutEffect(() => {
    setActiveTab(tabs.find((tab) => tab.id === searchParams.get('tab')) ?? tabs.find((tab) => tab.id === defaultTab))
  }, [searchParams])

  useEffect(() => {
    if (!switchRef.current) return

    for (const child of switchRef.current.children) {
      if (!child.classList.contains('active')) continue

      const elm = child as HTMLAnchorElement
      switchRef.current.scrollTo({
        left: switchRef.current.scrollWidth - (switchRef.current.scrollWidth - elm.offsetLeft) - elm.clientWidth,
        behavior: 'smooth',
      })
      break
    }
  }, [activeTab])

  function handleTabChange(id: string) {
    if (id === activeTab?.id) return
    setSearchParams({ tab: id }, { replace: true })
  }

  function handleKeyDown(e: KeyboardEvent<HTMLAnchorElement>) {
    switch (e.key) {
      case 'ArrowLeft':
      case 'ArrowUp':
      case 'a':
      case 'w':
        const prev = e.currentTarget.previousElementSibling as HTMLAnchorElement | null
        prev?.focus()
        break
      case 'ArrowRight':
      case 'ArrowDown':
      case 'd':
      case 's':
        const next = e.currentTarget.nextElementSibling as HTMLAnchorElement | null
        next?.focus()
        break
      case 'Enter':
        e.currentTarget.click()
        e.currentTarget.focus()
        break
    }
  }

  return (
    <div className="TabView">
      <div className="scroll-wrapper">
        <div className="switch-wrapper">
          <div className="switch" ref={switchRef} role="tablist">
            {tabs.map((tab, i) => (
              <a
                key={tab.id}
                tabIndex={i + 1}
                className={`trigger${tab.id === activeTab?.id ? ' active' : ''}`}
                onClick={() => handleTabChange(tab.id)}
                onKeyDown={handleKeyDown}
                role="tab"
              >
                <FontAwesomeIcon icon={tab.icon} />
                <span>{tab.title}</span>
              </a>
            ))}
          </div>
        </div>
      </div>

      {activeTab && activeTab.element}
    </div>
  )
}

export function Tab({ children }: { children: ReactNode }) {
  return <div role="tabpanel">{children}</div>
}
