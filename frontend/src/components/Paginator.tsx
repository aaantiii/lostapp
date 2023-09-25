import '@styles/components/Paginator.scss'
import { useEffect, useState } from 'react'
import Select, { SelectOption, SelectOptionGroup } from './Select'
import { Pagination } from '@api/types/pagination'
import { useSearchParams } from 'react-router-dom'
import usePageSize from '@hooks/usePageSize'

export interface PaginatorProps {
  pagination?: Pagination
  pageSize?: number
  onPageChange?: (page: number) => void
}

export default function Paginator({ pagination, pageSize, onPageChange }: PaginatorProps) {
  const [optionGroup, setOptionGroup] = useState<SelectOptionGroup>({ options: [] })
  const defaultPageSize = usePageSize(12, 20)
  const [_, setSearchParams] = useSearchParams()

  useEffect(() => {
    setSearchParams(
      (prev) => {
        prev.set('page', prev.get('page') ?? '1')
        prev.set('pageSize', prev.get('pageSize') ?? pageSize?.toString() ?? defaultPageSize.toString())
        return prev
      },
      { replace: true }
    )
  }, [defaultPageSize])

  useEffect(() => {
    if (!pagination) return

    const options: SelectOption[] = []
    for (let i = 1; i <= pagination.totalPages; i++) {
      options.push({ value: i.toString(), displayText: `Seite ${i}` })
    }

    setOptionGroup({ options, title: 'Seite wählen' })
    setSearchParams(
      (prev) => {
        prev.set('page', pagination.page.toString())
        prev.set('pageSize', pagination.pageSize.toString())
        return prev
      },
      { replace: true }
    )
  }, [pagination])

  function handlePageChange(page: number) {
    setSearchParams(
      (prev) => {
        prev.set('page', page.toString())
        return prev
      },
      { replace: true }
    )
    onPageChange?.(page)
  }

  return (
    <div className="Paginator">
      {pagination && (
        <>
          <span className="total-items">
            {pagination.totalItems} Ergebniss{pagination.totalItems > 1 ? 'e' : ''}
          </span>
          {pagination.totalPages > 1 && (
            <Select onChange={(page) => handlePageChange(Number(page))} value={pagination.page.toString()} optionGroups={[optionGroup]} />
          )}
        </>
      )}
    </div>
  )
}
