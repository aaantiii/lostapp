import '@styles/components/Button.scss'
import { useNavigate } from 'react-router-dom'
import { MouseEventHandler, forwardRef, useCallback } from 'react'

interface ButtonProps {
  children: string | JSX.Element
  to?: string | number
  disabled?: boolean
  onClick?: MouseEventHandler<HTMLButtonElement>
  type?: 'button' | 'submit'
  className?: string
}

const Button = forwardRef<HTMLButtonElement, ButtonProps>(({ children, to, disabled, onClick, type, className }, ref) => {
  const navigate = useNavigate()
  const handleClick = useCallback(
    (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
      if (typeof onClick == 'function') onClick(e)
      else if (typeof to == 'number') navigate(to)
      else if (typeof to == 'string') navigate(to)
    },
    [onClick, to]
  )

  return (
    <button ref={ref} className={`Button${className ? ` ${className}` : ''}`} disabled={disabled} onClick={handleClick} type={type ?? 'button'}>
      {children}
    </button>
  )
})

export default Button
