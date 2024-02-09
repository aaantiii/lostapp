import '@styles/components/Table.scss'
import Paginator, { PaginatorProps } from './Paginator'
import { ReactNode, useEffect, useRef } from 'react'
import { useSearchParams } from 'react-router-dom'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faChevronDown, faChevronUp } from '@fortawesome/free-solid-svg-icons'

type TableProps = Pick<PaginatorProps, 'pagination' | 'onPageChange' | 'showTotalItems'> & {
  header: ReactNode[]
  children: ReactNode
  className?: string
}

type TableColumnProps = {
  children: ReactNode[]
}

export default function Table({ children, className = '', header, ...paginatorProps }: TableProps) {
  const wrapperRef = useRef<HTMLDivElement>(null)

  useEffect(() => {
    wrapperRef.current?.scrollIntoView({
      behavior: 'smooth',
    })
  }, [paginatorProps])

  return (
    <div className={`Table ${className}`}>
      <Paginator {...paginatorProps}>
        <div className="scrollable-wrapper" ref={wrapperRef}>
          <table>
            <thead>
              <tr>
                {header.map((column, i) => (
                  <th key={i}>{column}</th>
                ))}
              </tr>
            </thead>
            <tbody>{children}</tbody>
          </table>
        </div>
      </Paginator>
    </div>
  )
}

export function TableRow({ children }: TableColumnProps) {
  return (
    <tr>
      {children.map((column, i) => (
        <td key={i}>{column}</td>
      ))}
    </tr>
  )
}

type TableOrderProps = {
  children: ReactNode
  queryName: string
  defaultValue: 'asc' | 'desc'
}

export function TableOrder({ children, queryName, defaultValue }: TableOrderProps) {
  const [searchParams, setSearchParams] = useSearchParams({
    [queryName]: defaultValue,
  })
  const currentOrder = searchParams.get(queryName)

  function handleOrderChange() {
    const newOrder = currentOrder === 'asc' ? 'desc' : 'asc'
    setSearchParams(
      (prev) => {
        prev.set(queryName, newOrder)
        prev.set('page', '1')
        return prev
      },
      { replace: true }
    )
  }

  return (
    <a role="button" className="TableOrder" onClick={handleOrderChange}>
      {children}
      {currentOrder === 'asc' ? <FontAwesomeIcon icon={faChevronUp} /> : <FontAwesomeIcon icon={faChevronDown} />}
    </a>
  )
}
