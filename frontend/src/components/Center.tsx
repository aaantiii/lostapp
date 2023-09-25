import '@styles/components/Center.scss'

interface CenteredContentProps {
  children?: any
}

export default function Center({ children }: CenteredContentProps) {
  return (
    <div className="Center">
      <div className="content">{children}</div>
    </div>
  )
}
