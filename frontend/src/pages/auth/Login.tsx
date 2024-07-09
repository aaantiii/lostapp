import routes, { buildURI } from '@api/routes'
import Button from '@components/Button'
import { useAuth } from '@context/authContext'
import { useEffect } from 'react'
import { useNavigate } from 'react-router-dom'

export default function Login() {
  const { user } = useAuth()
  const navigate = useNavigate()

  function handleLogin() {
    window.open(buildURI(routes.auth.login), '_self')
  }

  useEffect(() => {
    if (user !== undefined) navigate('/dashboard')
  }, [user])

  return (
    <main>
      <h1>Login</h1>
      <p>Drücke auf "Bei Discord anmelden", um dich mit deinem Discord Account anzumelden.</p>
      <p>
        Info: durch deine Anmeldung erhalten wir keinen Zugriff auf sensible Daten. Wir erhalten lediglich öffentliche Benutzerinfos wie bspw. deinen
        Benutzernamen sowie eine Liste von allen Servern, denen du beigetreten bist. Dabei wird allerdings nur überprüft, ob du auf dem Lost Family
        Discord bist, sowie deine Rollen, um dir die richtigen Berechtigungen zu geben.
      </p>
      <Button onClick={handleLogin}>Bei Discord anmelden.</Button>
    </main>
  )
}
