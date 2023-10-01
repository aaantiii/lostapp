import { useOutletContext, useParams } from 'react-router-dom'
import Button from '@components/Button'
import { CardList, Card } from '@components/Card'
import ExperienceLevel from '@components/ExperienceLevel'
import useDocumentTitle from '@hooks/useDocumentTitle'
import { formatPlayerClanRole, formatPlayerClanRoles, urlEncodeTag } from '@fmt/cocFormatter'
import { MemberOutletContext } from '@context/types'
import { useMessage } from '@context/messageContext'
import { Player } from '@api/types/player'
import routes from '@api/routes'
import { useQuery } from '@tanstack/react-query'
import { Kickpoint } from '@api/types/kickpoint'
import { addDaysToDate, dateFormatter, timeUntil } from '@fmt/intlFormatter'

export default function ViewMember() {
  const heading = useDocumentTitle('Member Details')
  const { clanTag, memberTag } = useParams()
  const { sendMessage } = useMessage()
  const { player: member, clanSettings, clan } = useOutletContext<MemberOutletContext>()

  const { data: kickpoints } = useQuery<Kickpoint[]>({
    queryKey: [routes.clans.members.kickpoints.byClanMember, { memberTag, clanTag }],
    enabled: member !== undefined,
  })

  function handleCopyTag(player: Player) {
    navigator.clipboard.writeText(player.tag)
    sendMessage({
      message: `Tag von ${player.name} kopiert!`,
      type: 'success',
    })
  }

  if (!member || !clan)
    return (
      <main>
        {heading}
        <p>Das angeforderte Mitglied konnte nicht gefunden werden.</p>
      </main>
    )

  return (
    <main>
      {heading}
      <section>
        {member && clan && (
          <CardList flexDirection="column">
            <Card
              title={member.name}
              description={formatPlayerClanRole(member.clans.find((c) => c.tag === clan.tag))}
              thumbnail={<ExperienceLevel level={member.expLevel} />}
              fields={[
                {
                  title: 'Season Wins',
                  value: member.attackWins,
                  style: { color: member.attackWins >= 80 ? 'green' : 'red' },
                  key: `wins${member.tag}`,
                },
              ]}
              buttons={[
                <Button key="copy-tag" onClick={() => handleCopyTag(member)}>
                  Tag kopieren
                </Button>,
                <Button key={clan.tag} to={`/member/clans/${urlEncodeTag(clan.tag)}`}>
                  {clan.name}
                </Button>,
              ]}
              key={member.tag}
            />
          </CardList>
        )}
      </section>
      <section>
        <h2>Aktive Kickpunkte</h2>
        {kickpoints && kickpoints.length > 0 && clanSettings ? (
          <CardList>
            {kickpoints.map((kickpoint) => (
              <Card
                key={kickpoint.id}
                title={`Kickpunkt #${kickpoint.id}`}
                description={kickpoint.description}
                fields={[
                  {
                    title: 'Anzahl Punkte',
                    value: kickpoint.amount,
                    key: 'amount',
                  },
                  {
                    title: 'Erhalten am',
                    value: dateFormatter.format(new Date(kickpoint.date)),
                    key: 'date',
                  },
                  {
                    title: 'Läuft ab in',
                    value: timeUntil(addDaysToDate(new Date(kickpoint.date), clanSettings.kickpointsExpireAfterDays)),
                    key: 'expires',
                  },
                  {
                    title: 'Erstellt von',
                    value: kickpoint.createdByUser.name,
                    key: 'createdBy',
                  },
                ]}
              />
            ))}
          </CardList>
        ) : (
          <p>Es wurden keine aktiven Kickpunkte gefunden.</p>
        )}
      </section>
    </main>
  )
}
