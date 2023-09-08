import { useNavigate } from 'react-router-dom'
import '../scss/components/Table.scss'
import Paginator, { PaginatorProps } from './Paginator'
import { urlEncodeTag } from '../fmt/cocFormatter'
import { numberFormatter, dateTimeFormatter } from '../fmt/formatters'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faUpRightFromSquare } from '@fortawesome/free-solid-svg-icons'

interface TableProps extends PaginatorProps {
  rowCountColumn?: boolean
  data: any[]
  columns: TableColumn[]
}

export interface TableColumn {
  prop: string
  heading: string
  link?: string
  linkIdProp?: string
  type?: 'number' | 'date' | 'string'
}

export function Table({ pagination, data, columns, rowCountColumn, onPageChange }: TableProps) {
  const navigate = useNavigate()

  function formatCellData(row: any, col: TableColumn): string | JSX.Element {
    let value = row[col.prop]

    if (col.type === 'number') {
      value = numberFormatter.format(value)
    } else if (col.type === 'date') {
      value = dateTimeFormatter.format(value)
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
    <div className="Table">
      <Paginator pagination={pagination} onPageChange={onPageChange} />
      <div className="scrollable-wrapper">
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
