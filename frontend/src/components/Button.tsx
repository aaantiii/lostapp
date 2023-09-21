import '@styles/components/Button.scss'
import { useNavigate } from 'react-router-dom'
import { MouseEventHandler, forwardRef, useCallback } from 'react'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faSpinner } from '@fortawesome/free-solid-svg-icons'

interface ButtonProps {
  children: string | JSX.Element
  isLoading?: boolean
  to?: string | number
  disabled?: boolean
  onClick?: MouseEventHandler<HTMLButtonElement>
  type?: 'button' | 'submit'
  className?: string
}

const Button = forwardRef<HTMLButtonElement, ButtonProps>(({ children, isLoading, to, disabled, onClick, type = 'button', className }, ref) => {
  const navigate = useNavigate()

  function handleClick(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) {
    if (typeof onClick === 'function') onClick(e)
    else if (typeof to === 'number') navigate(to)
    else if (typeof to === 'string') navigate(to)
  }

  return (
    <button ref={ref} className={`Button${className ? ` ${className}` : ''}`} disabled={disabled || isLoading} onClick={handleClick} type={type}>
      {isLoading ? <FontAwesomeIcon icon={faSpinner} className="loading-spinner" /> : children}
    </button>
  )
})

export default Button
