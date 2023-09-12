import { useOutletContext } from 'react-router-dom'
import Button from '@components/Button'
import { CardList, Card } from '@components/Card'
import ExperienceLevel from '@components/ExperienceLevel'
import useDocumentTitle from '@hooks/useDocumentTitle'
import { formatPlayerClanRoles, urlEncodeTag } from '@fmt/cocFormatter'
import { MemberOutletContext } from '@context/types'
import { useMessage } from '@context/messageContext'
import { Player } from '@api/types/player'

export default function ViewMember() {
  const heading = useDocumentTitle('Member Details')
  const { sendMessage } = useMessage()
  const { player: member } = useOutletContext<MemberOutletContext>()

  function handleCopyTag(player: Player) {
    navigator.clipboard.writeText(player.tag)
    sendMessage({
      message: `Tag von ${player.name} kopiert!`,
      type: 'success',
    })
  }

  return (
    <main>
      {heading}
      <section>
        {member?.clans && (
          <CardList flexDirection="column">
            <Card
              title={member.name}
              description={formatPlayerClanRoles(member)}
              thumbnail={<ExperienceLevel level={member.expLevel} />}
              fields={[
                {
                  title: 'Season Wins',
                  value: member.attackWins,
                  style: { color: member.attackWins >= 80 ? 'green' : 'red' },
                  key: `wins${member.tag}`,
                },
              ]}
              buttons={[
                <Button key="copy-tag" onClick={() => handleCopyTag(member)}>
                  Tag kopieren
                </Button>,
                ...member.clans.map((clan) => (
                  <Button key={clan.tag} to={`/member/clans/${urlEncodeTag(clan.tag)}`}>
                    {clan.name}
                  </Button>
                )),
              ]}
              key={member.tag}
            />
          </CardList>
        )}
      </section>
    </main>
  )
}
