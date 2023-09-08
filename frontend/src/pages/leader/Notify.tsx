import TextArea from '../../components/TextArea'
import useDocumentTitle from '../../hooks/useDocumentTitle'

export default function Notify() {
  const heading = useDocumentTitle('Nachricht senden')

  return (
    <main>
      {heading}
      <TextArea />
    </main>
  )
}
