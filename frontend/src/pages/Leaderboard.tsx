import { numberFormatter } from '@/utils/intlFormatter'
import routes from '@api/routes'
import { PaginatedResponse } from '@api/types/base'
import { ComparableStatistic, PlayerStatistic } from '@api/types/dtos'
import { ApiError, Clan } from '@api/types/models'
import Grid from '@components/Grid'
import QueryState from '@components/QueryState'
import Select from '@components/Select'
import { TableSkeleton } from '@components/Skeletons'
import Spacer from '@components/Spacer'
import Table from '@components/Table'
import useTitle from '@hooks/useTitle'
import { useQuery } from '@tanstack/react-query'
import { useEffect } from 'react'
import { useSearchParams } from 'react-router-dom'

export default function Leaderboard() {
  const title = useTitle('Leaderboard')
  const [searchParams] = useSearchParams({
    page: '1',
    limit: '10',
  })
  const { data, isLoading, error } = useQuery<PaginatedResponse<PlayerStatistic>, ApiError>({
    enabled: Boolean(searchParams.get('s')),
    queryKey: [
      routes.players.live.leaderboard,
      null,
      {
        page: searchParams.get('page') ?? '1',
        limit: searchParams.get('limit')!,
        statName: searchParams.get('s')!,
        clanTag: searchParams.get('clan') ?? undefined,
      },
    ],
  })

  return (
    <main>
      {title}
      <p>Längere Ladezeiten sind der Clash of Clans API geschuldet.</p>
      <LeaderboardFilters />
      <Spacer size="small" />
      {data?.items?.length ? (
        <Table pagination={data?.pagination} header={['#', 'Name', 'Clan', 'Wert']} rawData={data.items}>
          {data?.items.map((player) => (
            <tr key={player.tag}>
              <td>{player.placement}</td>
              <td>{player.name}</td>
              <td>{player.clanName}</td>
              <td>{numberFormatter.format(player.value)}</td>
            </tr>
          ))}
        </Table>
      ) : (
        <QueryState loading={isLoading} error={error} loader={<TableSkeleton />} />
      )}
    </main>
  )
}

function LeaderboardFilters() {
  const { data: statList } = useQuery<ComparableStatistic[]>({
    queryKey: [routes.players.stats.list],
  })
  const { data: clans } = useQuery<Clan[]>({
    queryKey: [routes.clans.list],
  })
  const [searchParams, setSearchParams] = useSearchParams()

  useEffect(() => {
    if (!statList || statList.length === 0) return
    if (searchParams.get('s')) return

    setSearchParams(
      (prev) => {
        prev.set('s', statList[0].name)
        return prev
      },
      { replace: true }
    )
  }, [statList])

  return (
    <Grid>
      {statList && (
        <Select
          disableClear
          label="Statistik"
          defaultValue={searchParams.get('s') ?? statList[0].name}
          onChange={(value) => {
            setSearchParams(
              (prev) => {
                prev.set('s', value!)
                prev.set('page', '1')
                return prev
              },
              { replace: true }
            )
          }}
          options={statList.map((stat) => ({
            value: stat.name,
            label: stat.displayName,
          }))}
        />
      )}
      {clans && (
        <Select
          label="Clan"
          defaultValue={searchParams.get('clan') ?? ''}
          placeholder="Alle anzeigen"
          onChange={(value) => {
            setSearchParams(
              (prev) => {
                if (!value) {
                  prev.delete('clan')
                  return prev
                }
                prev.set('clan', value!)
                prev.set('page', '1')
                return prev
              },
              { replace: true }
            )
          }}
          options={clans.map((clan) => ({
            value: clan.tag,
            label: clan.name,
          }))}
        />
      )}
    </Grid>
  )
}
