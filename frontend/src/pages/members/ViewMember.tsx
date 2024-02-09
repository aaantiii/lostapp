import { formatMemberClan } from '@/utils/cocFormatter'
import routes from '@api/routes'
import { User } from '@api/types/auth'
import { PaginatedResponse } from '@api/types/base'
import { LivePlayer } from '@api/types/coc'
import { ApiError } from '@api/types/models'
import { PlayersParams } from '@api/types/params'
import Card from '@components/Card'
import ExperienceLevel from '@components/ExperienceLevel'
import Grid from '@components/Grid'
import Paginator from '@components/Paginator'
import QueryState from '@components/QueryState'
import { useAuth } from '@context/authContext'
import useTitle from '@hooks/useTitle'
import { useQuery } from '@tanstack/react-query'
import { useParams, useSearchParams } from 'react-router-dom'

export default function ViewMember() {
  const { discordId } = useParams()
  const { user } = useAuth()
  const [searchParams] = useSearchParams({
    page: '1',
    pageSize: '12',
  })

  const {
    data: accounts,
    isLoading: accountsLoading,
    error: accountsError,
  } = useQuery<PaginatedResponse<LivePlayer>, ApiError>({
    queryKey: [
      routes.players.live.index,
      null,
      {
        discordId: discordId === '@me' ? user?.id : discordId,
        page: searchParams.get('page')!,
        pageSize: searchParams.get('pageSize')!,
      } satisfies PlayersParams,
    ],
  })

  const {
    data: fetchedUser,
    isLoading: userLoading,
    error: userError,
  } = useQuery<User, ApiError>({
    queryKey: [routes.users.byId, { id: discordId === '@me' ? user?.id : discordId }],
  })

  const title = useTitle(fetchedUser?.name ?? 'Mitglied Übersicht')

  return (
    <main>
      {title}
      {(accounts?.items?.length ?? 0) > 0 ? (
        <Paginator pagination={accounts?.pagination}>
          <Grid>
            {accounts?.items?.map((account) => (
              <Card
                key={account.tag}
                title={account.name}
                description={account.clanMembers?.map(formatMemberClan).join(', ') ?? 'Kein LOST-Mitglied'}
                thumbnail={<ExperienceLevel level={account.expLevel} />}
              />
            ))}
          </Grid>
        </Paginator>
      ) : (
        <QueryState loading={accountsLoading || userLoading} error={accountsError ?? userError} />
      )}
    </main>
  )
}
