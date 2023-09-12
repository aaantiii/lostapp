import useDocumentTitle from '@hooks/useDocumentTitle'
import CenteredContent from '@components/CenteredContent'
import Spacer from '@components/Spacer'
import Button from '@components/Button'
import { useNavigate } from 'react-router-dom'

export default function LoginFailed() {
  const heading = useDocumentTitle('❌ Login fehlgeschlagen ❌')
  const navigate = useNavigate()

  return (
    <main>
      <Spacer size="large" />
      {heading}
      <CenteredContent>
        <p>Während deiner Anmeldung ist ein Fehler aufgetreten. Dies kann unteranderem aus folgenden Gründen passieren:</p>
        <ol>
          <li>Aktiviere Cookies in deinem Browser (sollten standardmäßig aktiviert sein)</li>
          <li>Benutze keinen veralteten Browser (z.B. Internet Explorer). Verwende am besten die neuste Version von Chrome, Firefox oder Safari.</li>
        </ol>
        <Button onClick={() => navigate('/')}>Zur Startseite</Button>
      </CenteredContent>
    </main>
  )
}
