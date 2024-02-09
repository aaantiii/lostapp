import Button from '@components/Button'
import Grid from '@components/Grid'
import Switch from '@components/Switch'
import { useAuth } from '@context/authContext'
import { useTheme } from '@context/themeContext'
import { faArrowsToCircle } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import useTitle from '@hooks/useTitle'

export default function Account() {
  const { theme, toggleTheme } = useTheme()
  const { logout } = useAuth()
  const title = useTitle('Mein Account')

  return (
    <main>
      {title}
      <Grid mode="autofill">
        <Button onClick={logout}>
          <FontAwesomeIcon icon={faArrowsToCircle} /> Abmelden
        </Button>
      </Grid>
      <section>
        <h2>Einstellungen</h2>
        <Grid size="large">
          <Switch label="Dark Mode" defaultValue={theme === 'dark'} onChange={toggleTheme} />
        </Grid>
      </section>
    </main>
  )
}
