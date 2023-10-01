import { useOutletContext, useParams } from 'react-router-dom'
import { useQuery } from '@tanstack/react-query'
import { LeaderOutletContext } from '@context/types'
import routes from '@api/routes'
import useDocumentTitle from '@hooks/useDocumentTitle'
import { Kickpoint } from '@api/types/kickpoint'
import { Card, CardList } from '@components/Card'
import { addDaysToDate, dateFormatter, timeUntil } from '@fmt/intlFormatter'
import DialogNew from './DialogNew'
import { useMemo } from 'react'
import { SelectOptionGroup, selectOptionOther } from '@components/Select'
import DialogDelete from './DialogDelete'
import DialogEdit from './DialogEdit'

export default function Kickpoints() {
  const { clanTag, memberTag } = useParams()
  const { player, clanSettings } = useOutletContext<LeaderOutletContext>()
  const heading = useDocumentTitle(player ? `${player.name}'s Kickpunkte` : 'Kickpunkte')

  const { data: kickpoints, refetch: refreshKickpoints } = useQuery<Kickpoint[]>({
    queryKey: [routes.clans.members.kickpoints.byClanMember, { clanTag, memberTag }],
    enabled: clanTag !== undefined && memberTag !== undefined,
  })

  if (!clanSettings || !player || !kickpoints) {
    return (
      <main>
        {heading}
        <p>Beim Laden des Mitglieds ist ein Fehler aufgetreten.</p>
      </main>
    )
  }

  return (
    <main>
      {heading}
      <section>
        <h2>Mitglied</h2>
        <CardList>
          <Card
            key={player.tag}
            title={player.name}
            description={`Kickpunkte: ${kickpoints.reduce((prev, kickpoint) => prev + kickpoint.amount, 0)} von ${clanSettings.maxKickpoints}`}
            buttons={[<DialogNew key="new" onSuccess={refreshKickpoints} />]}
          />
        </CardList>
      </section>
      <section>
        <h2>Aktive Kickpunkte</h2>
        {kickpoints && kickpoints.length > 0 ? (
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
                buttons={[
                  <DialogEdit key="edit" kickpoint={kickpoint} onSuccess={refreshKickpoints} />,
                  <DialogDelete key="delete" kickpointId={kickpoint.id} onSuccess={refreshKickpoints} />,
                ]}
              />
            ))}
          </CardList>
        ) : (
          <p>{player.name} hat keine aktiven Kickpunkte.</p>
        )}
      </section>
    </main>
  )
}
