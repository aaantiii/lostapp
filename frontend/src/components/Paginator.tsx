import '../scss/components/Paginator.scss'
import { useEffect, useState } from 'react'
import Select, { SelectOption, SelectOptionGroup } from './Select'
import { Pagination } from '../api/types/pagination'

export interface PaginatorProps {
  pagination?: Pagination
  onPageChange: (page: number) => void
}

export default function Paginator({ pagination, onPageChange }: PaginatorProps) {
  const [optionGroup, setOptionGroup] = useState<SelectOptionGroup>({ options: [] })

  useEffect(() => {
    if (!pagination) return

    const options: SelectOption[] = []
    for (let i = 1; i <= pagination.totalPages; i++) {
      options.push({ value: i.toString(), displayText: `Seite ${i}` })
    }
    setOptionGroup({ options, title: 'Seite wählen' })
  }, [pagination])

  return (
    <div className="Paginator">
      {pagination && (
        <>
          <span className="total-items">
            {pagination.totalItems} Ergebniss{pagination.totalItems > 1 ? 'e' : ''}
          </span>
          {pagination.totalPages > 1 && (
            <Select onChange={(value) => onPageChange(Number(value))} value={pagination.page.toString()} optionGroups={[optionGroup]} />
          )}
        </>
      )}
    </div>
  )
}
