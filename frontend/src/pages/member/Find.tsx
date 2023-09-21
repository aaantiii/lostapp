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
import { useSearchParams } from 'react-router-dom'

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
  const [searchParams, setSearchParams] = useSearchParams()

  const [selectedSearchOption, setSelectedSearchOption] = useState(() => {
    for (const key of searchParams.keys()) {
      const option = searchOptionGroup.options.find((option) => option.value === key)
      if (option) return option
    }

    return searchOptionGroup.options[0]
  })
  const searchValue = searchParams.get(selectedSearchOption.value) ?? ''
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
        [selectedSearchOption.value]: searchValue,
      } satisfies PlayersParams,
    ],
    enabled: false,
    retry: false,
    cacheTime: 1000 * 60,
  })

  useEffect(() => {
    if (searchValue === '') return
    refetch()
  }, [page, searchParams])

  const handleCopyTag = useCallback((player: Player) => {
    navigator.clipboard.writeText(player.tag)
    sendMessage({
      message: `Tag von ${player.name} kopiert!`,
      type: 'success',
    })
  }, [])

  function handleOptionChange(value: string) {
    const option = searchOptionGroup.options.find((option) => option.value === value)
    if (!option) return

    setSelectedSearchOption(option)
  }

  function handleSubmit(newFormData: FindPlayerForm) {
    if (searchParams.get(newFormData.option) === newFormData.value) return
    newFormData.value = encodeURI(newFormData.value)
    setPage(1)
    setSearchParams({ [newFormData.option]: newFormData.value }, { replace: true })
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
              messages: [FormMessages.required],
            },
            {
              label: `Nach ${selectedSearchOption.displayText} suchen`,
              name: 'value',
              control: <Input type="search" placeholder={`Nach ${selectedSearchOption.displayText} suchen`} defaultValue={searchValue} />,
              messages: [
                selectedSearchOption.value === 'discordID' ? FormMessages.fixedLength(18) : FormMessages.minMaxLength(3, 30),
                FormMessages.required,
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
          <p>Keine Ergebnisse für {`${selectedSearchOption.displayText} "${searchValue}"`}</p>
        ) : (
          error?.response?.status === HttpStatusCode.BadRequest && <p>Ungültige Eingaben.</p>
        )}
      </section>
    </main>
  )
}
