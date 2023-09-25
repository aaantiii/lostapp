import '@styles/components/Select.scss'
import * as s from '@radix-ui/react-select'
import { faAngleDown, faCheck, faChevronDown, faChevronUp } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { forwardRef, useState } from 'react'
import Input from './Input'

interface SelectProps {
  optionGroups: SelectOptionGroup[]
  onChange?: (value: string) => void
  defaultValue?: string
  value?: string
  id?: string
  placeholder?: string
  label?: string
}

export interface SelectOptionGroup {
  title?: string
  options: SelectOption[]
}

export interface SelectOption {
  value: string
  realValue?: any
  displayText: string
}

const selectOptionValueOther = 'other'
export const selectOptionOther = (displayText: string): SelectOption => ({ value: selectOptionValueOther, displayText })

export default function Select({ onChange, defaultValue, label, value, placeholder, optionGroups, id }: SelectProps) {
  return (
    <div className="Select">
      {label && (
        <label htmlFor={id} className="label">
          {label}
        </label>
      )}
      <s.Root onValueChange={onChange} defaultValue={defaultValue?.toString()} value={value}>
        <s.Trigger className="SelectTrigger" id={id}>
          <s.Value placeholder={placeholder} />
          <s.Icon className="SelectIcon">
            <FontAwesomeIcon icon={faAngleDown} />
          </s.Icon>
        </s.Trigger>
        <s.Portal>
          <s.Content className="SelectContent">
            <s.ScrollUpButton className="SelectScrollButton">
              <FontAwesomeIcon icon={faChevronUp} />
            </s.ScrollUpButton>
            <s.Viewport className="SelectViewport">
              {optionGroups.map(({ title, options }, i) => [
                <s.Group key={title ?? i}>
                  <s.Label className="SelectLabel">{title ?? 'Auswählen'}</s.Label>
                  {options.map(({ value, displayText }) => (
                    <s.Item value={value.toString()} key={displayText} className="SelectItem">
                      <s.ItemText>{displayText}</s.ItemText>
                      <s.ItemIndicator className="SelectItemIndicator">
                        <FontAwesomeIcon icon={faCheck} />
                      </s.ItemIndicator>
                    </s.Item>
                  ))}
                </s.Group>,
                i < optionGroups.length - 1 ? <s.Separator className="SelectSeperator" key={`seperator${i}`} /> : null,
              ])}
            </s.Viewport>
            <s.ScrollDownButton className="SelectScrollButton">
              <FontAwesomeIcon icon={faChevronDown} />
            </s.ScrollDownButton>
          </s.Content>
        </s.Portal>
      </s.Root>
    </div>
  )
}

type SelectFormWrapperProps = SelectProps & {
  name?: string
  type?: string
}

export const SelectFormWrapper = forwardRef<HTMLInputElement, SelectFormWrapperProps>((props, ref) => {
  const [currentValue, setCurrentValue] = useState(props.defaultValue ?? '')

  function handleChange(value: string) {
    setCurrentValue(value)
    props.onChange?.(value)
  }

  return (
    <div className="SelectFormWrapper">
      <Select
        optionGroups={props.optionGroups}
        defaultValue={props.defaultValue}
        id={props.id}
        placeholder={props.placeholder}
        onChange={handleChange}
      />
      <Input name={props.name} ref={ref} type="hidden" value={currentValue} />
    </div>
  )
})
