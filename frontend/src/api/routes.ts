export default {
  auth: {
    session: '/auth/session',
    login: '/auth/discord/login',
    logout: '/auth/discord/logout',
  },
  users: {
    byId: '/users/:id',
  },
  players: {
    index: '/players',
    byTag: '/players/:tag',
    stats: {
      list: '/players/stats/list',
      leaderboard: '/players/stats/leaderboard/:statName',
    },
    live: {
      index: '/players/live',
      byTag: '/players/:tag/live',
    },
  },
  clans: {
    index: '/clans',
    byTag: '/clans/:clanTag',
    leading: '/clans/leading',
    settings: '/clans/:clanTag/settings',
    members: {
      kickpoints: {
        byClan: '/clans/:clanTag/members/kickpoints',
        byClanMember: '/clans/:clanTag/members/:memberTag/kickpoints',
        byId: '/clans/:clanTag/members/:memberTag/kickpoints/:kickpointId',
      },
    },
  },
} as const

export function replaceIds(path: string, ids?: any): string {
  if (ids) {
    for (const [prop, value] of Object.entries<string | number | undefined>(ids)) {
      if (value === undefined) throw new Error(`invalid id: ${prop} is ${value}`)
      path = path.replace(`:${prop}`, value.toString())
    }
  }

  return encodeURI(path)
}

export function buildURI(path: string) {
  return encodeURI(import.meta.env.VITE_API_URL + path)
}
