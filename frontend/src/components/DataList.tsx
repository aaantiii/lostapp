interface DataListProps {
  data: { title: string; value: string | number }[]
}

export default function DataList({ data }: DataListProps) {
  return (
    <dl>
      {data.map(({ title, value }) => (
        <>
          <dt key={title}>{title}</dt>
          <dd key={`${title}value`}>{value}</dd>
        </>
      ))}
    </dl>
  )
}
