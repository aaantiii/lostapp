import { useOutletContext } from 'react-router-dom'
import useDocumentTitle from '@hooks/useDocumentTitle'
import { LeaderOutletContext } from '@context/types'
import { Card, CardList } from '@components/Card'
import { urlEncodeTag } from '@fmt/cocFormatter'
import Button from '@components/Button'
import DataSearch from '@components/DataSearch'
import { ClanMemberRoleTranslated } from '@api/types/clan'
import { useQuery } from '@tanstack/react-query'
import routes from '@api/routes'
import { ClanMemberKickpoints } from '@api/types/kickpoint'

export default function ClanMembersIndex() {
  const { clan, clanSettings } = useOutletContext<LeaderOutletContext>()
  const heading = useDocumentTitle(`${clan ? clan.name : 'Clan'} Mitglieder`)

  const { data: clanMemberKickpoints } = useQuery<ClanMemberKickpoints[]>({
    queryKey: [routes.clans.members.kickpoints.all, { clanTag: urlEncodeTag(clan?.tag) }],
    enabled: clan !== undefined,
    refetchOnMount: 'always',
  })

  const membersWithMaxKickpoints = clanMemberKickpoints?.filter((member) => member.amount >= (clanSettings?.maxKickpoints ?? Infinity))

  return (
    <main>
      {heading}
      {membersWithMaxKickpoints && membersWithMaxKickpoints.length > 0 && (
        <section>
          <h2>&#9888; Kickpunkte-Limit überschritten &#9888;</h2>
          <p>Die maximale Anzahl von {clanSettings?.maxKickpoints ?? 0} Kickpunkten wurde von folgenden Mitgliedern erreicht:</p>
          <CardList>
            {membersWithMaxKickpoints.map((member) => (
              <Card
                key={member.tag}
                title={member.name}
                description={ClanMemberRoleTranslated.get(member.role)}
                fields={[
                  {
                    title: 'Kickpunkte',
                    value: member.amount,
                    key: member.tag,
                    style: (clanSettings?.maxKickpoints ?? 0) / 2 <= member.amount ? { color: 'red' } : {},
                  },
                ]}
                buttons={[
                  <Button key="kickpoints" to={`/leader/clans/${urlEncodeTag(clan?.tag)}/members/${urlEncodeTag(member.tag)}/kickpoints`}>
                    Kickpunkte ansehen
                  </Button>,
                ]}
              />
            ))}
          </CardList>
        </section>
      )}
      <section>
        <h2>Clan Member</h2>
        {clan && clanMemberKickpoints ? (
          <DataSearch key="member" title="Mitglieder" searchKeys={['name']} data={clanMemberKickpoints}>
            {(members?: ClanMemberKickpoints[]) => {
              return members ? (
                <CardList>
                  {members.map((member) => (
                    <Card
                      key={member.tag}
                      title={member.name}
                      description={ClanMemberRoleTranslated.get(member.role)}
                      fields={[
                        {
                          title: 'Kickpunkte',
                          value: member.amount,
                          key: member.tag,
                          style: (clanSettings?.maxKickpoints ?? Infinity) / 2 <= member.amount ? { color: 'red' } : {},
                        },
                      ]}
                      buttons={[
                        <Button key="kickpoints" to={`/leader/clans/${urlEncodeTag(clan.tag)}/members/${urlEncodeTag(member.tag)}/kickpoints`}>
                          Kickpunkte
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
          <p>Es wurden keine Mitglieder für {clan!.name} gefunden.</p>
        )}
      </section>
    </main>
  )
}
