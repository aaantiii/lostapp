import '@styles/components/Spacer.scss'

type SpacerProps = {
  size?: 'small' | 'medium' | 'large'
}

export default function Spacer({ size = 'medium' }: SpacerProps) {
  return <div className={`Spacer ${size}`}></div>
}
