import '@styles/components/Switch.scss'
import * as s from '@radix-ui/react-switch'
import { forwardRef, useId, useRef } from 'react'

type SwitchProps = {
  label?: string
  defaultValue?: boolean
  name?: string
  onChange?: (checked: boolean) => void
  id?: string
  required?: boolean
}

const Switch = forwardRef<HTMLInputElement, SwitchProps>(({ label, id, defaultValue, required, name, onChange }, forwardedRef) => {
  const newId = useId()
  if (!id) id = newId
  const checked = useRef(defaultValue ?? false)

  function handleCheckedChanged(c: boolean) {
    checked.current = c
    onChange?.(c)
  }

  return (
    <fieldset className="Switch">
      {label && (
        <label className="label" htmlFor={id}>
          {label}
        </label>
      )}
      <s.Root className="root" id={id} defaultChecked={defaultValue} onCheckedChange={handleCheckedChanged} required={required}>
        <s.Thumb className="thumb" />
      </s.Root>
      <input type="hidden" name={name} value={checked.current.toString()} ref={forwardedRef} />
    </fieldset>
  )
})

export default Switch
