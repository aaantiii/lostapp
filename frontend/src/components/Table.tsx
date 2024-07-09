import '@styles/components/Table.scss'
import Paginator, { PaginatorProps } from './Paginator'
import { ReactNode, useRef } from 'react'
import { useSearchParams } from 'react-router-dom'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faChevronDown, faChevronUp } from '@fortawesome/free-solid-svg-icons'

type TableProps = Omit<PaginatorProps, 'children'> & {
  header: ReactNode[]
  children: ReactNode
  className?: string
  rawData?: any
}

type TableColumnProps = {
  children: ReactNode[]
}

export default function Table({ children, className = '', header, rawData, ...paginatorProps }: TableProps) {
  const wrapperRef = useRef<HTMLDivElement>(null)
  const tableRef = useRef<HTMLTableElement>(null)

  function downloadCSV() {
    const table = tableRef.current
    if (!table) return

    const rows = table.querySelectorAll('tr')
    const csv = Array.from(rows)
      .map((row) => {
        const columns = row.querySelectorAll(':scope > *')
        return Array.from(columns)
          .map((column) => column.textContent)
          .join(';')
      })
      .join('\n')

    const blob = new Blob([csv], { type: 'text/csv' })
    const url = URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = 'table.csv'
    a.click()
    URL.revokeObjectURL(url)
  }

  function downloadJSON() {
    const blob = new Blob([JSON.stringify(rawData)], { type: 'application/json' })
    const url = URL.createObjectURL(blob)
    const a = document.createElement('a')
    a.href = url
    a.download = 'table.json'
    a.click()
    URL.revokeObjectURL(url)
  }

  return (
    <div className={`Table ${className}`}>
      <Paginator {...paginatorProps}>
        <div className="scrollable-wrapper" ref={wrapperRef}>
          <table ref={tableRef}>
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
        <div className="download">
          <a onClick={downloadCSV}>CSV</a>
          {rawData && <a onClick={downloadJSON}>JSON</a>}
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
