import { useEffect, useState } from 'react'
import Select, { SelectOption, SelectOptionGroup } from '@components/Select'
import { Table } from '@components/Table'
import useDocumentTitle from '@hooks/useDocumentTitle'
import { useQuery } from '@tanstack/react-query'
import { PlayersParams } from '@api/types/params'
import LoadingScreen from '@components/LoadingScreen'
import { urlDecodeTag, urlEncodeTag } from '@fmt/cocFormatter'
import routes from '@api/routes'
import { PaginatedResponse } from '@api/types/pagination'
import UnsubmittableForm from '@components/UnsubmittableForm'
import usePageSize from '@hooks/usePageSize'
import { ComparableStatistic, PlayerStatistic } from '@api/types/playerStats'
import { Clan } from '@api/types/clan'
import { useSearchParams } from 'react-router-dom'

const clanTagFilterAll: SelectOption = { value: '#all', displayText: 'Alle Lost Clans' }

export default function Leaderboard() {
  const heading = useDocumentTitle('Leaderboard')

  const pageSize = usePageSize(20, 30)
  const [searchParams, setSearchParams] = useSearchParams({ clan: clanTagFilterAll.value })
  const [clansSelectOptions, setClanOptionGroup] = useState({ options: [] } as SelectOptionGroup)
  const [statisticsSelectOptions, setStatisticsOptionGroup] = useState({ options: [] } as SelectOptionGroup)

  const { data: clans, isLoading: clansLoading } = useQuery<Clan[]>({
    queryKey: [routes.clans.all],
    staleTime: Infinity,
  })

  const { data: comparableStatistics, isLoading: achievementsLoading } = useQuery<ComparableStatistic[]>({
    queryKey: [routes.players.comparableStats],
    staleTime: Infinity,
  })

  const {
    data: paginatedPlayers,
    refetch: fetchLeaderboardPlayers,
    isError,
  } = useQuery<PaginatedResponse<PlayerStatistic>>({
    queryKey: [
      routes.players.leaderboard,
      { statsId: comparableStatistics?.find((s) => s.name === searchParams.get('statistic'))?.id },
      {
        page: Number(searchParams.get('page') ?? '1'),
        pageSize: Number(searchParams.get('pageSize') ?? '30'),
        clanTag: searchParams.get('clan') === clanTagFilterAll.value ? '' : searchParams.get('clan') ?? '',
      } satisfies PlayersParams,
    ],
    enabled: false,
    retry: false,
  })

  useEffect(() => {
    if (!clans) return

    const clanOptions: SelectOption[] = clans.map((clan) => ({ value: clan.tag, displayText: clan.name }))
    clanOptions.unshift(clanTagFilterAll)
    setClanOptionGroup({ title: 'Clan wählen', options: clanOptions })
  }, [clans])

  useEffect(() => {
    if (!comparableStatistics) return

    setStatisticsOptionGroup({
      title: 'Achievement wählen',
      options: comparableStatistics.map((achievement) => ({
        value: achievement.name,
        displayText: achievement.displayName,
      })),
    })
  }, [comparableStatistics])

  useEffect(() => {
    if (!comparableStatistics || !searchParams.get('statistic') || !searchParams.get('clan')) return

    fetchLeaderboardPlayers()
  }, [searchParams, comparableStatistics])

  if (clansLoading || achievementsLoading) return <LoadingScreen />

  return (
    <main>
      {heading}
      <section>
        <h2>Filter</h2>
        <UnsubmittableForm>
          <Select
            defaultValue={urlDecodeTag(searchParams.get('clan') ?? clanTagFilterAll.value)}
            optionGroups={[clansSelectOptions]}
            onChange={(tag) =>
              setSearchParams(
                (prev) => {
                  prev.set('clan', urlEncodeTag(tag))
                  return prev
                },
                { replace: true }
              )
            }
            placeholder="Clan auswählen"
          />
          <Select
            defaultValue={searchParams.get('statistic') ?? undefined}
            optionGroups={[statisticsSelectOptions]}
            onChange={(statistic) =>
              setSearchParams(
                (prev) => {
                  prev.set('statistic', statistic)
                  return prev
                },
                { replace: true }
              )
            }
            placeholder="Statistik auswählen"
          />
        </UnsubmittableForm>
      </section>
      <section>
        <h2>Statistiken</h2>
        {paginatedPlayers?.items ? (
          <Table
            data={paginatedPlayers.items}
            rowCountColumn
            columns={[
              { prop: 'playerName', heading: 'Spieler', link: '/member', linkIdProp: 'playerTag' },
              { prop: 'value', heading: 'Wert', type: 'number' },
              { prop: 'clanNames', heading: 'Clan' },
            ]}
            pagination={paginatedPlayers.pagination}
            pageSize={pageSize}
          />
        ) : (
          isError && <div className="center">Keine Spieler gefunden</div>
        )}
        <p>Momentan werden nur Statistiken der gesamten Zeit angezeigt. Das Filtern nach Zeiträumen wird in Zukunft möglich sein.</p>
      </section>
    </main>
  )
}
