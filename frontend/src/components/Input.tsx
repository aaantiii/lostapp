import '@styles/components/Input.scss'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { createRef, forwardRef, useEffect, useId, useRef } from 'react'
import { faCalendarDays, faSearch } from '@fortawesome/free-solid-svg-icons'

export type InputProps = Omit<React.InputHTMLAttributes<HTMLInputElement>, 'onChange' | 'title'> & {
  onChange?: (value: string) => void
  label?: string
}

const Input = forwardRef<HTMLInputElement, InputProps>((props, forwardedRef) => {
  const newId = useId()
  const inputRef = useRef<HTMLInputElement>()

  function handleDatePick() {
    inputRef.current?.showPicker()
  }

  useEffect(() => {
    if (props.type === 'date') {
      inputRef.current?.addEventListener('click', handleDatePick)
    }

    return () => {
      inputRef.current?.removeEventListener('click', handleDatePick)
    }
  }, [props.type])

  return (
    <fieldset className="Input">
      {props.label && <label htmlFor={props.id ?? newId}>{props.label}</label>}
      <div className="wrapper">
        <input
          ref={(e) => {
            inputRef.current = e ?? undefined
            if (typeof forwardedRef === 'function') forwardedRef(e)
            else if (forwardedRef) forwardedRef.current = e
          }}
          {...props}
          id={props.id ?? newId}
          type={props.type ?? 'text'}
          onChange={(e) => props.onChange?.(e.target.value)}
          onWheel={(e) => {
            if (props.type === 'number') {
              e.currentTarget.blur()
            }
          }}
        />
        {props.type === 'search' && (
          <div className="icon">
            <FontAwesomeIcon icon={faSearch} />
          </div>
        )}
        {props.type === 'date' && (
          <div className="icon">
            <FontAwesomeIcon icon={faCalendarDays} />
          </div>
        )}
      </div>
    </fieldset>
  )
})

export default Input
