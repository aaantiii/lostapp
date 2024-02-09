import '@styles/components/Layout.scss'

type LayoutProps = { children: any }

export default function Layout({ children }: LayoutProps) {
  return <div className="Layout">{children}</div>
}
