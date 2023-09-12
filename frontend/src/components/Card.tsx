import '@styles/components/Card.scss'
import { useRef, useEffect, CSSProperties } from 'react'

interface CardProps {
  title: string
  description?: string | string[]
  thumbnail?: JSX.Element
  fields?: {
    title?: string
    value: string | number
    style?: CSSProperties
    key: string
  }[]
  buttons?: JSX.Element[]
  key: string | number
}

interface CardListProps {
  children?: JSX.Element | JSX.Element[]
  flexDirection?: 'row' | 'column'
}

export function Card({ title, description, thumbnail, fields, buttons }: CardProps) {
  return (
    <div className="Card">
      <div className="border"></div>
      <div className="content">
        <div className="head">
          <div>
            <h4 className="title">{title}</h4>
            {description && (
              <div className="description">
                {Array.isArray(description) ? description.map((line, i) => <span key={i}>{line}</span>) : <span>{description}</span>}
              </div>
            )}
          </div>
          {thumbnail && <div className="thumbnail">{thumbnail}</div>}
        </div>
        {fields && (
          <ul className="fields">
            {fields.map((field) => (
              <li key={field.key} className="field">
                {field.title && `${field.title}: `}
                <span style={field.style}>{field.value}</span>
              </li>
            ))}
          </ul>
        )}
        {buttons && <div className="buttons">{buttons}</div>}
      </div>
    </div>
  )
}

export function CardList({ children, flexDirection }: CardListProps) {
  const cardListRef = useRef<HTMLDivElement>(null)

  useEffect(() => {
    if (!cardListRef.current) return

    const handleMouseMove = (e: MouseEvent) => {
      if (!cardListRef.current?.children) return

      for (const cardElm of cardListRef.current.children) {
        const card = cardElm as HTMLDivElement

        const rect = card.getBoundingClientRect()
        const x = e.clientX - rect.left
        const y = e.clientY - rect.top

        card.style.setProperty('--mouse-x', `${x}px`)
        card.style.setProperty('--mouse-y', `${y}px`)
      }
    }

    cardListRef.current.addEventListener('mousemove', handleMouseMove)
    return () => cardListRef.current?.removeEventListener('mousemove', handleMouseMove)
  }, [cardListRef.current])

  return (
    <div ref={cardListRef} className={`CardList ${flexDirection ?? 'row'}`}>
      {children}
    </div>
  )
}
