import { useQuery } from '@tanstack/react-query'
import useDocumentTitle from '../../hooks/useDocumentTitle'
import { Card, CardList } from '../../components/Card'
import Button from '../../components/Button'
import { useOutletContext } from 'react-router-dom'
import { LeaderOutletContext } from '../../types/context'
import { useAuth } from '../../context/authContext'
import { urlEncodeTag } from '../../fmt/cocFormatter'

export default function LeaderIndex() {
  const heading = useDocumentTitle('Anführer Übersicht')
  const { discordUser } = useAuth()
  const { leadingClans } = useOutletContext<LeaderOutletContext>()

  return (
    <main>
      <hgroup>
        {heading}
        <h2>Willkommen {discordUser?.username} 👋</h2>
        <h4>Deine Clans</h4>
      </hgroup>

      {leadingClans && leadingClans.length > 0 ? (
        <CardList>
          {leadingClans.map((clan) => (
            <Card
              key={clan.tag}
              title={clan.name}
              description={`Mitglieder: ${clan.members} / 50`}
              thumbnail={<img src={clan.badgeUrl} alt="Clan Badge" />}
              buttons={[
                <Button key="members" to={`/leader/clans/${urlEncodeTag(clan.tag)}/members`}>
                  Mitglieder verwalten
                </Button>,
                <Button key="settings" to={`/leader/clans/${urlEncodeTag(clan.tag)}/settings`}>
                  Clan-Einstellungen
                </Button>,
              ]}
            ></Card>
          ))}
        </CardList>
      ) : (
        <div className="center">Es ist ein Fehler aufgetreten.</div>
      )}
    </main>
  )
}
