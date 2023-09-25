import { Outlet, useParams } from 'react-router-dom'
import { useQuery } from '@tanstack/react-query'
import { AuthRole } from '../api/types/auth'
import ProtectedRoute from './ProtectedRoute'
import { MemberOutletContext } from '@context/types'
import LoadingScreen from '../components/LoadingScreen'
import { useAuth } from '../context/authContext'
import routes from '../api/routes'
import { useEffect, useState } from 'react'
import { urlDecodeTag } from '../fmt/cocFormatter'
import { Clan } from '../api/types/clan'
import { Player } from '../api/types/player'
import { ClanSettings } from '@/api/types/clanSettings'

export default function MemberRoute() {
  const { clanTag, memberTag } = useParams()
  const { discordUser } = useAuth()

  const { data: userPlayers, isLoading: userPlayersLoading } = useQuery<Player[]>({
    queryKey: [routes.players.all, null, { discordID: discordUser?.id }],
    enabled: discordUser !== undefined,
  })

  const { data: clan, isFetching: clanFetching } = useQuery<Clan>({
    queryKey: [routes.clans.byTag, { tag: clanTag }],
    enabled: clanTag !== undefined,
  })

  const [player, setPlayer] = useState<Player>()
  const {
    data: fetchedPlayer,
    refetch: fetchPlayer,
    isFetching: memberFetching,
  } = useQuery<Player>({
    queryKey: [routes.players.byTag, { tag: memberTag }],
    enabled: false,
  })

  const { data: clanSettings, isFetching: clanSettingsFetching } = useQuery<ClanSettings>({
    queryKey: [routes.clans.settings, { tag: clanTag }],
    enabled: clanTag !== undefined,
  })

  useEffect(() => {
    if (!memberTag || !userPlayers) return

    const player = userPlayers.find((account) => account.tag === urlDecodeTag(memberTag))
    if (player) setPlayer(player)
    else fetchPlayer()
  }, [memberTag, userPlayers])

  useEffect(() => {
    if (!fetchedPlayer) return
    setPlayer(fetchedPlayer)
  }, [fetchedPlayer])

  const isLoading = userPlayersLoading || clanFetching || memberFetching || clanSettingsFetching
  return (
    <ProtectedRoute requiredRole={AuthRole.Member}>
      {isLoading ? <LoadingScreen /> : <Outlet context={{ userPlayers, clan, player, clanSettings } satisfies MemberOutletContext} />}
    </ProtectedRoute>
  )
}
