import { boolToAscii } from '@/utils/worder'
import routes from '@api/routes'
import { ApiError, Clan } from '@api/types/models'
import Button from '@components/Button'
import Center from '@components/Center'
import DataList, { DataListItem } from '@components/DataList'
import Grid from '@components/Grid'
import QueryState from '@components/QueryState'
import Switch from '@components/Switch'
import { useAuth } from '@context/authContext'
import { useTheme } from '@context/themeContext'
import { faRightFromBracket } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import useTitle from '@hooks/useTitle'
import { useQuery } from '@tanstack/react-query'

export default function Account() {
  const { theme, toggleTheme } = useTheme()
  const { logout } = useAuth()
  const title = useTitle('Mein Account')

  return (
    <main>
      {title}
      <Grid mode="autofill">
        <Button onClick={logout}>
          <FontAwesomeIcon icon={faRightFromBracket} />
          <span> Abmelden</span>
        </Button>
      </Grid>
      <section>
        <h2>Einstellungen</h2>
        <Grid size="large">
          <Switch label="Dark Mode" defaultValue={theme === 'dark'} onChange={toggleTheme} />
        </Grid>
      </section>
      <UserInfos />
    </main>
  )
}

function UserInfos() {
  const { user } = useAuth()

  const {
    data: clans,
    error,
    isLoading,
  } = useQuery<Clan[], ApiError>({
    queryKey: [routes.clans.list],
  })

  function permissions() {
    const items: DataListItem[] = []
    if (!clans || !user) return items

    const joinClanNames = (tags: string[]) => {
      return tags.map((tag) => clans.find((c) => c.tag === tag)?.name).join(', ')
    }

    if (user.leaderOf?.length) {
      items.push({ label: 'Anführer', value: joinClanNames(user.leaderOf) })
    }
    if (user.coLeaderOf?.length) {
      items.push({ label: 'Vize-Anführer', value: joinClanNames(user.coLeaderOf) })
    }
    if (user.memberOf?.length) {
      items.push({ label: 'Mitglied', value: joinClanNames(user.memberOf) })
    }
    return items
  }

  return (
    <section>
      <h2>Benutzerinfos</h2>
      {clans ? (
        <Center>
          <DataList>
            {[
              { label: 'Benutzer', value: user?.name },
              { label: 'ID', value: user?.id },
              { label: 'Administrator', value: boolToAscii(user?.isAdmin) },
              ...permissions(),
            ]}
          </DataList>
        </Center>
      ) : (
        <QueryState loading={isLoading} error={error} />
      )}
    </section>
  )
}
