import '@styles/components/DataList.scss'
import { ReactNode } from 'react'

type DataListProps = {
  children: DataListEntry[]
}

export type DataListEntry = {
  label: string
  value: ReactNode
}

export default function DataList({ children }: DataListProps) {
  return (
    <div className="DataList">
      <dl>
        {children.map(({ label, value }) => (
          <div key={label + value}>
            <dt>{label}</dt>
            <dd>{value}</dd>
          </div>
        ))}
      </dl>
    </div>
  )
}
