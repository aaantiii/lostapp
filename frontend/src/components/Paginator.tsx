import '@styles/components/Paginator.scss'
import React, { useEffect, useRef } from 'react'
import { useSearchParams } from 'react-router-dom'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faChevronLeft, faChevronRight } from '@fortawesome/free-solid-svg-icons'
import { Pagination } from '@api/types/base'

export type PaginatorProps = {
  pagination?: Pagination
  showTotalItems?: boolean
  onPageChange?: (page: number) => void
  children: React.ReactNode
}

type PageSwitcherProps = Required<Omit<PaginatorProps, 'children' | 'showTotalItems'>>

export default function Paginator({ pagination, onPageChange, children, showTotalItems = true }: PaginatorProps) {
  const setSearchParams = useSearchParams()[1]
  const paginatorRef = useRef<HTMLDivElement>(null)

  useEffect(() => {
    if (!pagination) return

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
    <div className="Paginator" ref={paginatorRef}>
      {pagination && (
        <>
          {showTotalItems && <span className="total-items">{pagination.totalItems} Ergebnisse</span>}
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
