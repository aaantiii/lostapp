import { useNavigate, useOutletContext } from 'react-router-dom'
import Button from '@components/Button'
import { CardList, Card } from '@components/Card'
import ExperienceLevel from '@components/ExperienceLevel'
import useDocumentTitle from '@hooks/useDocumentTitle'
import { urlEncodeTag } from '@fmt/cocFormatter'
import { MemberOutletContext } from '@context/types'
import { ClanMemberRoleTranslated } from '@api/types/clan'

export default function Index() {
  const navigate = useNavigate()
  const { player } = useOutletContext<MemberOutletContext>()
  const heading = useDocumentTitle(player?.name ?? 'Mitglied ansehen')

  return (
    <main>
      <hgroup>
        {heading}
        <h4>{player ? `${player.name}'s Clans` : 'Member Accounts'}</h4>
      </hgroup>
      <section>
        {player ? (
          <CardList>
            {player.clans.map((clan) => (
              <Card
                title={player.name}
                description={`${ClanMemberRoleTranslated.get(clan.role)} in ${clan.name}`}
                thumbnail={<ExperienceLevel level={player.expLevel} />}
                key={clan.tag}
                buttons={[
                  <Button
                    key="member-details"
                    onClick={() => navigate(`/member/clans/${urlEncodeTag(clan.tag)}/members/${urlEncodeTag(player.tag)}`)}
                  >
                    Mitglied ansehen
                  </Button>,
                  <Button key="clan-details" onClick={() => navigate(`/member/clans/${urlEncodeTag(clan.tag)}`)}>
                    Clan Details
                  </Button>,
                ]}
              />
            ))}
          </CardList>
        ) : (
          <p>Dieser Spieler konnte nicht gefunden werden.</p>
        )}
      </section>
    </main>
  )
}
