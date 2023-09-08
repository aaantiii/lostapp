import { Session } from '../api/types/auth'
import { MessageProps } from '../components/Messages'
import { ClanSettings } from '../api/types/clanSettings'
import { Clan } from '../api/types/clan'
import { Player } from '../api/types/player'

export interface AuthContext extends Session {
  refreshSession: () => void
  logout: () => void
}

export interface MessageContext {
  messages: MessageProps[]
  sendMessage: (message: Omit<MessageProps, 'id'>) => void
  removeMessage: (id: number) => void
}

export interface MemberOutletContext {
  userPlayers?: Player[]
  clan?: Clan
  player?: Player
}

export interface LeaderOutletContext {
  leadingClans?: Clan[]
  clan?: Clan
  player?: Player
  clanSettings?: ClanSettings
  refreshClanSettings: () => void
}
