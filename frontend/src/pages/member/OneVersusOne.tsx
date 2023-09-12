import useDocumentTitle from '@hooks/useDocumentTitle'
import useNotImplemented from '@hooks/useNotImplemented'

export default function OneVersusOne() {
  const heading = useDocumentTitle('1 vs 1')
  useNotImplemented()

  return (
    <main>
      {heading}
      <p>
        <span className="bold">Info: </span>
        Momentan werden nur Statistiken der gesamten Zeit angezeigt. Das Filtern nach Zeiträumen wird in Zukunft möglich sein.
      </p>
    </main>
  )
}
