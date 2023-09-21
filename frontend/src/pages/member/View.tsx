import { useNavigate, useOutletContext } from 'react-router-dom'
import Button from '@components/Button'
import { CardList, Card } from '@components/Card'
import ExperienceLevel from '@components/ExperienceLevel'
import { useAuth } from '@context/authContext'
import useDocumentTitle from '@hooks/useDocumentTitle'
import { urlEncodeTag } from '@fmt/cocFormatter'
import { MemberOutletContext } from '@context/types'
import { ClanMemberRoleTranslated } from '@api/types/clan'

export default function Index() {
  const heading = useDocumentTitle('Member Übersicht')
  const navigate = useNavigate()
  const { discordUser } = useAuth()
  const { userPlayers } = useOutletContext<MemberOutletContext>()

  return (
    <main>
      <hgroup>
        {heading}
        <h2>Willkommen {discordUser?.name} 👋</h2>
        <h4>Deine Member Accounts</h4>
      </hgroup>
      <section>
        {userPlayers ? (
          <CardList>
            {userPlayers.map((account) => [
              ...account.clans.map((clan) => (
                <Card
                  title={account.name}
                  description={`${ClanMemberRoleTranslated.get(clan.role)} in ${clan.name}`}
                  thumbnail={<ExperienceLevel level={account.expLevel} />}
                  key={`${account.tag}${clan.tag}`}
                  fields={[
                    {
                      title: 'Season wins',
                      value: account.attackWins,
                      style: { color: account.attackWins >= 80 ? 'green' : 'red' },
                      key: `season-wins`,
                    },
                  ]}
                  buttons={[
                    <Button
                      key="member-details"
                      onClick={() => navigate(`/member/clans/${urlEncodeTag(clan.tag)}/members/${urlEncodeTag(account.tag)}`)}
                    >
                      Details
                    </Button>,
                    <Button key="clan-details" onClick={() => navigate(`/member/clans/${urlEncodeTag(clan.tag)}`)}>
                      Clan Details
                    </Button>,
                  ]}
                />
              )),
            ])}
          </CardList>
        ) : (
          <p>Du hast noch keine Clash of Clans Accounts verknüpft.</p>
        )}
      </section>
    </main>
  )
}
