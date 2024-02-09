package types

import "time"

type Session struct {
	User  *AuthUser `json:"user"`
	Token string    `json:"-"`

	LastRefreshed time.Time `json:"-"`
	LastUsed      time.Time `json:"-"`
}

func (s *Session) NeedsRefresh(maxAge time.Duration) bool {
	return time.Since(s.LastRefreshed) > maxAge
}

type AuthUser struct {
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	AvatarURL  string   `json:"avatarUrl"`
	MemberOf   []string `json:"memberOf,omitempty"`
	CoLeaderOf []string `json:"coLeaderOf,omitempty"`
	LeaderOf   []string `json:"leaderOf,omitempty"`
	IsAdmin    bool     `json:"isAdmin"`
}

type AuthRole string

const (
	AuthRoleMember   AuthRole = "member"
	AuthRoleCoLeader AuthRole = "coLeader"
	AuthRoleLeader   AuthRole = "leader"
	AuthRoleAdmin    AuthRole = "~~~admin~~~"
)

func NewSession(user *AuthUser, token string) Session {
	currentTime := time.Now()
	return Session{
		User:          user,
		Token:         token,
		LastRefreshed: currentTime,
		LastUsed:      currentTime,
	}
}
