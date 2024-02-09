import '@styles/components/Grid.scss'
import { ReactNode } from 'react'

type GridProps = {
  children: ReactNode
  mode?: 'autofit' | 'autofill'
  size?: 'medium' | 'large'
}

export default function Grid({ children, mode = 'autofit', size = 'medium' }: GridProps) {
  return <div className={`Grid ${mode} ${size}`}>{children}</div>
}
