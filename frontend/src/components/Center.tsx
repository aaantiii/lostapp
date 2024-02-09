import '@styles/components/Center.scss'

type CenterProps = {
  children: React.ReactNode
}

export default function Center({ children }: CenterProps) {
  return <div className="Center">{children}</div>
}
