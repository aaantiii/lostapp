import useDocumentTitle from '../../hooks/useDocumentTitle'

export default function Overview() {
  const heading = useDocumentTitle('Übersicht')

  return <main>{heading}</main>
}
