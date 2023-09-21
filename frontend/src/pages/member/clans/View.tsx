import { useOutletContext, useParams } from 'react-router-dom'
import useDocumentTitle from '@hooks/useDocumentTitle'
import { CardList, Card } from '@components/Card'
import { MemberOutletContext } from '@context/types'
import DataSearch from '@components/DataSearch'
import { ClanMember, ClanMemberRoleTranslated } from '@api/types/clan'
import DataList from '@components/DataList'
import DataChangelog from '@components/DataChangelog'
import Button from '@components/Button'
import { urlEncodeTag } from '@fmt/cocFormatter'

export default function ViewClan() {
  const { clanTag, memberTag } = useParams()
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
        <h2>Kickpunkte-Einstellungen</h2>
        <DataChangelog data={clanSettings} type="updated" />
        <DataList
          data={[
            {
              title: 'Kickpunkte bis zum Kick',
              value: `${clanSettings.maxKickpoints} Kickpunkte`,
            },
            {
              title: 'Minimale Season Wins',
              value: `${clanSettings.minSeasonWins} Siege`,
            },
            {
              title: 'Dauer bis zum Abbau',
              value: `${clanSettings.kickpointsExpireAfterDays} Tage`,
            },
            {
              title: 'Bestrafung für Season Wins',
              value: `${clanSettings.kickpointsSeasonWins} Kickpunkte`,
            },
            {
              title: 'Bestrafung für verpassten CK-Angriff',
              value: `${clanSettings.kickpointsCWMissed} Kickpunkte`,
            },
            {
              title: 'Bestrafung für CK-Fail',
              value: `${clanSettings.kickpointsCWFail} Kickpunkte`,
            },
            {
              title: 'Bestrafung für verpassten CKL-Angriff',
              value: `${clanSettings.kickpointsCWLMissed} Kickpunkte`,
            },
            {
              title: 'Bestrafung für CKL 0-Sterne Angriff',
              value: `${clanSettings.kickpointsCWLZeroStars} Kickpunkte`,
            },
            {
              title: 'Bestrafung für CKL 1-Stern Angriff',
              value: `${clanSettings.kickpointsCWLOneStar} Kickpunkte`,
            },
            {
              title: 'Bestrafung für verpassten Raid',
              value: `${clanSettings.kickpointsRaidMissed} Kickpunkte`,
            },
            {
              title: 'Bestrafung für Raid-Fail',
              value: `${clanSettings.kickpointsRaidFail} Kickpunkte`,
            },
            {
              title: 'Bestrafung für Clan Spiele',
              value: `${clanSettings.kickpointsClanGames} Kickpunkte`,
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
                    <Card
                      key={member.tag}
                      title={member.name}
                      description={ClanMemberRoleTranslated.get(member.role)}
                      buttons={[
                        <Button to={`/member/clans/${clanTag}/members/${urlEncodeTag(member.tag)}`} key="view">
                          Mitglied ansehen
                        </Button>,
                      ]}
                    />
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
