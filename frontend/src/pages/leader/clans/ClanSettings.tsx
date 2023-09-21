import { useOutletContext } from 'react-router-dom'
import useDocumentTitle from '@hooks/useDocumentTitle'
import routes from '@api/routes'
import { useEffect, useState } from 'react'
import Form from '@components/Form'
import { UpdateClanSettings } from '@api/types/clanSettings'
import Input from '@components/Input'
import { useMessage } from '@context/messageContext'
import FormMessages from '@/validation/formMessages'
import { LeaderOutletContext } from '@context/types'
import DataChangelog from '@components/DataChangelog'
import client from '@api/client'
import { replaceRouteIds } from '@api/urlBuilder'
import { urlEncodeTag } from '@fmt/cocFormatter'

export default function ClanSettings() {
  const { clan, clanSettings, refreshClanSettings } = useOutletContext<LeaderOutletContext>()
  const heading = useDocumentTitle(clan ? `${clan.name} Einstellungen` : 'Clan Einstellungen')
  const { sendMessage } = useMessage()
  const [isSaving, setIsSaving] = useState(false)

  async function saveChanges(data: UpdateClanSettings) {
    if (!clan) return

    setIsSaving(true)
    const { status } = await client.put(replaceRouteIds(routes.clans.settings, { tag: urlEncodeTag(clan.tag) }), data)
    setIsSaving(false)

    if (status < 300) {
      sendMessage({ type: 'success', message: `Die Claneinstellungen wurden erfolgreich geändert.` })
      refreshClanSettings()
      return
    }

    switch (status) {
      case 400:
        sendMessage({ type: 'error', message: 'Die Claneinstellungen konnten wegen ungültigen Eingaben nicht geändert werden.' })
        break
      case 404:
        sendMessage({ type: 'error', message: `Der Clan "${clan.tag}" konnte nicht gefunden werden.` })
        break
      default:
        sendMessage({ type: 'error', message: 'Beim Speichern der Änderungen ist ein Fehler aufgetreten.' })
        break
    }
  }

  function handleSubmit(data: UpdateClanSettings) {
    if (!data || !clanSettings || !clan) return

    for (const k in data) {
      const key = k as keyof UpdateClanSettings
      if (data[key] !== clanSettings[key]) {
        saveChanges(data)
        return
      }
    }

    sendMessage({
      type: 'error',
      message: 'Es wurden keine Änderungen vorgenommen!',
    })
  }

  return (
    <main>
      {heading}
      <DataChangelog data={clanSettings} type="updated" />
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
              messages: [FormMessages.minMaxNumber(1, 100), FormMessages.needInteger, FormMessages.required],
              type: 'number',
            },
            {
              label: 'Season Wins (Minimum)',
              name: 'minSeasonWins',
              control: <Input type="number" inputMode="numeric" defaultValue={clanSettings.minSeasonWins} />,
              messages: [FormMessages.minMaxNumber(0, 200), FormMessages.needInteger, FormMessages.required],
              type: 'number',
            },
            {
              label: 'Tage bis zum Abbau',
              name: 'kickpointsExpireAfterDays',
              control: <Input type="number" inputMode="numeric" defaultValue={clanSettings.kickpointsExpireAfterDays} />,
              messages: [FormMessages.minMaxNumber(7, 100), FormMessages.needInteger, FormMessages.required],
              type: 'number',
            },
            {
              label: 'Kickpunkte für Season Wins',
              name: 'kickpointsSeasonWins',
              control: <Input type="number" inputMode="numeric" defaultValue={clanSettings.kickpointsSeasonWins} />,
              messages: [FormMessages.minMaxNumber(0, 10), FormMessages.needInteger, FormMessages.required],
              type: 'number',
            },
            {
              label: 'Kickpunkte für verpasste CK-Angriffe',
              name: 'kickpointsCWMissed',
              control: <Input type="number" inputMode="numeric" defaultValue={clanSettings.kickpointsCWMissed} />,
              messages: [FormMessages.minMaxNumber(0, 10), FormMessages.needInteger, FormMessages.required],
              type: 'number',
            },
            {
              label: 'Kickpunkte für CK-Fail',
              name: 'kickpointsCWFail',
              control: <Input type="number" inputMode="numeric" defaultValue={clanSettings.kickpointsCWFail} />,
              messages: [FormMessages.minMaxNumber(0, 10), FormMessages.needInteger, FormMessages.required],
              type: 'number',
            },
            {
              label: 'Kickpunkte für verpasste CKL-Angriffe',
              name: 'kickpointsCWLMissed',
              control: <Input type="number" inputMode="numeric" defaultValue={clanSettings.kickpointsCWLMissed} />,
              messages: [FormMessages.minMaxNumber(0, 10), FormMessages.needInteger, FormMessages.required],
              type: 'number',
            },
            {
              label: 'Kickpunkte für CKL 0 Sterne',
              name: 'kickpointsCWLZeroStars',
              control: <Input type="number" inputMode="numeric" defaultValue={clanSettings.kickpointsCWLZeroStars} />,
              messages: [FormMessages.minMaxNumber(0, 10), FormMessages.needInteger, FormMessages.required],
              type: 'number',
            },
            {
              label: 'Kickpunkte für CKL 1 Stern',
              name: 'kickpointsCWLOneStar',
              control: <Input type="number" inputMode="numeric" defaultValue={clanSettings.kickpointsCWLOneStar} />,
              messages: [FormMessages.minMaxNumber(0, 10), FormMessages.needInteger, FormMessages.required],
              type: 'number',
            },
            {
              label: 'Kickpunkte für verpasste Raids',
              name: 'kickpointsRaidMissed',
              control: <Input type="number" inputMode="numeric" defaultValue={clanSettings.kickpointsRaidMissed} />,
              messages: [FormMessages.minMaxNumber(0, 10), FormMessages.needInteger, FormMessages.required],
              type: 'number',
            },
            {
              label: 'Kickpunkte für Raid Fail',
              name: 'kickpointsRaidFail',
              control: <Input type="number" inputMode="numeric" defaultValue={clanSettings.kickpointsRaidFail} />,
              messages: [FormMessages.minMaxNumber(0, 10), FormMessages.needInteger, FormMessages.required],
              type: 'number',
            },
            {
              label: 'Kickpunkte für Clan Spiele',
              name: 'kickpointsClanGames',
              control: <Input type="number" inputMode="numeric" defaultValue={clanSettings.kickpointsClanGames} />,
              messages: [FormMessages.minMaxNumber(0, 10), FormMessages.needInteger, FormMessages.required],
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
