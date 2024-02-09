import { ClanMemberRoleTranslated, PlayerClan } from '@api/types/coc'
import { Member } from '@api/types/models'

export function formatPlayerClanRole(clan?: PlayerClan) {
  if (!clan) return ''
  return `${ClanMemberRoleTranslated.get(clan.role)} in ${clan.name}`
}

export function formatPlayerClanRoles(clans: PlayerClan[]) {
  return clans.map(formatPlayerClanRole).join('\n')
}

export function formatMemberClan(member: Member) {
  if (!member.clan) return ''
  return `${ClanMemberRoleTranslated.get(member.clanRole)} in ${member.clan.name}`
}

export function urlEncodeTag(tag?: string) {
  if (tag?.startsWith('#')) return tag.slice(1)
  return tag ?? ''
}

export function urlDecodeTag(tag?: string) {
  if (!tag) return ''
  if (tag?.startsWith('#')) return tag
  return '#' + tag
}
