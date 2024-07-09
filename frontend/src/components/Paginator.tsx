import '@styles/components/Paginator.scss'
import React, { useEffect } from 'react'
import { useSearchParams } from 'react-router-dom'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faChevronLeft, faChevronRight } from '@fortawesome/free-solid-svg-icons'
import { Pagination } from '@api/types/base'
import Select from './Select'

export type PaginatorProps = {
  pagination?: Pagination
  limits?: number[]
  showTotalItems?: boolean
  onPageChange?: (page: number) => void
  children: React.ReactNode
}

type PageSwitcherProps = Required<Omit<PaginatorProps, 'children' | 'showTotalItems' | 'limits'>>

export default function Paginator({ pagination, onPageChange, children, showTotalItems = true, limits = [10, 20, 30, 40, 50] }: PaginatorProps) {
  const [searchParams, setSearchParams] = useSearchParams()

  useEffect(() => {
    if (!pagination) return

    setSearchParams(
      (prev) => {
        prev.set('page', pagination.page.toString())
        prev.set('limit', pagination.limit.toString())
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
          <div className="header">
            {showTotalItems && <span className="total-items">{pagination.totalItems} Ergebnisse</span>}
            <Select
              label="Ergebnisse pro Seite"
              defaultValue={searchParams.get('limit') ?? limits[0].toString()}
              disableClear
              onChange={(value) => {
                setSearchParams(
                  (prev) => {
                    prev.set('limit', value!)
                    prev.set('page', '1')
                    return prev
                  },
                  { replace: true }
                )
                onPageChange?.(1)
              }}
              options={limits.map((limit) => ({ value: limit.toString(), label: limit.toString() }))}
            />
          </div>
          {children}
          <PageSwitcher pagination={pagination} onPageChange={handlePageChange} />
        </>
      )}
    </div>
  )
}

function PageSwitcher({ pagination, onPageChange }: PageSwitcherProps) {
  if (pagination.totalPages <= 1) return null // dont render

  return (
    <div className="PageSwitcher">
      <div className="wrapper">
        <button onClick={() => onPageChange(pagination.page - 1)} disabled={pagination.page <= 1}>
          <FontAwesomeIcon icon={faChevronLeft} />
        </button>
        {pagination.navigation.map((page) => (
          <button key={page} onClick={() => onPageChange(page)} disabled={pagination.page === page} className="navigation-element">
            {page}
          </button>
        ))}
        <button onClick={() => onPageChange(pagination.page + 1)} disabled={pagination.page >= pagination.totalPages}>
          <FontAwesomeIcon icon={faChevronRight} />
        </button>
      </div>
    </div>
  )
}
