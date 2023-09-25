import { ClanMemberRoleTranslated } from '@api/types/clan'
import { PlayerClan } from '@api/types/player'

export function formatPlayerClanRole(clan?: PlayerClan): string {
  if (!clan) return ''
  return `${ClanMemberRoleTranslated.get(clan.role)} in ${clan.name}`
}

export function formatPlayerClanRoles(clans: PlayerClan[]): string {
  return clans.map(formatPlayerClanRole).join('\n')
}

export function urlEncodeTag(tag?: string): string {
  if (tag?.startsWith('#')) return tag.slice(1)
  return tag ?? ''
}

export function urlDecodeTag(tag?: string): string {
  if (tag?.startsWith('#')) return tag
  return tag ? `#${tag}` : ''
}
