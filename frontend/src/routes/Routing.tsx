import { Navigate, Route, Routes } from 'react-router-dom'
import pages from '@pages/pages'
import LeaderRoute from './LeaderRoute'
import MemberRoute from './MemberRoute'

export default function Routing() {
  return (
    <Routes>
      <Route index element={<pages.Home />} />
      <Route path="report-bug" element={<pages.ReportBug />} />
      <Route path="request-feature" element={<pages.RequestFeature />} />
      <Route path="apply" element={<pages.Apply />} />

      <Route path="error/:code" element={<pages.Error />} />

      <Route path="legal">
        <Route path="imprint" element={<pages.legal.Imprint />} />
        <Route path="privacy" element={<pages.legal.Privacy />} />
      </Route>

      <Route path="auth/login">
        <Route index element={<pages.auth.login.Index />} />
        <Route path="success" element={<pages.auth.login.Success />} />
        <Route path="failed" element={<pages.auth.login.Failed />} />
      </Route>

      <Route path="member" element={<MemberRoute />}>
        <Route index element={<pages.member.Index />} />
        <Route path="stats" element={<pages.member.Stats />} />
        <Route path="leaderboard" element={<pages.member.Leaderboard />} />
        <Route path="1v1" element={<pages.member.OneVersusOne />} />
        <Route path="find" element={<pages.member.Find />} />
        <Route path="clans">
          <Route index element={<pages.member.clans.Index />} />
          <Route path=":clanTag">
            <Route index element={<pages.member.clans.View />} />
            <Route path="members">
              <Route path=":memberTag" element={<pages.member.clans.members.View />} />
            </Route>
          </Route>
        </Route>
      </Route>

      <Route path="leader" element={<LeaderRoute />}>
        <Route index element={<pages.leader.Index />} />
        <Route path="clans">
          <Route path=":clanTag">
            <Route path="settings" element={<pages.leader.clans.ClanSettings />} />
            <Route path="members">
              <Route index element={<pages.leader.clans.members.Index />} />
              <Route path=":memberTag">
                <Route path="kickpoints" element={<pages.leader.clans.members.Kickpoints />} />
              </Route>
            </Route>
          </Route>
        </Route>
        <Route path="notify" element={<pages.leader.Notify />} />
      </Route>

      <Route path="*" element={<Navigate to="/error/404" />} />
    </Routes>
  )
}
