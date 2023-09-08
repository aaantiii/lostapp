import useDocumentTitle from '../hooks/useDocumentTitle'
import useNotImplemented from '../hooks/useNotImplemented'

export default function RequestFeature() {
  useDocumentTitle('Feature vorschlagen')
  useNotImplemented()

  return (
    <main className="p-request-feature">
      <h1>Feature vorschlagen</h1>
    </main>
  )
}
