import { SelectFormWrapper, SelectOption, SelectOptionGroup, selectOptionOther } from '@components/Select'
import { useMessage } from '@context/messageContext'
import { CreateKickpoint } from '@api/types/kickpoint'
import FormDialog from '@components/FormDialog'
import Input from '@components/Input'
import FormMessages from '@/validation/formMessages'
import TextArea from '@components/TextArea'
import { useMemo, useState } from 'react'
import routes from '@api/routes'
import { useOutletContext, useParams } from 'react-router-dom'
import client from '@api/client'
import { replaceRouteIds } from '@api/urlBuilder'
import { HttpStatusCode } from 'axios'
import { LeaderOutletContext } from '@context/types'

interface DialogNewMemberProps {
  onSuccess: () => void
}

export default function DialogNew({ onSuccess }: DialogNewMemberProps) {
  const { clanTag, memberTag } = useParams()
  const { sendMessage } = useMessage()
  const reasonOptions = useReasonSelectOptions()
  const [selectedReasonOption, setSelectedReasonOption] = useState<undefined | SelectOption>(reasonOptions.options[0])
  const [isCreating, setIsCreating] = useState(false)

  async function createKickpoint(data: CreateKickpoint) {
    setIsCreating(true)
    const { status } = await client.post(replaceRouteIds(routes.clans.members.kickpoints.byTag, { clanTag, memberTag }), data)
    setIsCreating(false)

    switch (status) {
      case HttpStatusCode.Created:
        sendMessage({ type: 'success', message: `Kickpunkt wurde erfolgreich hinzugefügt.` })
        onSuccess()
        break
      case HttpStatusCode.BadRequest:
        sendMessage({ type: 'error', message: 'Kickpunkt konnte wegen ungültigen Eingaben nicht hinzugefügt werden.' })
        break
      case HttpStatusCode.NotFound:
        sendMessage({ type: 'error', message: `Mitglied "${memberTag}" wurde nicht gefunden.` })
        break
      default:
        sendMessage({ type: 'error', message: 'Beim Speichern der Änderungen ist ein Fehler aufgetreten.' })
        break
    }
  }

  function handleSubmit(data: CreateKickpoint) {
    if (!selectedReasonOption) return

    if (selectedReasonOption.value !== selectOptionOther('').value) {
      data.amount = parseInt(selectedReasonOption.realValue)
    }

    if (data.amount <= 0) {
      sendMessage({
        type: 'warning',
        message: 'Kickpunkt wurde nicht gespeichert da die Anzahl 0 ist.',
      })
      return
    }

    createKickpoint(data)
  }

  return (
    <FormDialog
      title="Neuer Kickpunkt"
      description="Fülle das Formular aus, um einen neuen Kickpunkt hinzuzufügen."
      isLoading={isCreating}
      fields={[
        {
          name: 'reason',
          label: 'Grund',
          control: (
            <SelectFormWrapper
              optionGroups={[reasonOptions]}
              type="number"
              placeholder="Grund wählen"
              defaultValue={selectedReasonOption?.value}
              onChange={(value) => setSelectedReasonOption(reasonOptions.options.find((option) => option.value === value))}
            />
          ),
          messages: [FormMessages.required],
          type: 'number',
          noSubmit: true,
        },
        selectedReasonOption?.value === selectOptionOther('').value
          ? {
              name: 'amount',
              label: 'Anzahl Kickpunkte',
              control: <Input placeholder="Kickpunkte" type="number" inputMode="numeric" />,
              messages: [FormMessages.required, FormMessages.minMaxNumber(1, 10)],
              type: 'number',
            }
          : null,
        {
          name: 'description',
          label: 'Beschreibung',
          control: <TextArea placeholder='z.B. "CKL September"' />,
          messages: [FormMessages.required, FormMessages.minMaxLength(5, 100)],
        },
        {
          name: 'date',
          label: 'Datum',
          control: <Input placeholder="Datum" type="date" />,
          messages: [FormMessages.required, FormMessages.maxDate(new Date())],
          type: 'date',
        },
      ]}
      submitText="Speichern"
      onSubmit={handleSubmit}
    />
  )
}

function useReasonSelectOptions(): SelectOptionGroup {
  const { clanSettings } = useOutletContext<LeaderOutletContext>()

  const options = useMemo<SelectOptionGroup>(() => {
    if (!clanSettings) return { title: 'Grund', options: [] }

    return {
      title: 'Grund wählen',
      options: [
        { value: 'kickpointsSeasonWins', displayText: 'Season Wins', realValue: clanSettings.kickpointsSeasonWins },
        { value: 'kickpointsCWMissed', displayText: 'CK-Angriff nicht gemacht', realValue: clanSettings.kickpointsCWMissed },
        { value: 'kickpointsCWFail', displayText: 'CK-Fail', realValue: clanSettings.kickpointsCWFail },
        { value: 'kickpointsCWLMissed', displayText: 'CKL-Angriff nicht gemacht', realValue: clanSettings.kickpointsCWLMissed },
        { value: 'kickpointsCWLZeroStars', displayText: 'CKL 0 Sterne Angriff', realValue: clanSettings.kickpointsCWLZeroStars },
        { value: 'kickpointsCWLOneStar', displayText: 'CKL 1 Sterne Angriff', realValue: clanSettings.kickpointsCWLOneStar },
        { value: 'kickpointsRaidFail', displayText: 'Raid Fail', realValue: clanSettings.kickpointsRaidFail },
        { value: 'kickpointsRaidMissed', displayText: 'Raid nicht gemacht', realValue: clanSettings.kickpointsRaidMissed },
        { value: 'kickpointsClanGames', displayText: 'Clan Spiele nicht gemacht', realValue: clanSettings.kickpointsClanGames },
        selectOptionOther('Anderer Grund'),
      ],
    }
  }, [clanSettings])

  return options
}
