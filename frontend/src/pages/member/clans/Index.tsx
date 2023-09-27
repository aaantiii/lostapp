import { useEffect, useState } from 'react'
import { CardList, Card } from '@components/Card'
import Button from '@components/Button'
import useDocumentTitle from '@hooks/useDocumentTitle'
import { urlEncodeTag } from '@fmt/cocFormatter'
import { useQuery } from '@tanstack/react-query'
import { useAuth } from '@context/authContext'
import routes from '@api/routes'
import LoadingScreen from '@components/LoadingScreen'
import { Clan } from '@api/types/clan'
import { useMessage } from '@context/messageContext'

export default function Index() {
  const heading = useDocumentTitle('Lost Family')

  const { sendMessage } = useMessage()
  const { discordUser } = useAuth()
  const [totalClanMembers, setTotalClanMembers] = useState(0)
  const {
    data: lostClans,
    isLoading,
    isError,
  } = useQuery<Clan[]>({
    queryKey: [routes.clans.all],
    enabled: discordUser !== undefined,
  })

  useEffect(() => {
    if (!lostClans) return
    setTotalClanMembers(lostClans.reduce((sum, clan) => sum + clan.members, 0))
  }, [lostClans])

  function handleCopyTag(clan: Clan) {
    navigator.clipboard.writeText(clan.tag)
    sendMessage({
      message: `Tag von ${clan.name} kopiert!`,
      type: 'success',
    })
  }

  if (isLoading) return <LoadingScreen />

  return (
    <main>
      <hgroup>
        {heading}
        <h4>Derzeit hat die Lost Family {totalClanMembers} Mitglieder 💙</h4>
      </hgroup>
      {isError ? (
        <p>Beim Abrufen der Clans ist ein Fehler aufgetreten.</p>
      ) : (
        lostClans && (
          <CardList>
            {lostClans.map((clan) => (
              <Card
                title={clan.name}
                description={`Member: ${clan.members} / 50`}
                thumbnail={<img src={clan.badgeUrl} alt={`${clan.name} Clan Badge`} />}
                buttons={[
                  <Button to={`/member/clans/${urlEncodeTag(clan.tag)}`} key="link">
                    Clan ansehen
                  </Button>,
                  <Button onClick={() => handleCopyTag(clan)} key="copy-tag">
                    Tag kopieren
                  </Button>,
                ]}
                key={clan.tag}
              />
            ))}
          </CardList>
        )
      )}
    </main>
  )
}
