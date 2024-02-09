import '@styles/components/Card.scss'
import { ReactNode } from 'react'

type CardProps = {
  title: string
  description?: string
  image?: string
  thumbnail?: ReactNode
  fields?: {
    label: string
    value: string
  }[]
  buttons?: React.ReactNode
}

export default function Card({ title, description, image, thumbnail, fields, buttons }: CardProps) {
  return (
    <div className="Card">
      <div className="header">
        {image && (
          <div className="image">
            <img className="image" src={image} alt={title} loading="lazy" />
          </div>
        )}
        <div className="body">
          <div className="content">
            <span className="title" role="heading">
              {title}
            </span>
            {description && <span className="description">{description}</span>}
          </div>
          {thumbnail && (
            <div className="thumbnail">{typeof thumbnail === 'string' ? <img src={thumbnail} alt={title} loading="lazy" /> : thumbnail}</div>
          )}
        </div>
      </div>
      {fields && (
        <div className="fields">
          {fields.map(({ label, value }) => (
            <div className="field" key={label}>
              <span className="label">{label}</span>
              <span className="value">{value}</span>
            </div>
          ))}
        </div>
      )}
      {buttons && <div className="buttons">{buttons}</div>}
    </div>
  )
}
