import '../scss/components/Select.scss'
import * as s from '@radix-ui/react-select'
import { faAngleDown, faCheck, faChevronDown, faChevronUp } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { forwardRef, useState } from 'react'

interface SelectProps {
  optionGroups: SelectOptionGroup[]
  onChange?: (value: string) => void
  defaultValue?: string
  value?: string
  id?: string
  placeholder?: string
}

export interface SelectOptionGroup {
  title?: string
  options: SelectOption[]
}

export interface SelectOption {
  value: string
  displayText: string
}

export const selectOptionNone: SelectOption = { value: '', displayText: '' }

export default function Select({ onChange, defaultValue, value, placeholder, optionGroups, id }: SelectProps) {
  return (
    <div className="Select">
      <s.Root onValueChange={onChange} defaultValue={defaultValue} value={value}>
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
                    <s.Item value={value} key={value} className="SelectItem">
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

export const SelectFormWrapper = forwardRef<HTMLInputElement, SelectProps & { name?: string }>((props, ref) => {
  const [value, setValue] = useState(props.defaultValue ?? '')

  function handleChange(value: string) {
    setValue(value)
    props.onChange?.(value)
  }

  return (
    <div className="SelectFormWrapper">
      <input name={props.name} ref={ref} type="hidden" value={value} />
      <Select
        optionGroups={props.optionGroups}
        defaultValue={props.defaultValue}
        value={value}
        id={props.id}
        placeholder={props.placeholder}
        onChange={handleChange}
      />
    </div>
  )
})
