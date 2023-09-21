import { useNavigate } from 'react-router-dom'
import useDocumentTitle from '@hooks/useDocumentTitle'
import Center from '@components/Center'
import Spacer from '@components/Spacer'
import Button from '@components/Button'
import Content from '@components/Content'

export default function LoginFailed() {
  const heading = useDocumentTitle('❌ Login fehlgeschlagen ❌')
  const navigate = useNavigate()

  return (
    <main>
      <Spacer size="large" />
      <Content>
        {heading}
        <Center>
          <p>Während deiner Anmeldung ist ein Fehler aufgetreten. Dies kann unteranderem aus folgenden Gründen passieren:</p>
          <ol>
            <li>Aktiviere Cookies in deinem Browser (sollten standardmäßig aktiviert sein)</li>
            <li>
              Benutze keinen veralteten Browser (z.B. Internet Explorer). Verwende am besten die neustes Version von Chrome, da die App dafür
              optimiert ist.
            </li>
          </ol>
          <Button onClick={() => navigate('/')}>Zur Startseite</Button>
        </Center>
      </Content>
    </main>
  )
}
