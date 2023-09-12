import useDocumentTitle from '@hooks/useDocumentTitle'

export default function Stats() {
  const heading = useDocumentTitle('Statistiken')

  return <main>{heading}</main>
}
