import { useEffect, useState } from 'react'
import Button from '@components/Button'
import { useAuth } from '@context/authContext'
import useDashboardNavigate from '@hooks/useDashboardNavigate'
import routes from '@api/routes'
import CenteredContent from '@components/CenteredContent'
import useDocumentTitle from '@hooks/useDocumentTitle'
import Spacer from '@components/Spacer'

const width = 400
const height = 600
const popupWindowArgs = `
  height=${height},
  width=${width},
  top=${(screen.height - height) / 4},
  left=${(screen.width - width) / 2},
  resizable=yes,
  scrollbars=yes,
  toolbar=yes,
  menubar=no,
  location=no,
  directories=no,
  status=yes`

export default function Index() {
  const heading = useDocumentTitle('Anmelden')
  const overviewRedirect = useDashboardNavigate()
  const { discordUser, refreshSession } = useAuth()

  useEffect(() => {
    if (discordUser) overviewRedirect()
  }, [discordUser])

  function handleLogin() {
    window.open(`${import.meta.env.VITE_SERVICE_API}/${routes.auth.login}`, '_self')
  }

  return (
    <main>
      <Spacer size="large" />
      {heading}
      <CenteredContent>
        <p>Drücke unten auf den Button um dich anzumelden. Du wirst dafür auf die Website von Discord weitergeleitet.</p>
        <Button onClick={handleLogin} disabled={discordUser !== undefined}>
          Weiter zu Discord
        </Button>
      </CenteredContent>
    </main>
  )
}
