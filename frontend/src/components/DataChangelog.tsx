import '@styles/components/DataChangelog.scss'
import { CreatedByUser, UpdatedByUser } from '@api/types/user'
import UserAvatar from './UserAvatar'
import { dateTimeFormatter } from '@/utils/intlFormatter'

interface DataChangelogProps {
  data?: UpdatedByUser | CreatedByUser
  type: 'created' | 'updated' | 'createdOrUpdated'
}

export default function DataChangelog({ data, type }: DataChangelogProps) {
  if (!data) return null

  function updatedChangelog(updatedData: UpdatedByUser) {
    if (!updatedData.updatedAt || !updatedData.updatedByUser) return null
    return (
      <div className="DataChangelog">
        <UserAvatar user={updatedData.updatedByUser} noHover nameFirst></UserAvatar>
        <span>Zuletzt geändert: {dateTimeFormatter.format(new Date(updatedData.updatedAt))}</span>
      </div>
    )
  }

  function createdChangelog(createdData: CreatedByUser) {
    if (!createdData.createdAt || !createdData.createdByUser) return null
    return (
      <div className="DataChangelog">
        <UserAvatar user={createdData.createdByUser} noHover nameFirst></UserAvatar>
        <span>Erstellt am: {dateTimeFormatter.format(new Date(createdData.createdAt))}</span>
      </div>
    )
  }

  const createdByUser = data as CreatedByUser
  const updatedByUser = data as UpdatedByUser

  if (type === 'updated') return updatedChangelog(updatedByUser)
  if (type === 'created') return createdChangelog(createdByUser)
  if (type === 'createdOrUpdated') return updatedChangelog(updatedByUser) ?? createdChangelog(createdByUser)

  return null
}
