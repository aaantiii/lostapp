import { useEffect } from 'react'
import Input from './Input'
import useDebouncedValue from '@hooks/useDebouncedValue'
import { useSearchParams } from 'react-router-dom'

type SearchQueryProps = {
  placeholder?: string
  label?: string
  inputMode?: 'numeric' | 'search'
}

export default function SearchQuery({ placeholder, label, inputMode = 'search' }: SearchQueryProps) {
  const [searchParams, setSearchParams] = useSearchParams()
  const [searchValueDebounced, setSearchValue] = useDebouncedValue(searchParams.get('q') ?? '')

  useEffect(() => {
    if (searchValueDebounced === searchParams.get('q')) return
    setSearchParams(
      (prev) => {
        if (searchValueDebounced === '') {
          prev.delete('q')
          return prev
        }

        prev.set('q', searchValueDebounced)
        prev.set('page', '1')
        return prev
      },
      { replace: true }
    )
  }, [searchValueDebounced])

  return (
    <Input
      type="search"
      placeholder={placeholder}
      label={label}
      defaultValue={searchValueDebounced}
      onChange={(value) => setSearchValue(value.trim())}
      inputMode={inputMode}
    />
  )
}
