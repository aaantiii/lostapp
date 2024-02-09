import '@styles/components/Select.scss'
import { faCaretDown, faCheck, faXmark } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { MouseEvent, forwardRef, useEffect, useId, useRef, useState } from 'react'

type SelectProps = {
  options: SelectOption[]
  id?: string
  placeholder?: string
  label?: string
  required?: boolean
  name?: string
  defaultValue?: string | string[]
  disableClear?: boolean
  onChange?: (value: string | undefined) => void
  onMultiChange?: (value: string[] | undefined) => void
}

export type SelectOption = {
  label: string
  value: string
  disabled?: boolean
}

const Select = forwardRef<HTMLInputElement, SelectProps>(
  ({ onChange, onMultiChange, label, required, defaultValue, disableClear, name, placeholder, options = [], id }, forwardedRef) => {
    const [isOpen, setIsOpen] = useState(false)
    const [selectedValue, setSelectedValue] = useState(defaultValue)
    const [higlightedIndex, setHighlightedIndex] = useState(0)
    const wrapperRef = useRef<HTMLDivElement>(null)

    const newId = useId()

    useEffect(() => {
      if (isOpen) {
        document.body.style.overflow = 'hidden'
        setHighlightedIndex(0)
      } else {
        document.body.style.overflow = 'auto'
      }

      return () => {
        document.body.style.overflow = 'auto'
      }
    }, [isOpen])

    useEffect(() => {
      function handleKeyDown(e: KeyboardEvent) {
        if (e.target !== wrapperRef.current) return
        document.body.style.overflow = 'hidden'
        switch (e.code) {
          case 'Enter':
          case 'Space':
            setIsOpen((prev) => !prev)
            if (isOpen) selectOption(options[higlightedIndex].value)
            break
          case 'ArrowDown':
          case 'ArrowUp': {
            if (!isOpen) {
              setIsOpen(true)
              break
            }

            const newIndex = e.code === 'ArrowDown' ? higlightedIndex + 1 : higlightedIndex - 1
            if (newIndex >= 0 && newIndex < options.length) setHighlightedIndex(newIndex)
            break
          }
          case 'Escape':
            setIsOpen(false)
            break
        }
      }

      wrapperRef.current?.addEventListener('keydown', handleKeyDown)
      const timeoutId = setTimeout(() => {}, 0)
      return () => {
        wrapperRef.current?.removeEventListener('keydown', handleKeyDown)
        clearTimeout(timeoutId)
      }
    }, [isOpen, higlightedIndex, options])

    function selectOption(value: string) {
      if (Array.isArray(selectedValue)) {
        if (selectedValue.includes(value)) {
          const newOptions = selectedValue.filter((opt) => opt !== value)
          setSelectedValue(newOptions)
          onMultiChange?.(newOptions)
        } else {
          const newOptions = [...selectedValue, value]
          setSelectedValue(newOptions)
          onMultiChange?.(newOptions)
        }
      } else {
        if (value === selectedValue) return
        setSelectedValue(value)
        onChange!(value)
        setIsOpen(false)
      }
    }

    function optionFromValue(value?: string) {
      return options.find((opt) => opt.value === value)
    }

    function clearOptions(e: MouseEvent<HTMLButtonElement>) {
      e.stopPropagation()
      Array.isArray(selectedValue) ? setSelectedValue([]) : setSelectedValue(undefined)
      onChange?.(undefined)
      onMultiChange?.(undefined)
    }

    function isOptionSelected(option: SelectOption) {
      return Array.isArray(selectedValue) ? selectedValue.includes(option.value) : selectedValue === option.value
    }

    function handleLabelClick() {
      wrapperRef.current?.focus()
      wrapperRef.current?.click()
    }

    return (
      <fieldset className="Select">
        {label && (
          <label htmlFor={id ?? newId} className="label" onClick={handleLabelClick}>
            {label}
          </label>
        )}
        <div
          ref={wrapperRef}
          className="wrapper"
          tabIndex={0}
          onClick={() => setIsOpen((prev) => !prev)}
          onBlur={() => setIsOpen(false)}
          role="button"
        >
          <span className="value">
            {Array.isArray(selectedValue)
              ? selectedValue.length === 0
                ? placeholder
                : selectedValue.map((value) => (
                    <button
                      className="option-badge"
                      key={value}
                      onClick={(e) => {
                        e.stopPropagation()
                        selectOption(value)
                      }}
                    >
                      {optionFromValue(value)?.label}
                      <FontAwesomeIcon icon={faXmark} className="remove-button" />
                    </button>
                  ))
              : optionFromValue(selectedValue)?.label ?? placeholder}
          </span>
          {!disableClear && (
            <>
              <button className="clear-button" onClick={clearOptions}>
                <FontAwesomeIcon icon={faXmark} />
              </button>
              <div className="divider"></div>
            </>
          )}
          <div className="caret">
            <FontAwesomeIcon icon={faCaretDown} />
          </div>
          <ul className={`options ${isOpen ? 'show' : ''}`}>
            {options.map((opt, i) => (
              <li
                key={opt.value}
                className={`option ${isOptionSelected(opt) ? 'selected ' : ''}${i === higlightedIndex ? 'highlighted ' : ''}${
                  opt.disabled ? 'disabled' : ''
                }`}
                onClick={(e) => {
                  e.stopPropagation()
                  if (!opt.disabled) selectOption(opt.value)
                }}
                onMouseEnter={() => setHighlightedIndex(i)}
              >
                <FontAwesomeIcon icon={faCheck} className="check-mark" />
                {opt.label}
              </li>
            ))}
          </ul>
        </div>
        <input type="hidden" id={id ?? newId} name={name} value={selectedValue ?? ''} ref={forwardedRef} required={required} />
      </fieldset>
    )
  }
)

export default Select
