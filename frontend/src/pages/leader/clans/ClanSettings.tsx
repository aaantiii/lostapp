import { useOutletContext } from 'react-router-dom'
import useDocumentTitle from '../../../hooks/useDocumentTitle'
import { useQuery } from '@tanstack/react-query'
import routes from '../../../api/routes'
import { useEffect, useState } from 'react'
import Form from '../../../components/Form'
import { UpdatedClanSettings } from '../../../api/types/clanSettings'
import Input from '../../../components/Input'
import { PUT } from '../../../api/queryFunctions'
import { useMessage } from '../../../context/messageContext'
import FormMessages from '../../../validation/formMessages'
import { LeaderOutletContext } from '../../../types/context'
import { urlEncodeTag } from '../../../fmt/cocFormatter'

export default function ClanSettings() {
  const { clan, clanSettings, refreshClanSettings } = useOutletContext<LeaderOutletContext>()
  const heading = useDocumentTitle(clan ? `${clan.name} Einstellungen` : 'Clan Einstellungen')
  const { sendMessage } = useMessage()
  const [updatedSettings, setUpdatedSettings] = useState<UpdatedClanSettings>()
  const { refetch: saveClanSettings, isFetching: isSaving } = useQuery({
    queryKey: [routes.clans.settings, { tag: urlEncodeTag(clan?.tag) }, updatedSettings],
    enabled: false,
    retry: false,
    queryFn: PUT<typeof updatedSettings>,
    cacheTime: 0,
    onSuccess: () => {
      sendMessage({
        type: 'success',
        message: 'Die Einstellungen wurden erfolgreich gespeichert.',
      })
      refreshClanSettings()
    },
    onError: () => {
      sendMessage({
        type: 'error',
        message: 'Beim Speichern der Einstellungen ist ein Fehler aufgetreten.',
      })
    },
  })

  useEffect(() => {
    if (!updatedSettings || !clanSettings || !clan) return

    for (const key in updatedSettings) {
      if (updatedSettings[key as keyof UpdatedClanSettings] !== clanSettings[key as keyof UpdatedClanSettings]) {
        saveClanSettings()
        return
      }
    }

    sendMessage({
      type: 'error',
      message: 'Es wurden keine Änderungen vorgenommen!',
    })
  }, [updatedSettings])

  function handleSubmit(data: UpdatedClanSettings) {
    setUpdatedSettings(data)
  }

  return (
    <main>
      {heading}
      <h2>Kickpunkte</h2>
      {clanSettings ? (
        <Form
          onSubmit={handleSubmit}
          submitText="Einstellungen Speichern"
          isLoading={isSaving}
          fields={[
            {
              label: 'Kickpunkte bis zum Kick',
              name: 'maxKickpoints',
              control: <Input type="number" inputMode="numeric" defaultValue={clanSettings.maxKickpoints} />,
              messages: [FormMessages.minMaxNumber(1, 100), FormMessages.needInteger, FormMessages.valueMissing],
              type: 'number',
            },
            {
              label: 'Season Wins (Minimum)',
              name: 'minSeasonWins',
              control: <Input type="number" inputMode="numeric" defaultValue={clanSettings.minSeasonWins} />,
              messages: [FormMessages.minMaxNumber(0, 200), FormMessages.needInteger, FormMessages.valueMissing],
              type: 'number',
            },
            {
              label: 'Tage bis zum Abbau',
              name: 'kickpointsExpireAfterDays',
              control: <Input type="number" inputMode="numeric" defaultValue={clanSettings.kickpointsExpireAfterDays} />,
              messages: [FormMessages.minMaxNumber(7, 100), FormMessages.needInteger, FormMessages.valueMissing],
              type: 'number',
            },
            {
              label: 'Kickpunkte für Season Wins',
              name: 'kickpointsSeasonWins',
              control: <Input type="number" inputMode="numeric" defaultValue={clanSettings.kickpointsSeasonWins} />,
              messages: [FormMessages.minMaxNumber(0, 10), FormMessages.needInteger, FormMessages.valueMissing],
              type: 'number',
            },
            {
              label: 'Kickpunkte für verpasste CK-Angriffe',
              name: 'kickpointsCWMissed',
              control: <Input type="number" inputMode="numeric" defaultValue={clanSettings.kickpointsCWMissed} />,
              messages: [FormMessages.minMaxNumber(0, 10), FormMessages.needInteger, FormMessages.valueMissing],
              type: 'number',
            },
            {
              label: 'Kickpunkte für CK-Fail',
              name: 'kickpointsCWFail',
              control: <Input type="number" inputMode="numeric" defaultValue={clanSettings.kickpointsCWFail} />,
              messages: [FormMessages.minMaxNumber(0, 10), FormMessages.needInteger, FormMessages.valueMissing],
              type: 'number',
            },
            {
              label: 'Kickpunkte für verpasste CKL-Angriffe',
              name: 'kickpointsCWLMissed',
              control: <Input type="number" inputMode="numeric" defaultValue={clanSettings.kickpointsCWLMissed} />,
              messages: [FormMessages.minMaxNumber(0, 10), FormMessages.needInteger, FormMessages.valueMissing],
              type: 'number',
            },
            {
              label: 'Kickpunkte für CKL 0 Sterne',
              name: 'kickpointsCWLZeroStars',
              control: <Input type="number" inputMode="numeric" defaultValue={clanSettings.kickpointsCWLZeroStars} />,
              messages: [FormMessages.minMaxNumber(0, 10), FormMessages.needInteger, FormMessages.valueMissing],
              type: 'number',
            },
            {
              label: 'Kickpunkte für CKL 1 Stern',
              name: 'kickpointsCWLOneStar',
              control: <Input type="number" inputMode="numeric" defaultValue={clanSettings.kickpointsCWLOneStar} />,
              messages: [FormMessages.minMaxNumber(0, 10), FormMessages.needInteger, FormMessages.valueMissing],
              type: 'number',
            },
            {
              label: 'Kickpunkte für verpasste Raids',
              name: 'kickpointsRaidMissed',
              control: <Input type="number" inputMode="numeric" defaultValue={clanSettings.kickpointsRaidMissed} />,
              messages: [FormMessages.minMaxNumber(0, 10), FormMessages.needInteger, FormMessages.valueMissing],
              type: 'number',
            },
            {
              label: 'Kickpunkte für Raid Fail',
              name: 'kickpointsRaidFail',
              control: <Input type="number" inputMode="numeric" defaultValue={clanSettings.kickpointsRaidFail} />,
              messages: [FormMessages.minMaxNumber(0, 10), FormMessages.needInteger, FormMessages.valueMissing],
              type: 'number',
            },
            {
              label: 'Kickpunkte für Clan Spiele',
              name: 'kickpointsClanGames',
              control: <Input type="number" inputMode="numeric" defaultValue={clanSettings.kickpointsClanGames} />,
              messages: [FormMessages.minMaxNumber(0, 10), FormMessages.needInteger, FormMessages.valueMissing],
              type: 'number',
            },
          ]}
        ></Form>
      ) : (
        <p>Fehler beim Laden der Claneinstellungen.</p>
      )}
    </main>
  )
}
