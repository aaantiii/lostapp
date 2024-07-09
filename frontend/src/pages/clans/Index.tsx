import { urlDecodeTag, urlEncodeTag } from '@/utils/cocFormatter'
import routes from '@api/routes'
import { PaginatedResponse } from '@api/types/base'
import { LiveClan } from '@api/types/coc'
import { ApiError } from '@api/types/models'
import { ClansParams } from '@api/types/params'
import Button from '@components/Button'
import Card from '@components/Card'
import Grid from '@components/Grid'
import Paginator from '@components/Paginator'
import QueryState from '@components/QueryState'
import { GridSkeleton } from '@components/Skeletons'
import useTitle from '@hooks/useTitle'
import { useQuery } from '@tanstack/react-query'
import { useSearchParams } from 'react-router-dom'

export default function ClansIndex() {
  const title = useTitle('LOST Clans')
  const [searchParams] = useSearchParams({
    page: '1',
    limit: '12',
  })

  const {
    data: clans,
    isLoading,
    error,
  } = useQuery<PaginatedResponse<LiveClan>, ApiError>({
    queryKey: [
      routes.clans.live.index,
      null,
      {
        page: searchParams.get('page')!,
        limit: searchParams.get('limit')!,
      } as ClansParams,
    ],
  })

  return (
    <main>
      {title}
      {(clans?.items?.length ?? 0) > 0 ? (
        <Paginator pagination={clans?.pagination} limits={[12, 24, 36, 48]}>
          <Grid>
            {clans?.items.map((clan) => (
              <Card
                key={clan.tag}
                title={clan.name}
                description={clan.tag}
                thumbnail={clan.badgeUrls.small}
                fields={[{ label: 'Mitglieder', value: `${clan.members} / 50` }]}
                buttons={[
                  <Button key="view" to={`/clans/${urlEncodeTag(clan.tag)}/members`}>
                    Mitglieder
                  </Button>,
                ]}
              />
            ))}
          </Grid>
        </Paginator>
      ) : (
        <QueryState loading={isLoading} error={error} loader={<GridSkeleton />} />
      )}
    </main>
  )
}
