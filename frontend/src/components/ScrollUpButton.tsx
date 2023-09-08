import '../scss/components/ScrollUpButton.scss'
import { faChevronUp } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { useEffect, useRef } from 'react'

export default function ScrollUpButton() {
  const buttonRef = useRef<HTMLAnchorElement>(null)

  useEffect(() => {
    handleScroll()

    window.addEventListener('scroll', handleScroll)
    return () => window.removeEventListener('scroll', handleScroll)
  }, [])

  function handleScroll() {
    if (!buttonRef.current) return

    if (window.scrollY > window.innerHeight / 2) {
      buttonRef.current.classList.add('visible')
    } else {
      buttonRef.current.classList.remove('visible')
    }
  }

  return (
    <a onClick={() => window.scrollTo(0, 0)} ref={buttonRef} className="ScrollUpButton">
      <FontAwesomeIcon icon={faChevronUp} />
    </a>
  )
}
