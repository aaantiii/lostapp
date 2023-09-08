import '../scss/components/Content.scss'

interface ContentProps {
  children: JSX.Element[] | JSX.Element
}

export default function Content({ children }: ContentProps) {
  return <div className="Content">{children}</div>
}
