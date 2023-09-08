import ExpandableList, { ExpandableListItem } from '../components/ExpandableList'
import { AuthRole } from '../api/types/auth'
import { useAuth } from '../context/authContext'
import { useMemo } from 'react'
import {
  faAddressCard,
  faBell,
  faChartLine,
  faChess,
  faDashboard,
  faHouseChimney,
  faRankingStar,
  faSearch,
  faShieldHalved,
} from '@fortawesome/free-solid-svg-icons'

export default function useNavItems() {
  const { userRole } = useAuth()

  const navItems = useMemo(() => {
    const items = [
      <ExpandableList title="Lost Clans" key="lost-clans">
        <ExpandableListItem title="Startseite" href="/" icon={faHouseChimney} />
      </ExpandableList>,
    ]

    if (userRole === undefined || userRole === AuthRole.User) {
      items.push(
        <ExpandableList title="Bewerbung" key="user">
          <ExpandableListItem disabled title="Übersicht" href="/user" icon={faDashboard} />
          <ExpandableListItem disabled title="Mitglied werden" href="/user/apply" icon={faAddressCard} />
        </ExpandableList>
      )

      return items
    }

    if (userRole >= AuthRole.Leader) {
      items.push(
        <ExpandableList title="Leader" key="leader">
          <ExpandableListItem title="Übersicht" href="/leader" icon={faDashboard} />
          <ExpandableListItem disabled title="Nachricht senden" href="/leader/notify" icon={faBell} />
        </ExpandableList>
      )
    }

    if (userRole >= AuthRole.Member) {
      items.push(
        <ExpandableList title="Member" key="member">
          <ExpandableListItem title="Übersicht" href="/member" icon={faDashboard} />
          <ExpandableListItem title="Member suchen" href="/member/find" icon={faSearch} />
          <ExpandableListItem disabled title="Statistiken" href="/member/stats" icon={faChartLine} />
          <ExpandableListItem title="Leaderboard" href="/member/leaderboard" icon={faRankingStar} />
          <ExpandableListItem title="1v1 Vergleich" href="/member/1v1" icon={faChess} />
          <ExpandableListItem title="Lost Clans" href="/member/clans" icon={faShieldHalved} />
        </ExpandableList>
      )
    }

    return items
  }, [userRole])

  return navItems
}
