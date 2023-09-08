import '../scss/components/Spacer.scss'

interface SpacerProps {
  size: 'small' | 'medium' | 'large'
}

export default function Spacer({ size }: SpacerProps) {
  return <div className={`Spacer ${size}`}></div>
}
