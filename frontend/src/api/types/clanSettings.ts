import { UpdatedByUser } from './user'

export interface ClanSettings extends UpdatedByUser {
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

export type UpdateClanSettings = Omit<ClanSettings, 'clanTag' | 'updatedAt' | 'updatedByUser'>
