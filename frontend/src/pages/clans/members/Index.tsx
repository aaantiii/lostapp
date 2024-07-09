import { urlEncodeTag } from '@/utils/cocFormatter'
import client from '@api/client'
import routes, { replaceIds } from '@api/routes'
import { AuthRole } from '@api/types/auth'
import { ClanMemberRoleTranslated } from '@api/types/coc'
import { ApiError, Clan, Kickpoint } from '@api/types/models'
import AlertDialog from '@components/AlertDialog'
import Button from '@components/Button'
import Card from '@components/Card'
import Grid from '@components/Grid'
import GridSearch from '@components/GridSearch'
import NotAvailable from '@components/NotAvailable'
import QueryState from '@components/QueryState'
import RoleRender from '@components/RoleRender'
import { GridSkeleton } from '@components/Skeletons'
import Spacer from '@components/Spacer'
import { useMessages } from '@context/messagesContext'
import { faCheck, faTimes, faXmark } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import useTitle from '@hooks/useTitle'
import { useQuery, useQueryClient } from '@tanstack/react-query'
import { AxiosError } from 'axios'
import { useState } from 'react'
import { useParams } from 'react-router-dom'

export default function ClanMembers() {
  const { clanTag } = useParams()
  const {
    data: clan,
    isLoading,
    error,
  } = useQuery<Clan, ApiError>({
    queryKey: [routes.clans.byTag, { clanTag }],
  })
  const title = useTitle(clan?.name ? `${clan.name} Mitglieder` : 'Clan Übersicht')

  function sumKickpoints(kickpoints: Kickpoint[] | undefined) {
    return kickpoints?.reduce((prev, cur) => prev + cur.amount, 0) ?? 0
  }

  return (
    <main>
      {title}
      <RoleRender role={AuthRole.ClanCoLeader}>
        <Grid mode="autofill">
          <NotAvailable trigger={<Button>Mitglied hinzufügen</Button>} />
        </Grid>
        <Spacer size="small" />
      </RoleRender>
      {clan ? (
        <GridSearch>
          {(query) => {
            const matches = clan.clanMembers
              .filter((m) => m.player?.name.toLowerCase().includes(query.toLowerCase()))
              .toSorted((a) => (a.player!.name.startsWith(query) ? 1 : -1))

            return matches.map((m) => (
              <Card
                title={m.player!.name}
                key={m.playerTag}
                description={ClanMemberRoleTranslated.get(m.clanRole)}
                fields={[
                  {
                    label: 'Kickpunkte',
                    value: sumKickpoints(m.kickpoints).toString(),
                  },
                ]}
                buttons={[
                  <Button key="kickpoints" to={`/clans/${urlEncodeTag(clan.tag)}/members/${urlEncodeTag(m.playerTag)}/kickpoints`}>
                    Kickpunkte
                  </Button>,
                  <RoleRender role={AuthRole.ClanCoLeader} clanTag={clan.tag} key="delete">
                    <AlertDeleteMember clanTag={clan.tag} memberTag={m.playerTag} clanName={clan.name} memberName={m.player!.name} />
                  </RoleRender>,
                ]}
              />
            ))
          }}
        </GridSearch>
      ) : (
        <QueryState loading={isLoading} error={error} loader={<GridSkeleton />} />
      )}
    </main>
  )
}

type AlertDeleteMemberProps = {
  clanTag: string
  memberTag: string
  memberName: string
  clanName: string
}

function AlertDeleteMember({ clanTag, memberTag, memberName, clanName }: AlertDeleteMemberProps) {
  const { sendMessage } = useMessages()
  const [isDeleting, setIsDeleting] = useState(false)
  const queryClient = useQueryClient()

  async function handleDelete() {
    setIsDeleting(true)
    const url = replaceIds(routes.clans.members.byTag, { clanTag: urlEncodeTag(clanTag), memberTag: urlEncodeTag(memberTag) })
    try {
      await client.delete(url)
    } catch (err) {
      if (err instanceof AxiosError && err.response?.data) {
        const msg = err.response.data as ApiError
        sendMessage({ message: msg.message, type: 'error' })
      }
      return
    } finally {
      setIsDeleting(false)
    }

    await queryClient.invalidateQueries({ queryKey: [routes.clans.byTag, { clanTag: urlEncodeTag(clanTag) }] })
    sendMessage({ message: `${memberName} wurde aus ${clanName} entfernt.`, type: 'success' })
  }

  return (
    <AlertDialog
      title="Mitglied entfernen"
      description={`Möchtest du ${memberName} wirklich aus ${clanName} entfernen? Achtung: Diese Aktion löscht auch alle Kickpunkte und kann nicht rückgängig gemacht werden.`}
      trigger={<Button className="red">Mitglied entfernen</Button>}
      confirm={
        <Button className="green" loading={isDeleting} onClick={handleDelete}>
          <FontAwesomeIcon icon={faCheck} className="large" />
        </Button>
      }
      cancel={
        <Button className="red" disabled={isDeleting}>
          <FontAwesomeIcon icon={faXmark} className="large" />
        </Button>
      }
    />
  )
}
