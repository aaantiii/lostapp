import '@styles/components/Content.scss'

interface ContentProps {
  children: any
}

export default function Content({ children }: ContentProps) {
  return <div className="Content">{children}</div>
}
