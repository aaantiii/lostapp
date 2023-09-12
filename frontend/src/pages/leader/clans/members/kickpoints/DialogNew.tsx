import { SelectFormWrapper, SelectOptionGroup } from '@components/Select'
import { useMessage } from '@context/messageContext'
import { LeaderOutletContext } from '@context/types'
import { CreateKickpoint } from '@api/types/kickpoint'
import Dialog from '@components/Dialog'
import { useOutletContext } from 'react-router-dom'
import { useMemo } from 'react'
import Input from '@components/Input'

interface DialogNewMemberProps {}

function useReasonOptions() {
  const { clanSettings } = useOutletContext<LeaderOutletContext>()

  const options = useMemo<SelectOptionGroup>(() => {
    return {
      title: 'Grund wählen',
      options: [
        { value: 'kickpointsSeasonWins', displayText: 'Season Wins' },
        { value: 'kickpointsCWMissed', displayText: 'CK-Angriff nicht gemacht' },
        { value: 'kickpointsCWFail', displayText: 'CK-Fail' },
        { value: 'kickpointsCWLMissed', displayText: 'CKL-Angriff nicht gemacht' },
        { value: 'kickpointsCWLZeroStars', displayText: 'CKL 0 Sterne Angriff' },
        { value: 'kickpointsCWLOneStar', displayText: 'CKL 1 Sterne Angriff' },
        { value: 'kickpointsRaidFail', displayText: 'Raid Fail' },
        { value: 'kickpointsRaidMissed', displayText: 'Raid nicht gemacht' },
        { value: 'kickpointsClanGames', displayText: 'Clan Spiele nicht gemacht' },
        { value: 'other', displayText: 'Sonstiges' },
      ],
    }
  }, [clanSettings])

  return options
}

const selectOptionGroupReason: SelectOptionGroup = {
  title: 'Grund wählen',
  options: [
    { value: 'kickpointsSeasonWins', displayText: 'Season Wins' },
    { value: 'kickpointsCWMissed', displayText: 'CK-Angriff nicht gemacht' },
    { value: 'kickpointsCWFail', displayText: 'CK-Fail' },
    { value: 'kickpointsCWLMissed', displayText: 'CKL-Angriff nicht gemacht' },
    { value: 'kickpointsCWLZeroStars', displayText: 'CKL 0 Sterne Angriff' },
    { value: 'kickpointsCWLOneStar', displayText: 'CKL 1 Sterne Angriff' },
    { value: 'kickpointsRaidFail', displayText: 'Raid Fail' },
    { value: 'kickpointsRaidMissed', displayText: 'Raid nicht gemacht' },
    { value: 'kickpointsClanGames', displayText: 'Clan Spiele nicht gemacht' },
    { value: 'other', displayText: 'Sonstiges' },
  ],
}

export default function DialogNew() {
  const { clanSettings } = useOutletContext<LeaderOutletContext>()
  const { sendMessage } = useMessage()

  const selectOptionGroupReason1 = useMemo<SelectOptionGroup | undefined>(() => {
    if (!clanSettings) return

    return {
      title: 'Grund wählen',
      options: [
        { value: 'kickpointsSeasonWins', displayText: 'Season Wins' },
        { value: 'kickpointsCWMissed', displayText: 'CK-Angriff nicht gemacht' },
        { value: 'kickpointsCWFail', displayText: 'CK-Fail' },
        { value: 'kickpointsCWLMissed', displayText: 'CKL-Angriff nicht gemacht' },
        { value: 'kickpointsCWLZeroStars', displayText: 'CKL 0 Sterne Angriff' },
        { value: 'kickpointsCWLOneStar', displayText: 'CKL 1 Sterne Angriff' },
        { value: 'kickpointsRaidFail', displayText: 'Raid Fail' },
        { value: 'kickpointsRaidMissed', displayText: 'Raid nicht gemacht' },
        { value: 'kickpointsClanGames', displayText: 'Clan Spiele nicht gemacht' },
        { value: 'other', displayText: 'Sonstiges' },
      ],
    }
  }, [clanSettings])

  function handleSubmit(data: CreateKickpoint) {}

  return (
    <Dialog
      title="Neuer Kickpunkt"
      description="Fülle das Formular aus, um einen neuen Kickpunkt hinzuzufügen."
      fields={[
        {
          name: 'reason',
          label: 'Grund',
          control: (
            <SelectFormWrapper optionGroups={[selectOptionGroupReason]} type="text" placeholder="Grund wählen" inputPlaceholder="Anzahl Kickpunkte" />
          ),
        },
        {
          name: 'description',
          label: 'Beschreibung',
          control: <Input placeholder='z.B. "CKL September"' />,
        },
      ]}
      submitText="Speichern"
      onSubmit={handleSubmit}
    />
  )
}
