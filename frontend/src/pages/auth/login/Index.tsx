import { useEffect, useState } from 'react'
import Button from '@components/Button'
import { useAuth } from '@context/authContext'
import useDashboardNavigate from '@hooks/useDashboardNavigate'
import routes from '@api/routes'
import Center from '@components/Center'
import useDocumentTitle from '@hooks/useDocumentTitle'
import Spacer from '@components/Spacer'
import Content from '@components/Content'

export default function Index() {
  const heading = useDocumentTitle('Anmelden')
  const overviewRedirect = useDashboardNavigate()
  const { discordUser } = useAuth()

  useEffect(() => {
    if (discordUser) overviewRedirect()
  }, [discordUser])

  function handleLogin() {
    window.open(`${import.meta.env.VITE_API_URL}/${routes.auth.login}`, '_self')
  }

  return (
    <main>
      <Spacer size="large" />
      <Content>
        {heading}
        <Center>
          <p>Drücke unten auf den Button, um dich anzumelden. Du wirst dafür auf die Website von Discord weitergeleitet.</p>
          <Button onClick={handleLogin} disabled={discordUser !== undefined}>
            Weiter zu Discord
          </Button>
        </Center>
      </Content>
    </main>
  )
}
