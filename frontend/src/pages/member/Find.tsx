import { useCallback, useEffect, useRef, useState } from 'react'
import Input from '@components/Input'
import useDocumentTitle from '@hooks/useDocumentTitle'
import { SelectFormWrapper, SelectOptionGroup } from '@components/Select'
import { useQuery } from '@tanstack/react-query'
import { PlayersParams } from '@api/types/params'
import { PaginatedResponse } from '@api/types/pagination'
import { CardList, Card } from '@components/Card'
import ExperienceLevel from '@components/ExperienceLevel'
import { formatPlayerClanRoles, urlEncodeTag } from '@fmt/cocFormatter'
import Form from '@components/Form'
import Button from '@components/Button'
import LoadingSpinner from '@components/LoadingSpinner'
import Paginator from '@components/Paginator'
import { useMessage } from '@context/messageContext'
import routes from '@api/routes'
import FormMessages from '@/validation/formMessages'
import usePageSize from '@hooks/usePageSize'
import { AxiosError, HttpStatusCode } from 'axios'
import { Player } from '@api/types/player'

interface FindPlayerForm {
  option: string
  value: string
}

const searchOptionGroup: SelectOptionGroup = {
  title: 'Filter wählen',
  options: [
    { value: 'name', displayText: 'Name' },
    { value: 'tag', displayText: 'Tag' },
    { value: 'clanName', displayText: 'Clan Name' },
    { value: 'clanTag', displayText: 'Clan Tag' },
    { value: 'discordID', displayText: 'Discord ID' },
  ],
}

export default function Find() {
  const heading = useDocumentTitle('Mitglied suchen')
  const pageSize = usePageSize(10, 16)
  const { sendMessage } = useMessage()

  const [formData, setFormData] = useState<FindPlayerForm>({ option: '', value: '' })
  const [selectedSearchOption, setSelectedSearchOption] = useState(searchOptionGroup.options[0])
  const [page, setPage] = useState(1)

  const {
    data: players,
    refetch,
    isFetching,
    error,
  } = useQuery<PaginatedResponse<Player>, AxiosError>({
    queryKey: [
      routes.players.all,
      null,
      {
        pageSize,
        page,
        [formData.option]: formData.value,
      } satisfies PlayersParams,
    ],
    enabled: false,
    retry: false,
    cacheTime: 1000 * 60,
  })

  useEffect(() => {
    if (!formData.option || !formData.value) return

    refetch()
  }, [page, formData])

  const handleCopyTag = useCallback((player: Player) => {
    navigator.clipboard.writeText(player.tag)
    sendMessage({
      message: `Tag von ${player.name} kopiert!`,
      type: 'success',
    })
  }, []) // [sendMessage]

  function handleOptionChange(value: string) {
    const option = searchOptionGroup.options.find((option) => option.value === value)
    if (!option) return

    setSelectedSearchOption(option)
  }

  function handleSubmit(newFormData: FindPlayerForm) {
    if (newFormData.option === formData.option && newFormData.value === formData.value) return
    if (newFormData.option.toLowerCase().includes('tag')) newFormData.value = urlEncodeTag(newFormData.value)
    setPage(1)
    setFormData(newFormData)
  }

  return (
    <main>
      {heading}
      <section>
        <h2>Filter</h2>
        <Form
          onSubmit={handleSubmit}
          submitText="Suchen"
          isLoading={isFetching}
          fields={[
            {
              label: 'Filter',
              name: 'option',
              control: (
                <SelectFormWrapper
                  placeholder="Filter auswählen"
                  optionGroups={[searchOptionGroup]}
                  onChange={handleOptionChange}
                  defaultValue={selectedSearchOption.value}
                />
              ),
              messages: [FormMessages.valueMissing],
            },
            {
              label: `Nach ${selectedSearchOption.displayText} suchen`,
              name: 'value',
              control: <Input type="search" placeholder={`Nach ${selectedSearchOption.displayText} suchen`} />,
              messages: [
                selectedSearchOption.value === 'discordID' ? FormMessages.fixedLength(18) : FormMessages.minMaxLength(3, 30),
                FormMessages.valueMissing,
              ],
            },
          ]}
        ></Form>
      </section>
      <section>
        <h2>Suchergebnisse</h2>
        <Paginator pagination={players?.pagination} onPageChange={setPage} />
        {isFetching && <LoadingSpinner />}
        {players?.items && players.items.length > 0 && (
          <CardList>
            {players.items.map((player) => (
              <Card
                title={player.name}
                thumbnail={<ExperienceLevel level={player.expLevel} />}
                description={formatPlayerClanRoles(player)}
                key={player.tag}
                buttons={[
                  <Button to={`/member/${urlEncodeTag(player.tag)}`} key="view-player">
                    Spieler ansehen
                  </Button>,
                  <Button onClick={() => handleCopyTag(player)} key="copy-tag">
                    Tag kopieren
                  </Button>,
                ]}
              />
            ))}
          </CardList>
        )}
        {error?.response?.status === HttpStatusCode.NotFound ? (
          <p>Keine Ergebnisse für {`${selectedSearchOption.displayText} "${formData.value}"`}</p>
        ) : (
          error && <p>Es ist ein Fehler aufgetreten</p>
        )}
      </section>
    </main>
  )
}
