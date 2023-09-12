import { FormMessageProps } from '@radix-ui/react-form'

namespace FormMessages {
  export const badInput: FormMessageProps = {
    id: 'badInput',
    match: 'badInput',
    children: 'Ungültiger Wert.',
  }

  export const valueMissing: FormMessageProps = {
    id: 'valueMissing',
    match: (value: string) => value === '',
    children: 'Eingabe erforderlich.',
  }

  export const needInteger: FormMessageProps = {
    id: 'isInteger',
    match: (value: string) => !Number.isInteger(Number(value)),
    children: 'Wert muss eine ganze Zahl sein.',
  }

  export const minMaxNumber = (min: number, max?: number): FormMessageProps => ({
    id: 'minMax',
    match: (value: string) => {
      if (value === '') return false

      const num = Number(value)
      if (Number.isNaN(num)) return false
      return num < min || (max !== undefined && num > max)
    },
    children: `Wert muss zwischen ${min} und ${max} liegen.`,
  })

  export const minMaxLength = (min: number, max?: number): FormMessageProps => ({
    id: 'minMaxText',
    match: (value: string) => value.length !== 0 && (value.length < min || (max !== undefined && value.length > max)),
    children: `Wert muss zwischen ${min} und ${max} Zeichen lang sein.`,
  })

  export const fixedLength = (length: number): FormMessageProps => ({
    id: 'fixedLength',
    match: (value: string) => value.length !== 0 && value.length !== length,
    children: `Wert muss ${length} Zeichen lang sein.`,
  })
}

export default FormMessages
