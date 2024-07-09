import { formatMemberClan, urlEncodeTag } from '@/utils/cocFormatter'
import routes from '@api/routes'
import { PaginatedResponse } from '@api/types/base'
import { LivePlayer } from '@api/types/coc'
import { ApiError } from '@api/types/models'
import { PlayersParams } from '@api/types/params'
import Button from '@components/Button'
import Card from '@components/Card'
import ExperienceLevel from '@components/ExperienceLevel'
import Grid from '@components/Grid'
import Paginator from '@components/Paginator'
import QueryState from '@components/QueryState'
import { GridSkeleton } from '@components/Skeletons'
import { useAuth } from '@context/authContext'
import useTitle from '@hooks/useTitle'
import { useQuery } from '@tanstack/react-query'
import { useSearchParams } from 'react-router-dom'

export default function Dashboard() {
  const title = useTitle('Dashboard')

  return (
    <main>
      {title}
      <Warnings />
      <MyAccounts />
    </main>
  )
}

function Warnings() {
  return <section></section>
}

function CoLeaderHeader() {
  return <section></section>
}

function MyAccounts() {
  const { user } = useAuth()
  const [searchParams] = useSearchParams({
    page: '1',
    limit: '12',
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
        discordId: user!.id,
        isMember: 'true',
        page: searchParams.get('page')!,
        limit: searchParams.get('limit')!,
      } satisfies PlayersParams,
    ],
  })

  return (
    <section>
      <h2>Deine Accounts</h2>
      {(accounts?.items?.length ?? 0) > 0 ? (
        <Paginator pagination={accounts?.pagination} limits={[12, 24]}>
          <Grid>
            {accounts?.items?.map((account) =>
              account.clanMembers?.map((member) => (
                <Card
                  key={account.tag}
                  title={account.name}
                  description={formatMemberClan(member)}
                  thumbnail={<ExperienceLevel level={account.expLevel} />}
                  buttons={[
                    <Button key="clan" to={`/clans/${urlEncodeTag(member.clanTag)}`}>
                      Clan
                    </Button>,
                  ]}
                />
              ))
            )}
          </Grid>
        </Paginator>
      ) : (
        <QueryState loading={accountsLoading} error={accountsError} loader={<GridSkeleton />} />
      )}
    </section>
  )
}
