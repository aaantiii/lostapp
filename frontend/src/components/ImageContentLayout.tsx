import '@styles/components/ImageContentLayout.scss'

type ImageContentLayoutProps = {
  image: {
    url: string
    title: string
  }
  children: React.ReactNode
}

export default function ImageContentLayout({ image, children }: ImageContentLayoutProps) {
  return (
    <div className="ImageContentLayout">
      <img src={image.url} alt={image.title} loading="eager" />
      <div className="content">{children}</div>
    </div>
  )
}
