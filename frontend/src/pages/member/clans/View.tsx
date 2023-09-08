import { useOutletContext } from 'react-router-dom'
import useDocumentTitle from '../../../hooks/useDocumentTitle'
import { CardList, Card } from '../../../components/Card'
import { MemberOutletContext } from '../../../types/context'
import DataSearch from '../../../components/DataSearch'
import { ClanMember, ClanMemberRoleTranslated } from '../../../api/types/clan'

export default function ViewClan() {
  const { clan } = useOutletContext<MemberOutletContext>()
  const heading = useDocumentTitle(`${clan ? clan.name : 'Clan'} Mitglieder`)

  return (
    <main>
      {heading}
      <section>
        <h2>Clan Member</h2>
        {clan && clan.members > 0 ? (
          <DataSearch key="member" title="Mitglieder" searchKeys={['name']} data={clan.memberList}>
            {(members?: ClanMember[]) => {
              return members && members.length > 0 ? (
                <CardList>
                  {members.map((member) => (
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
      <section>
        <h2>Kickpunkte</h2>
      </section>
    </main>
  )
}
