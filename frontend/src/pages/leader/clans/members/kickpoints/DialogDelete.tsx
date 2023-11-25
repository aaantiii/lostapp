import client from '@api/client'
import routes from '@api/routes'
import { buildURL } from '@api/urlBuilder'
import ConfirmDialog from '@components/ConfirmDialog'
import { useMessage } from '@context/messageContext'
import { HttpStatusCode } from 'axios'
import { useState } from 'react'
import { useParams } from 'react-router-dom'

interface DialogDeleteProps {
  kickpointId: number
  onSuccess: () => void
}

export default function DialogDelete({ kickpointId, onSuccess }: DialogDeleteProps) {
  const { clanTag, memberTag } = useParams()
  const { sendMessage } = useMessage()
  const [isDeleting, setIsDeleting] = useState(false)

  async function deleteKickpoint() {
    setIsDeleting(true)
    const { status } = await client.delete(buildURL(routes.clans.members.kickpoints.byId, { clanTag, memberTag, kickpointId }))
    setIsDeleting(false)

    if (status < 300) {
      sendMessage({ type: 'success', message: `Kickpunkt #${kickpointId} wurde erfolgreich gelöscht.` })
      onSuccess()
      return
    }

    switch (status) {
      case HttpStatusCode.NotFound:
        sendMessage({ type: 'error', message: `Kickpunkt #${kickpointId} wurde nicht gefunden.` })
        break
      default:
        sendMessage({ type: 'error', message: 'Beim Löschen des Kickpunkts ist ein Fehler aufgetreten.' })
        break
    }
  }

  return (
    <ConfirmDialog
      title="Kickpunkt löschen"
      description="Möchtest du diesen Kickpunkt wirklich löschen?"
      confirmText="Ja"
      cancelText="Nein"
      onConfirm={deleteKickpoint}
      triggerButtonColor="red"
    />
  )
}
