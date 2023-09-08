export function validateFindQuery(searchValue: string, searchOption: string): string | undefined {
  if (!searchOption) return 'Wähle einen Filter aus'
  if ((!searchValue || searchValue.length < 3) && searchOption !== 'discordID') return 'Der Suchbegriff muss mindestens 3 Zeichen lang sein.'
  if (searchOption === 'discordID' && (searchValue.length !== 18 || !searchValue.match(/\d{18}/))) return 'Ungültige Discord ID'
}
