import { replacePlaceholders } from '@util/jsonPlaceholders'
import { FormMessageProps } from '@radix-ui/react-form'
import { useTranslation } from 'react-i18next'
import { useLanguage } from '@context/languageContext'

export default function useFormMessages() {
  const [t] = useTranslation('components/form', { keyPrefix: 'messages' })
  const { formatDate, formatNumber } = useLanguage()

  const valueMissing: FormMessageProps = {
    id: 'valueMissing',
    match: 'valueMissing',
    children: t('valueMissing'),
  }

  const badInput: FormMessageProps = {
    id: 'badInput',
    match: 'badInput',
    children: t('badInput'),
  }

  const required: FormMessageProps = {
    id: 'required',
    match: (value: string) => value.length === 0,
    children: t('required'),
  }

  const needInteger: FormMessageProps = {
    id: 'needInteger',
    match: (value: string) => !Number.isInteger(Number(value)),
    children: t('needInteger'),
  }

  const notZero: FormMessageProps = {
    id: 'notZero',
    match: (value: string) => Number(value) === 0,
    children: t('notZero'),
  }

  const minMaxNumber = (min: number, max: number): FormMessageProps => ({
    id: 'minMaxNumber',
    match: (value: string) => {
      if (value === '') return false

      const num = Number(value)
      if (Number.isNaN(num)) return false
      return num < min || num > max
    },
    children: replacePlaceholders(t('minMaxNumber'), { min: formatNumber(min), max: formatNumber(max) }),
  })

  const minMaxLength = (min: number, max: number): FormMessageProps => ({
    id: 'minMaxLength',
    match: (value: string) => value.length !== 0 && (value.length < min || value.length > max),
    children: replacePlaceholders(t('minMaxLength'), { min: formatNumber(min), max: formatNumber(max) }),
  })

  const minDate = (min: Date): FormMessageProps => ({
    id: 'minDate',
    match: (value: string) => {
      if (value === '') return false
      console.log(value, new Date(value))
      const valueTruncated = truncateToDay(new Date(value))
      const minTruncated = truncateToDay(min)

      return valueTruncated < minTruncated
    },
    children: replacePlaceholders(t('minDate'), { min: formatDate(min) }),
  })

  const maxDate = (max: Date): FormMessageProps => ({
    id: 'maxDate',
    match: (value: string) => {
      if (value === '') return false

      const valueTruncated = truncateToDay(new Date(value))
      const maxTruncated = truncateToDay(max)

      return valueTruncated > maxTruncated
    },
    children: t('maxDate', { maxDate: max }),
  })

  const dateRange = (min: Date, max: Date): FormMessageProps => ({
    id: 'dateRange',
    match: (value: string) => {
      if (value === '') return false

      const valueTruncated = truncateToDay(new Date(value))
      const minTruncated = truncateToDay(min)
      const maxTruncated = truncateToDay(max)

      return valueTruncated < minTruncated || valueTruncated > maxTruncated
    },
    children: t('dateRange', { min: formatDate(min), max: formatDate(max) }),
  })

  const oneOf = (options: string[]): FormMessageProps => ({
    id: 'oneOf',
    match: (value: string) => !options.includes(value),
    children: t('oneOf', { options }),
  })

  return {
    valueMissing,
    badInput,
    required,
    needInteger,
    notZero,
    minMaxNumber,
    minMaxLength,
    minDate,
    maxDate,
    dateRange,
    oneOf,
  }
}

function truncateToDay(date: Date) {
  return new Date(date.getFullYear(), date.getMonth(), date.getDate())
}
