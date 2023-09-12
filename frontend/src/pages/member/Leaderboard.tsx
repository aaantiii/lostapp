import { useEffect, useState } from 'react'
import Select, { SelectOption, SelectOptionGroup } from '@components/Select'
import { Table } from '@components/Table'
import useDocumentTitle from '@hooks/useDocumentTitle'
import { useQuery } from '@tanstack/react-query'
import { PlayersParams } from '@api/types/params'
import LoadingScreen from '@components/LoadingScreen'
import { urlEncodeTag } from '@fmt/cocFormatter'
import routes from '@api/routes'
import { PaginatedResponse } from '@api/types/pagination'
import UnsubmittableForm from '@components/UnsubmittableForm'
import usePageSize from '@hooks/usePageSize'
import { ComparableStatistic, PlayerStatistic } from '@api/types/playerStats'
import { Clan } from '@api/types/clan'

const clanTagFilterAll: SelectOption = { value: 'all', displayText: 'Alle Lost Clans' }

export default function Leaderboard() {
  const heading = useDocumentTitle('Leaderboard')

  const [page, setPage] = useState(1)
  const pageSize = usePageSize(20, 30)
  const [clansSelectOptions, setClanOptionGroup] = useState({ options: [] } as SelectOptionGroup)
  const [statisticsSelectOptions, setAchievementOptionGroup] = useState({ options: [] } as SelectOptionGroup)
  const [selectedClan, setSelectedClan] = useState(clanTagFilterAll.value)
  const [selectedStatistic, setSelectedStatistic] = useState<ComparableStatistic>()

  const { data: clans, isLoading: clansLoading } = useQuery<Clan[]>({
    queryKey: [routes.clans.all, null, { minifyData: true }],
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
      { statsId: selectedStatistic?.id },
      {
        page,
        pageSize,
        clanTag: selectedClan === clanTagFilterAll.value ? '' : urlEncodeTag(selectedClan),
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

    setAchievementOptionGroup({
      title: 'Achievement wählen',
      options: comparableStatistics.map((achievement) => ({
        value: achievement.name,
        displayText: achievement.displayName,
      })),
    })
  }, [comparableStatistics])

  useEffect(() => {
    if (!selectedStatistic || !selectedClan) return

    fetchLeaderboardPlayers()
  }, [selectedStatistic, selectedClan, page, pageSize])

  useEffect(() => {
    setPage(1)
  }, [selectedStatistic, selectedClan])

  if (clansLoading || achievementsLoading) return <LoadingScreen />

  return (
    <main>
      {heading}
      <section>
        <h2>Filter</h2>
        <UnsubmittableForm>
          <Select defaultValue={selectedClan} optionGroups={[clansSelectOptions]} onChange={setSelectedClan} placeholder="Clan auswählen" />
          <Select
            optionGroups={[statisticsSelectOptions]}
            onChange={(name) => setSelectedStatistic(comparableStatistics?.find((s) => s.name === name))}
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
            onPageChange={setPage}
            pagination={paginatedPlayers.pagination}
          />
        ) : (
          isError && <div className="center">Keine Spieler gefunden</div>
        )}
        <p>Momentan werden nur Statistiken der gesamten Zeit angezeigt. Das Filtern nach Zeiträumen wird in Zukunft möglich sein.</p>
      </section>
    </main>
  )
}
