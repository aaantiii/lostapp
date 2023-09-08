import { useOutletContext, useParams } from 'react-router-dom'
import { useQuery } from '@tanstack/react-query'
import { LeaderOutletContext } from '../../../../types/context'
import routes from '../../../../api/routes'
import useDocumentTitle from '../../../../hooks/useDocumentTitle'
import { Kickpoint } from '../../../../api/types/kickpoint'
import { Card, CardList } from '../../../../components/Card'
import { addDaysToDate, dateTimeFormatter } from '../../../../fmt/formatters'

export default function Kickpoints() {
  const heading = useDocumentTitle('Kickpunkte')
  const { clanTag, memberTag } = useParams()
  const { player, clanSettings } = useOutletContext<LeaderOutletContext>()

  const { data: kickpoints } = useQuery<Kickpoint[]>({
    queryKey: [routes.clans.members.kickpoints.byTag, { clanTag, memberTag }],
    enabled: clanTag !== undefined && memberTag !== undefined,
  })

  return (
    <main>
      {heading}
      {player && <p className="bold">Auf dieser Seite kannst du die Kickpunkte von {player.name} verwalten.</p>}
      <section>
        <h2>Kickpunkte verwalten</h2>
      </section>
      {clanSettings && (
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
                      title: 'Datum',
                      value: dateTimeFormatter.format(new Date(kickpoint.date)),
                      key: 'date',
                    },
                    {
                      title: 'Läuft ab',
                      value: dateTimeFormatter.format(addDaysToDate(new Date(kickpoint.date), clanSettings.kickpointsExpireAfterDays)),
                      key: 'expires',
                    },
                  ]}
                  buttons={[]}
                />
              ))}
            </CardList>
          ) : (
            player && <p>{player.name} hat keine aktiven Kickpunkte.</p>
          )}
        </section>
      )}
    </main>
  )
}
