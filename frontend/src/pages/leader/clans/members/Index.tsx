import { useOutletContext } from 'react-router-dom'
import useDocumentTitle from '../../../../hooks/useDocumentTitle'
import { LeaderOutletContext } from '../../../../types/context'
import { Card, CardList } from '../../../../components/Card'
import { urlEncodeTag } from '../../../../fmt/cocFormatter'
import Button from '../../../../components/Button'
import DataSearch from '../../../../components/DataSearch'
import { ClanMemberRoleTranslated } from '../../../../api/types/clan'
import { useQuery } from '@tanstack/react-query'
import routes from '../../../../api/routes'
import { ClanMemberKickpoints } from '../../../../api/types/kickpoint'

export default function ClanMembersIndex() {
  const { clan, clanSettings } = useOutletContext<LeaderOutletContext>()
  const heading = useDocumentTitle(`${clan ? clan.name : 'Clan'} Mitglieder`)

  const { data: clanMemberKickpoints } = useQuery<ClanMemberKickpoints[]>({
    queryKey: [routes.clans.members.kickpoints.all, { clanTag: urlEncodeTag(clan?.tag) }],
    enabled: clan !== undefined,
  })

  return (
    <main>
      {heading}
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
                        <Button key="manage" to={`/leader/clans/${urlEncodeTag(clan.tag)}/members/${urlEncodeTag(member.tag)}/manage`}>
                          Mitglied verwalten
                        </Button>,
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
