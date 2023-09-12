import useDocumentTitle from '@hooks/useDocumentTitle'
import useNotImplemented from '@hooks/useNotImplemented'

export default function ReportBug() {
  useDocumentTitle('Fehler melden')
  useNotImplemented()
  return (
    <main className="p-report-bug">
      <h1>Fehler melden</h1>
    </main>
  )
}
