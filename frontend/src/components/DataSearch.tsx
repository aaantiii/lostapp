import '../scss/components/DataSearch.scss'
import { useEffect, useState } from 'react'
import Input from './Input'
import useDebouncedValue from '../hooks/useDebouncedValue'

interface DataSearchProps {
  data: any[]
  searchKeys: string[]
  title?: string
  children: (filteredData?: any[]) => JSX.Element
}

export default function DataSearch({ data, searchKeys, title, children }: DataSearchProps) {
  const [result, setResult] = useState<(typeof data)[]>()
  const [searchValue, setSearchValue] = useDebouncedValue('', 150)

  useEffect(() => {
    if (!data) return
    if (!searchValue) return setResult(data)

    for (const key of searchKeys) {
      const res = data.filter((item) => item[key].toLowerCase().includes(searchValue))
      if (res.length > 0) {
        setResult(res)
        return
      }
    }

    setResult(undefined)
  }, [data, searchValue])

  return (
    <div className="DataSearch">
      <div className="header">
        <h3>{title}</h3>
        <Input
          disabled={!data || data.length === 0}
          type="search"
          placeholder="Suchen"
          onChange={(value) => setSearchValue(value.trim().toLowerCase())}
        />
      </div>
      {<div className="content">{children(result)}</div>}
    </div>
  )
}
