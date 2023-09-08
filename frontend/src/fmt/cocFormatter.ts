import { ClanMemberRoleTranslated } from '../api/types/clan'
import { Player } from '../api/types/player'

export function formatPlayerClanRoles(player: Player): string[] {
  if (!player.clans || player.clans.length === 0) return ['kein Clan']

  const formattedRoles: string[] = []

  for (const clan of player.clans) {
    formattedRoles.push(`${ClanMemberRoleTranslated.get(clan.role)} in ${clan.name}\n`)
  }

  return formattedRoles
}

// urlEncodeCocTag removes the leading '#' from a COC tag.
// With a leading '#' the tag would be interpreted as anchor instead of route id in URLs.
export function urlEncodeTag(tag?: string): string {
  if (tag?.startsWith('#')) return tag.slice(1)
  return tag ?? ''
}

export function urlDecodeTag(tag?: string): string {
  if (tag?.startsWith('#')) return tag
  return tag ? `#${tag}` : ''
}
