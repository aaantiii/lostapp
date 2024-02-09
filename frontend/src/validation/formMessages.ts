import { dateFormatter } from '@/utils/intlFormatter'
import { FormMessageProps } from '@radix-ui/react-form'

namespace FormMessages {
  function prepareTextValue(value: string): string {
    return value.trim().replaceAll(/\s\s+/g, ' ')
  }

  export const badInput: FormMessageProps = {
    id: 'badInput',
    match: 'badInput',
    children: 'Ungültiger Wert.',
  }

  export const required: FormMessageProps = {
    id: 'required',
    match: (value: string) => value.length === 0,
    children: 'Eingabe erforderlich.',
  }

  export const needInteger: FormMessageProps = {
    id: 'needInteger',
    match: (value: string) => !Number.isInteger(Number(value)),
    children: 'Wert muss eine ganze Zahl sein.',
  }

  export const minMaxNumber = (min: number, max?: number): FormMessageProps => ({
    id: 'minMaxNumber',
    match: (value: string) => {
      if (value === '') return false

      const num = Number(value)
      if (Number.isNaN(num)) return false
      return num < min || (max !== undefined && num > max)
    },
    children: `Wert muss zwischen ${min} und ${max} liegen.`,
  })

  export const minMaxLength = (min: number, max?: number): FormMessageProps => ({
    id: 'minMaxLength',
    match: (value: string) => value.length !== 0 && (value.length < min || (max !== undefined && value.length > max)),
    children: `Wert muss zwischen ${min} und ${max} Zeichen lang sein.`,
  })

  export const fixedLength = (length: number): FormMessageProps => ({
    id: 'fixedLength',
    match: (value: string) => {
      const v = prepareTextValue(value)
      return v.length !== length && v.length !== 0
    },
    children: `Wert muss ${length} Zeichen lang sein.`,
  })

  export const maxDate = (maxDate: Date): FormMessageProps => ({
    id: 'dateSmallerThan',
    match: (value: string) => {
      if (value === '') return false

      const dateValue = new Date(value)
      const dateValueTruncated = new Date(dateValue.getFullYear(), dateValue.getMonth(), dateValue.getDate())
      const maxDateTruncated = new Date(maxDate.getFullYear(), maxDate.getMonth(), maxDate.getDate())

      return dateValueTruncated > maxDateTruncated
    },
    children: `Datum darf höchstens ${dateFormatter.format(maxDate)} sein.`,
  })

  export const spacesForbidden: FormMessageProps = {
    id: 'spacesForbidden',
    match: (value: string) => value.includes(' '),
    children: 'Leerzeichen sind nicht erlaubt.',
  }
}

export default FormMessages
