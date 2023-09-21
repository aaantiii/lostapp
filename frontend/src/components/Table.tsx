import '@styles/components/Table.scss'
import { useNavigate } from 'react-router-dom'
import Paginator, { PaginatorProps } from './Paginator'
import { urlEncodeTag } from '@fmt/cocFormatter'
import { numberFormatter, dateFormatter } from '@fmt/intlFormatter'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faUpRightFromSquare } from '@fortawesome/free-solid-svg-icons'
import { useEffect, useRef } from 'react'

interface TableProps extends PaginatorProps {
  rowCountColumn?: boolean
  data: any[]
  columns: TableColumn[]
  width?: number
}

export interface TableColumn {
  prop: string
  heading: string
  link?: string
  linkIdProp?: string
  type?: 'number' | 'date' | 'string'
}

export function Table({ pagination, data, columns, rowCountColumn, width, onPageChange }: TableProps) {
  const navigate = useNavigate()
  const wrapperRef = useRef<HTMLDivElement>(null)

  useEffect(() => {
    wrapperRef.current?.scrollTo(0, 0)
  }, [data])

  function formatCellData(row: any, col: TableColumn): string | JSX.Element {
    let value = row[col.prop]

    if (col.type === 'number') {
      value = numberFormatter.format(value)
    } else if (col.type === 'date') {
      value = dateFormatter.format(value)
    }

    if (col.link) {
      const href = col.linkIdProp ? encodeURI(col.link + '/' + urlEncodeTag(row[col.linkIdProp])) : col.link

      return (
        <a onClick={() => navigate(href)}>
          {value}
          <span> </span>
          <FontAwesomeIcon className="icon" icon={faUpRightFromSquare} />{' '}
        </a>
      )
    }

    return value
  }

  return (
    <div className="Table" style={width ? { width } : {}}>
      <Paginator pagination={pagination} onPageChange={onPageChange} />
      <div className="scrollable-wrapper" ref={wrapperRef}>
        <table>
          <thead>
            <tr>
              {rowCountColumn && <th>#</th>}
              {columns.map(({ prop, heading }) => (
                <th key={prop}>{heading}</th>
              ))}
            </tr>
          </thead>
          <tbody>
            {data.map((row, i) => (
              <tr key={i}>
                {rowCountColumn && <td>{pagination ? pagination.page * pagination.pageSize - pagination.pageSize + i + 1 : i + 1}</td>}
                {columns.map((col) => (
                  <td key={col.prop}>{formatCellData(row, col)}</td>
                ))}
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  )
}
