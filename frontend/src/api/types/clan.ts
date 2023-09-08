export interface Clan {
  memberList: ClanMember[]
  warTies?: number
  description?: string
  warLosses?: number
  clanCapital?: ClanCapital
  warLeague?: WarLeague
  badgeUrl: string
  tag: string
  name: string
  level: number
  warWins: number
  members: number
  isWarLogPublic: boolean
}

export enum ClanMemberRole {
  Leader = 'leader',
  CoLeader = 'coLeader',
  Elder = 'admin',
  Member = 'member',
}

export const ClanMemberRoleTranslated: ReadonlyMap<ClanMemberRole, string> = new Map([
  [ClanMemberRole.Leader, 'Anführer'],
  [ClanMemberRole.CoLeader, 'Vize Anführer'],
  [ClanMemberRole.Elder, 'Ältester'],
  [ClanMemberRole.Member, 'Mitglied'],
])

export interface ClanMember {
  tag: string
  name: string
  role: ClanMemberRole
  expLevel: number
}

export interface ClanCapital {
  CapitalHallLevel?: number
  districts?: { id: number; name: string; districtHallLevel: number }[]
}

export interface ClanWar {
  state: 'notInWar' | 'preparation' | 'inWar' | 'warEnded'
  teamSize: number
  startTime: string
  preparationStartTime: string
  endTime: string
  clan: WarClan
  opponent: WarClan
  attacksPerMember?: number
}

export interface WarClan {
  tag: string
  name: string
  badgeUrl: string
  clanLevel: number
  attacks: number
  stars: number
  destructionPercentage: number
  members: ClanWarMember[]
}

export interface ClanWarMember {
  tag: string
  name: string
  mapPosition: number
  townhallLevel: number
  opponentAttacks: number
  bestOpponentAttack?: ClanWarAttack
  attacks?: ClanWarAttack[]
}

export interface ClanWarAttack {
  order: number
  attackerTag: string
  defenderTag: string
  stars: number
  duration: number
  destructionPercentage: number
}

export interface WarLogClan {
  tag?: string
  name?: string
  badgeUrl: string
  clanLevel: number
  attacks?: number
  stars: number
  destructionPercentage: number
  expEarned?: number
}

export interface ClanWarLogEntry {
  result: 'win' | 'lose' | 'tie' | null
  endTime: string
  teamSize: number
  attacksPerMember?: number
  clan: WarLogClan
  opponent: WarLogClan
}

export interface ClanWarLog {
  items: ClanWarLogEntry[]
}

export interface ClanWarLeagueGroup {
  state: 'notInWar' | 'preparation' | 'inWar' | 'ended'
  season: string
  clans: ClanWarLeagueClan[]
  rounds: ClanWarLeagueRound[]
}

export interface ClanWarLeagueClan {
  name: string
  tag: string
  clanLevel: number
  badgeUrl: string
  members: ClanWarLeagueClanMember[]
}

export interface ClanWarLeagueClanMember {
  name: string
  tag: string
  townHallLevel: number
}

export interface ClanWarLeagueRound {
  warTags: string[]
}

export interface CapitalRaidSeason {
  state: 'ongoing' | 'ended'
  startTime: string
  endTime: string
  CapitalTotalLoot: number
  raidsCompleted: number
  totalAttacks: number
  enemyDistrictsDestroyed: number
  offensiveReward: number
  defensiveReward: number
  members?: CapitalRaidSeasonMember[]
  attackLog: CapitalRaidSeasonAttackLog[]
  defenseLog: CapitalRaidSeasonDefenseLog[]
}

export interface CapitalRaidSeasonMember {
  tag: string
  name: string
  attacks: number
  attackLimit: number
  bonusAttackLimit: number
  CapitalResourcesLooted: number
}

export interface CapitalRaidSeasonClan {
  tag: string
  name: string
  level: number
  badgeUrls: {
    small: string
    large: string
    medium: string
  }
}

export interface CapitalRaidSeasonDistrict {
  id: number
  name: string
  districtHallLevel: number
  destructionPercent: number
  attackCount: number
  totalLooted: number
}

export interface CapitalRaidSeasonAttackLog {
  defender: CapitalRaidSeasonClan
  attackCount: number
  districtCount: number
  districtsDestroyed: number
  districts: CapitalRaidSeasonDistrict[]
}

export interface CapitalRaidSeasonDefenseLog {
  attacker: CapitalRaidSeasonClan
  attackCount: number
  districtCount: number
  districtsDestroyed: number
  districts: CapitalRaidSeasonDistrict[]
}

export interface CapitalRaidSeasons {
  items: CapitalRaidSeason[]
}

export interface ClanRanking {
  clanLevel: number
  clanPoints: number
  location: Location
  members: number
  tag: string
  name: string
  rank: number
  previousRank: number
  badgeUrl: string
}

export interface ClanVersusRanking {
  clanLevel: number
  location: Location
  members: number
  tag: string
  name: string
  rank: number
  previousRank: number
  badgeUrl: string
  clanVersusPoints: number
}

export interface ClanCapitalRanking {
  clanLevel: number
  clanPoints: number
  location: Location
  members: number
  tag: string
  name: string
  rank: number
  previousRank: number
  badgeUrl: string
  clanCapitalPoints: number
}

export interface ClanCapitalRankingList {
  items: ClanCapitalRanking[]
}

export interface CapitalLeague {
  id: number
  name: string
}

export interface CapitalLeagueList {
  items: CapitalLeague[]
}

export interface LeagueSeasonList {
  items: {
    id: string
  }[]
}
export interface WarLeague {
  id: number
  name: string
}
