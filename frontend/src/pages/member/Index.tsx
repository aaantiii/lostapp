import { useNavigate, useOutletContext } from 'react-router-dom'
import Button from '../../components/Button'
import { CardList, Card } from '../../components/Card'
import ExperienceLevel from '../../components/ExperienceLevel'
import { useAuth } from '../../context/authContext'
import useDocumentTitle from '../../hooks/useDocumentTitle'
import { formatPlayerClanRoles, urlEncodeTag } from '../../fmt/cocFormatter'
import { MemberOutletContext } from '../../types/context'
import { useMessage } from '../../context/messageContext'
import { Player } from '../../api/types/player'

export default function Index() {
  const heading = useDocumentTitle('Member Übersicht')

  const navigate = useNavigate()
  const { sendMessage } = useMessage()
  const { discordUser } = useAuth()
  const { userPlayers: cocAccounts } = useOutletContext<MemberOutletContext>()

  function handleCopyTag(player: Player) {
    navigator.clipboard.writeText(player.tag)
    sendMessage({
      message: `Tag von ${player.name} kopiert!`,
      type: 'success',
    })
  }

  return (
    <main>
      <hgroup>
        {heading}
        <h2>Willkommen {discordUser?.username} 👋</h2>
        <h4>Deine Accounts</h4>
      </hgroup>

      {cocAccounts ? (
        <CardList>
          {cocAccounts.map((account) => (
            <Card
              title={account.name}
              description={formatPlayerClanRoles(account)}
              thumbnail={<ExperienceLevel level={account.expLevel} />}
              key={`card-${account.tag}`}
              fields={[
                {
                  title: 'Season wins',
                  value: account.attackWins,
                  style: { color: account.attackWins >= 80 ? 'green' : 'red' },
                  key: `season-wins`,
                },
              ]}
              buttons={[
                <Button key="show-player-details" onClick={() => navigate(`/member/${urlEncodeTag(account.tag)}`)}>
                  Details
                </Button>,
                <Button key="copy-tag" onClick={() => handleCopyTag(account)}>
                  Tag kopieren
                </Button>,
              ]}
            />
          ))}
        </CardList>
      ) : (
        <p>Du hast noch keine Clash of Clans Accounts verknüpft.</p>
      )}
    </main>
  )
}
