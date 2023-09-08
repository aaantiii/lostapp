import useDocumentTitle from '../../hooks/useDocumentTitle'

export default function Apply() {
  const heading = useDocumentTitle('Bewerbung')

  return (
    <main>
      {heading}
      <p>Auf dieser Seite kannst du dich für einen Clan der Lost Family bewerben und so Teil von uns werden ❤️</p>
    </main>
  )
}
