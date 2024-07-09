import { useSearchParams } from 'react-router-dom'
import Grid from './Grid'
import SearchQuery from './SearchQuery'
import Spacer from './Spacer'

type GridSearchProps = {
  children: (q: string) => React.ReactNode
}

export default function GridSearch({ children }: GridSearchProps) {
  const [searchParams] = useSearchParams()
  return (
    <div className="GridSearch">
      <SearchQuery placeholder="Mitglied suchen" />
      <Spacer size="small" />
      <Grid>{children(searchParams.get('q') ?? '')}</Grid>
    </div>
  )
}
