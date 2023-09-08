export interface ClanSettings {
  minSeasonWins: number
  maxKickpoints: number
  kickpointsExpireAfterDays: number
  kickpointsSeasonWins: number
  kickpointsCWMissed: number
  kickpointsCWFail: number
  kickpointsCWLMissed: number
  kickpointsCWLZeroStars: number
  kickpointsCWLOneStar: number
  kickpointsRaidFail: number
  kickpointsRaidMissed: number
  kickpointsClanGames: number
}

export type UpdatedClanSettings = Omit<ClanSettings, 'clanTag'>
