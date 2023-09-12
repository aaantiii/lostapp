import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import '@styles/components/Input.scss'
import { forwardRef } from 'react'
import { faSearch } from '@fortawesome/free-solid-svg-icons'

type InputProps = Omit<React.InputHTMLAttributes<HTMLInputElement>, 'onChange' | 'title'> & { onChange?: (value: string) => void }

const Input = forwardRef<HTMLInputElement, InputProps>((props, ref) => {
  return (
    <div className="Input">
      <input
        ref={ref}
        {...props}
        type={props.type ?? 'text'}
        onChange={(e) => props.onChange?.(e.target.value)}
        onWheel={(e) => {
          if (props.type === 'number') {
            e.currentTarget.blur()
          }
        }}
      />
      <div className="icon">{props.type === 'search' && <FontAwesomeIcon icon={faSearch} />}</div>
    </div>
  )
})

export default Input
