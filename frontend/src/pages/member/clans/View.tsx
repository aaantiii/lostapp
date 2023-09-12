import { useOutletContext } from 'react-router-dom'
import useDocumentTitle from '@hooks/useDocumentTitle'
import { CardList, Card } from '@components/Card'
import { MemberOutletContext } from '@context/types'
import DataSearch from '@components/DataSearch'
import { ClanMember, ClanMemberRoleTranslated } from '@api/types/clan'
import DataList from '@components/DataList'

export default function ViewClan() {
  const { clan, clanSettings } = useOutletContext<MemberOutletContext>()
  const heading = useDocumentTitle(clan?.name ?? 'Clan Details')

  if (!clan || !clanSettings) {
    return (
      <main>
        {heading}
        <p>Beim Laden des Clans ist ein Fehler aufgetreten.</p>
      </main>
    )
  }

  return (
    <main>
      {heading}
      <section>
        <h2>Kickpunkte</h2>
        <DataList
          data={[
            {
              title: 'Kickpunkte bis zum Kick',
              value: clanSettings.maxKickpoints,
            },
            {
              title: 'Minimale Season Wins',
              value: clanSettings.minSeasonWins,
            },
            {
              title: 'Tage bis zum Abbau',
              value: clanSettings.kickpointsExpireAfterDays,
            },
            {
              title: 'Kickpunkte für Season Wins',
              value: clanSettings.kickpointsSeasonWins,
            },
            {
              title: 'Kickpunkte für verpassten CK-Angriff',
              value: clanSettings.kickpointsCWMissed,
            },
            {
              title: 'Kickpunkte für CK-Fail',
              value: clanSettings.kickpointsCWFail,
            },
            {
              title: 'Kickpunkte für verpassten CKL-Angriff',
              value: clanSettings.kickpointsCWLMissed,
            },
            {
              title: 'Kickpunkte für CKL 0-Sterne Angriff',
              value: clanSettings.kickpointsCWLZeroStars,
            },
            {
              title: 'Kickpunkte für CKL 1-Stern Angriff',
              value: clanSettings.kickpointsCWLOneStar,
            },
            {
              title: 'Kickpunkte für verpassten Raid',
              value: clanSettings.kickpointsRaidMissed,
            },
            {
              title: 'Kickpunkte für Raid-Fail',
              value: clanSettings.kickpointsRaidFail,
            },
            {
              title: 'Kickpunkte für Clan Spiele',
              value: clanSettings.kickpointsClanGames,
            },
          ]}
        />
      </section>
      <section>
        <h2>Clan Member</h2>
        {clan.members > 0 ? (
          <DataSearch key="member" title="Mitglieder" searchKeys={['name']} data={clan.memberList}>
            {(results?: ClanMember[]) => {
              return results && results.length > 0 ? (
                <CardList>
                  {results.map((member) => (
                    <Card key={member.tag} title={member.name} description={ClanMemberRoleTranslated.get(member.role)} />
                  ))}
                </CardList>
              ) : (
                <p>Keine Mitglieder gefunden.</p>
              )
            }}
          </DataSearch>
        ) : (
          <p>{clan ? `${clan.name} hat keine Mitglieder.` : 'Beim Abrufen der Mitglieder ist ein Fehler aufgetreten.'}</p>
        )}
      </section>
    </main>
  )
}
