import routes from '@api/routes'
import { PaginatedResponse } from '@api/types/base'
import { useQuery } from '@tanstack/react-query'
import { useParams } from 'react-router-dom'

export default function ClanByTag() {
  const { clanTag } = useParams()

  return (
    <main>
      <ClanEvents />
    </main>
  )
}

function ClanEvents() {
  const { clanTag } = useParams()

  const {
    data: events,
    isLoading,
    error,
  } = useQuery<PaginatedResponse<ClanEvent>>({
    queryKey: [routes.clans.events, { clanTag }],
  })

  if (!isLoading && !events) return null

  return (
    <section>
      <h2>Events</h2>
      {}
    </section>
  )
}
