import { Outlet, useParams } from 'react-router-dom'
import { AuthRole } from '@api/types/auth'
import ProtectedRoute from './ProtectedRoute'
import { useQuery } from '@tanstack/react-query'
import routes from '@api/routes'
import LoadingScreen from '@components/LoadingScreen'
import { LeaderOutletContext } from '@context/types'
import { urlDecodeTag } from '@fmt/cocFormatter'
import { ClanSettings } from '@api/types/clanSettings'
import { Clan } from '@api/types/clan'
import { Player } from '@api/types/player'

export default function LeaderRoute() {
  const { clanTag, memberTag } = useParams()

  const { data: leadingClans, isLoading: leadingClansLoading } = useQuery<Clan[]>({
    queryKey: [routes.clans.leading],
    cacheTime: Infinity,
  })

  const clan = leadingClans?.find((clan) => clan.tag === urlDecodeTag(clanTag))

  const { data: player, isFetching: memberFetching } = useQuery<Player>({
    queryKey: [routes.players.byTag, { tag: memberTag }],
    enabled: memberTag !== undefined,
  })

  const {
    data: clanSettings,
    refetch: refreshClanSettings,
    isFetching: clanSettingsFetching,
  } = useQuery<ClanSettings>({
    queryKey: [routes.clans.settings, { clanTag }],
    enabled: clanTag !== undefined,
    staleTime: 1000 * 60,
    cacheTime: 1000 * 60,
  })

  const isLoading = leadingClansLoading || memberFetching || clanSettingsFetching
  return (
    <ProtectedRoute requiredRole={AuthRole.Leader}>
      {isLoading ? (
        <LoadingScreen />
      ) : (
        <Outlet context={{ leadingClans, clan, clanSettings, player, refreshClanSettings } satisfies LeaderOutletContext} />
      )}
    </ProtectedRoute>
  )
}
