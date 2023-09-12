import '@styles/components/CenteredContent.scss'

interface CenteredContentProps {
  children?: JSX.Element | JSX.Element[]
}

export default function CenteredContent({ children }: CenteredContentProps) {
  return (
    <div className="CenteredContent">
      <div className="content">{children}</div>
    </div>
  )
}
