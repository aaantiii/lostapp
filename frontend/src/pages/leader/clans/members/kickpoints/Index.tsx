import { useOutletContext, useParams } from 'react-router-dom'
import { useQuery } from '@tanstack/react-query'
import { LeaderOutletContext } from '@context/types'
import routes from '@api/routes'
import useDocumentTitle from '@hooks/useDocumentTitle'
import { Kickpoint } from '@api/types/kickpoint'
import { Card, CardList } from '@components/Card'
import { addDaysToDate, dateTimeFormatter, timeUntil } from '@fmt/intlFormatter'
import Button from '@components/Button'
import DialogNew from './DialogNew'
import DialogEdit from './DialogEdit'
import DialogDelete from './DialogDelete'
import { useState } from 'react'

export default function Kickpoints() {
  const heading = useDocumentTitle('Kickpunkte')
  const { clanTag, memberTag } = useParams()
  const { player, clanSettings } = useOutletContext<LeaderOutletContext>()

  const { data: kickpoints } = useQuery<Kickpoint[]>({
    queryKey: [routes.clans.members.kickpoints.byTag, { clanTag, memberTag }],
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
            title={player.name}
            key={player.tag}
            description={`Kickpunkte: ${kickpoints.reduce((prev, kickpoint) => prev + kickpoint.amount, 0)} von ${clanSettings.maxKickpoints}`}
            buttons={[<DialogNew />]}
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
                description={kickpoint.reason}
                fields={[
                  {
                    title: 'Anzahl Punkte',
                    value: kickpoint.amount,
                    key: 'amount',
                  },
                  {
                    title: 'Erhalten am',
                    value: dateTimeFormatter.format(new Date(kickpoint.date)),
                    key: 'date',
                  },
                  {
                    title: 'Läuft ab in',
                    value: timeUntil(addDaysToDate(new Date(kickpoint.date), clanSettings.kickpointsExpireAfterDays)),
                    key: 'expires',
                  },
                ]}
                buttons={[
                  <Button key="edit">Kickpunkt bearbeiten</Button>,
                  <Button className="red" key="delete">
                    Kickpunkt löschen
                  </Button>,
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
