import useTitle from '@hooks/useTitle'

export default function Imprint() {
  const title = useTitle('Impressum')

  return (
    <main>
      {title}
      <p>Deine Mama</p>
    </main>
  )
}
