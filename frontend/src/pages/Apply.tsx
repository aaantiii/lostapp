import Button from '@components/Button'
import Center from '@components/Center'
import Content from '@components/Content'
import Link from '@components/Link'
import Spacer from '@components/Spacer'
import useDocumentTitle from '@hooks/useDocumentTitle'

export default function ApplyDiscord() {
  const heading = useDocumentTitle('Bewerbung')

  return (
    <main>
      <Spacer size="large" />
      <Content>
        <hgroup>
          {heading}
          <h2>Werde ein Teil von uns!</h2>
        </hgroup>
        <Center>
          <p>
            Du bist ein aktiver Clasher auf der Suche nach Gleichgesinnten? Wir sind ständig auf der Suche nach talentierten Spielern, die gerne in
            einer freundlichen und kooperativen Umgebung clashen.
          </p>
          <p>
            Wenn du Lust hast, dich uns anzuschließen und Teil unserer Clash of Clans-Familie zu werden, dann warte nicht länger – lass uns noch heute
            deine Bewerbung über unseren{' '}
            <Link href="https://discord.gg/XzFSKkBAEB" newWindow>
              Bewerber Discord
            </Link>{' '}
            zukommen!
          </p>
          <Button onClick={() => window.open('https://discord.gg/XzFSKkBAEB', '_blank')}>Jetzt bewerben</Button>
        </Center>
      </Content>
    </main>
  )
}
