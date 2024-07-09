import { addDaysToDate, dateTimeFormatter, timeUntil } from '@/utils/intlFormatter'
import routes from '@api/routes'
import { PaginatedResponse } from '@api/types/base'
import { ApiError, ClanSettings, Kickpoint } from '@api/types/models'
import QueryState from '@components/QueryState'
import { TableSkeleton } from '@components/Skeletons'
import Table from '@components/Table'
import useTitle from '@hooks/useTitle'
import { useQuery } from '@tanstack/react-query'
import { useParams, useSearchParams } from 'react-router-dom'

export default function KickpointsIndex() {
  const title = useTitle('Kickpunkte')
  const { clanTag, memberTag } = useParams()
  const [searchParams] = useSearchParams({
    page: '1',
    limit: '10',
  })
  const {
    data: kickpoints,
    isLoading: kickpointsLoading,
    error: kickpointsError,
  } = useQuery<PaginatedResponse<Kickpoint>, ApiError>({
    retry: false,
    queryKey: [
      routes.clans.members.kickpoints.byClanMember,
      { clanTag, memberTag },
      {
        page: searchParams.get('page')!,
        limit: searchParams.get('limit')!,
      },
    ],
  })
  const {
    data: settings,
    isLoading: settingsLoading,
    error: settingsError,
  } = useQuery<ClanSettings, ApiError>({
    queryKey: [routes.clans.settings, { clanTag }],
  })

  return (
    <main>
      {title}
      {kickpoints?.items?.length && settings ? (
        <Table
          rawData={kickpoints}
          header={['ID', 'Beschreibung', 'Kickpunkte', 'Läuft ab in', 'Erhalten am', 'Erstellt von']}
          pagination={kickpoints.pagination}
        >
          {kickpoints.items.map((k) => (
            <tr>
              <td>{k.id}</td>
              <td>{k.description}</td>
              <td>{k.amount}</td>
              <td>{timeUntil(addDaysToDate(new Date(k.date), settings.kickpointsExpireAfterDays))}</td>
              <td>{dateTimeFormatter.format(new Date(k.createdAt))}</td>
              <td>{k.createdByUser.name}</td>
            </tr>
          ))}
        </Table>
      ) : (
        <QueryState loading={kickpointsLoading || settingsLoading} error={kickpointsError ?? settingsError} loader={<TableSkeleton />} />
      )}
    </main>
  )
}
