import FormMessages from '@/validation/formMessages'
import client from '@api/client'
import routes from '@api/routes'
import { Kickpoint, UpdateKickpoint } from '@api/types/kickpoint'
import { buildURL } from '@api/urlBuilder'
import FormDialog from '@components/FormDialog'
import Input from '@components/Input'
import TextArea from '@components/TextArea'
import { useMessage } from '@context/messageContext'
import { toIsoString } from '@fmt/intlFormatter'
import { HttpStatusCode } from 'axios'
import { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'

interface DialogEditProps {
  kickpoint: Kickpoint
  onSuccess: () => void
}

export default function DialogEdit({ kickpoint, onSuccess }: DialogEditProps) {
  const { clanTag, memberTag } = useParams()
  const { sendMessage } = useMessage()
  const [isSaving, setIsSaving] = useState(false)

  async function saveChanges(data: UpdateKickpoint) {
    setIsSaving(true)
    const { status } = await client.put(buildURL(routes.clans.members.kickpoints.byId, { clanTag, memberTag, kickpointId: -1 }), data)
    setIsSaving(false)

    if (status < 300) {
      sendMessage({ type: 'success', message: `Kickpunkt #${kickpoint.id} wurde erfolgreich geändert.` })
      onSuccess()
      return
    }

    switch (status) {
      case HttpStatusCode.BadRequest:
        sendMessage({ type: 'error', message: 'Kickpunkt konnte wegen ungültigen Eingaben nicht geändert werden.' })
        break
      case HttpStatusCode.NotFound:
        sendMessage({ type: 'error', message: `Kickpunkt #${kickpoint.id} wurde nicht gefunden.` })
        break
      default:
        sendMessage({ type: 'error', message: 'Beim Speichern der Änderungen ist ein Fehler aufgetreten.' })
        break
    }
  }

  return (
    <FormDialog
      title="Kickpunkt bearbeiten"
      description='Wenn du die Änderungen vorgenommen hast, klicke auf "Änderungen speichern".'
      isLoading={isSaving}
      fields={[
        {
          name: 'amount',
          label: 'Anzahl Kickpunkte',
          control: <Input placeholder="Kickpunkte" type="number" inputMode="numeric" defaultValue={kickpoint.amount} />,
          messages: [FormMessages.required, FormMessages.minMaxNumber(1, 10)],
          type: 'number',
        },
        {
          name: 'description',
          label: 'Beschreibung',
          control: <TextArea placeholder='z.B. "CKL September"' defaultValue={kickpoint.description} />,
          messages: [FormMessages.required, FormMessages.minMaxLength(5, 100)],
        },
        {
          name: 'date',
          label: 'Datum',
          control: <Input placeholder="Datum" type="date" defaultValue={toIsoString(new Date(kickpoint.date)).substring(0, 10)} />,
          messages: [FormMessages.required, FormMessages.maxDate(new Date())],
          type: 'date',
        },
      ]}
      submitText="Änderungen speichern"
      onSubmit={saveChanges}
    />
  )
}
