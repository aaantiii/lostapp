export default {
  auth: {
    session: 'auth/session',
    login: 'auth/discord/login',
    logout: 'auth/discord/logout',
  },
  players: {
    all: 'players',
    byTag: 'players/:tag',
    leaderboard: 'players/leaderboard/:statsId',
    comparableStats: 'players/comparable-stats',
  },
  clans: {
    all: 'clans',
    byTag: 'clans/:clanTag',
    leading: 'clans/leading',
    settings: 'clans/:clanTag/settings',
    members: {
      kickpoints: {
        byClan: 'clans/:clanTag/members/kickpoints',
        byClanMember: 'clans/:clanTag/members/:memberTag/kickpoints',
        byId: 'clans/:clanTag/members/:memberTag/kickpoints/:kickpointId',
      },
    },
  },
  notifications: 'notifications',
} as const
