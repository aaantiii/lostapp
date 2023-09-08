package types

import "time"

type AuthRole uint8

const (
	AuthRoleUser AuthRole = 1 + iota
	AuthRoleMember
	AuthRoleLeader
	AuthRoleAdmin
)

type Session struct {
	DiscordUser *DiscordUser `json:"discordUser"`
	AuthRole    AuthRole     `json:"userRole"`

	AccessToken string `json:"-"`

	lastRefreshed time.Time
	lastUsed      time.Time
}

// Sessions maps a session token to a Session.
type Sessions map[string]*Session

func NewSession(user *DiscordUser, token string) *Session {
	currentTime := time.Now()
	return &Session{
		DiscordUser:   user,
		AccessToken:   token,
		lastRefreshed: currentTime,
		lastUsed:      currentTime,
	}
}

func (session *Session) Refresh(discordUser *DiscordUser) {
	session.DiscordUser = discordUser
	session.lastRefreshed = time.Now()
}

// LastRefreshed returns the time.Duration since the session data was last refreshed.
func (session *Session) LastRefreshed() time.Duration {
	return time.Since(session.lastRefreshed)
}

// LastUsed returns the time.Duration since the session was last used.
func (session *Session) LastUsed() time.Duration {
	return time.Since(session.lastUsed)
}

// HasPermission returns true if the Session.AuthRole is greater than or equal to the required role.
func (session *Session) HasPermission(requiredRole AuthRole) bool {
	session.lastUsed = time.Now() // this function is called every time the session is used
	return session.AuthRole >= requiredRole
}
