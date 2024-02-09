import '@styles/components/Button.scss'
import { faSpinner } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { ButtonHTMLAttributes, forwardRef } from 'react'
import { useNavigate } from 'react-router-dom'

type ButtonProps = ButtonHTMLAttributes<HTMLButtonElement> & {
  children: React.ReactNode
  to?: string | number
  loading?: boolean
}

const Button = forwardRef<HTMLButtonElement, ButtonProps>(({ children, to, loading, type = 'button', className = '', ...buttonProps }, ref) => {
  const navigate = useNavigate()

  function handleClick(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) {
    if (typeof buttonProps.onClick === 'function') buttonProps.onClick(e)
    else if (typeof to === 'number') navigate(to)
    else if (typeof to === 'string') navigate(to)
  }

  return (
    <button {...buttonProps} type={type} className={`Button ${className}`} disabled={buttonProps.disabled || loading} onClick={handleClick} ref={ref}>
      {loading ? <FontAwesomeIcon icon={faSpinner} className="loading-spinner" /> : children}
    </button>
  )
})

export default Button
