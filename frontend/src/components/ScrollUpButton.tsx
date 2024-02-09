import '@styles/components/ScrollUpButton.scss'
import { faChevronUp } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { useEffect, useRef, useState } from 'react'

export default function ScrollUpButton() {
  const buttonRef = useRef<HTMLAnchorElement>(null)
  const [isVisible, setIsVisible] = useState(false)

  useEffect(() => {
    handleScroll()
    window.addEventListener('scroll', handleScroll)
    return () => window.removeEventListener('scroll', handleScroll)
  }, [])

  function handleScroll() {
    if (!buttonRef.current) return

    if (window.scrollY > window.innerHeight * 0.75) {
      buttonRef.current.classList.add('visible')
      setIsVisible(true)
    } else {
      buttonRef.current.classList.remove('visible')
      setIsVisible(false)
    }
  }

  return (
    <a onClick={() => window.scrollTo(0, 0)} ref={buttonRef} className="ScrollUpButton">
      <FontAwesomeIcon icon={faChevronUp} />
    </a>
  )
}
