package services

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"golang.org/x/oauth2"

	"backend/api/repos"
	"backend/api/types"
	"backend/client"
	"backend/env"
)

type IAuthService interface {
	Config() *AuthServiceConfig
	Session(token string) (*types.Session, bool)
	CreateSession(token string) (*types.Session, error)
	RefreshSession(token string) (*types.Session, error)
	DeleteSession(token string) error
	AuthCodeURL(state string) string
	ExchangeCode(code string) (*oauth2.Token, error)
}

type AuthService struct {
	membersRepo        repos.IMembersRepo
	config             AuthServiceConfig
	discordOAuthConfig *oauth2.Config
	sessions           types.Sessions
	httpClient         *http.Client
	discordClient      *client.DiscordClient
}

type AuthServiceConfig struct {
	MaxSessionIdleTime time.Duration // MaxSessionIdleTime defines how long a session will stay in cache without being used.
	MaxSessionDataAge  time.Duration // MaxSessionDataAge is the time after which a session is considered outdated. The data will be refetched when the session is used.
}

func NewAuthService(membersRepo repos.IMembersRepo) *AuthService {
	service := &AuthService{
		membersRepo: membersRepo,
		config: AuthServiceConfig{
			MaxSessionIdleTime: 12 * time.Hour,
			MaxSessionDataAge:  15 * time.Minute,
		},
		discordOAuthConfig: &oauth2.Config{
			ClientID:     env.DISCORD_CLIENT_ID.Value(),
			ClientSecret: env.DISCORD_CLIENT_SECRET.Value(),
			RedirectURL:  env.DISCORD_OAUTH_REDIRECT_URL.Value(),
			Scopes:       []string{"identify"},
			Endpoint: oauth2.Endpoint{
				AuthURL:   client.DiscordAuthRoute.URL(),
				TokenURL:  client.DiscordTokenRoute.URL(),
				AuthStyle: oauth2.AuthStyleInParams,
			},
		},
		httpClient:    client.NewHttpClient(),
		discordClient: client.NewDiscordClient(),
		sessions:      make(types.Sessions),
	}
	go service.handleUnusedSessions()

	return service
}

func (service *AuthService) Config() *AuthServiceConfig {
	return &service.config
}

// Session returns a types.Session associated with the provided token. An error is returned if the session does not exist.
func (service *AuthService) Session(token string) (*types.Session, bool) {
	session, found := service.sessions[token]
	return session, found
}

// CreateSession initializes a new types.Session with the provided token and gets the user and guilds for that token from the discord api.
// An error is returned if the provided token already exists or if it's not a valid token.
func (service *AuthService) CreateSession(token string) (*types.Session, error) {
	if _, found := service.Session(token); found {
		return nil, errors.New("session with provided token already exists")
	}

	discordUser, err := service.discordClient.FetchUser(token)
	if err != nil {
		return nil, err
	}

	session := types.NewSession(discordUser, token)
	if err = service.grantRoleToSession(session); err != nil {
		return nil, err
	}

	service.sessions[token] = session
	return session, nil
}

// RefreshSession refetches the user from discord and updates the session in cache.
// An error is returned if the provided token does not exist or if the request to discord fails.
func (service *AuthService) RefreshSession(token string) (*types.Session, error) {
	session, found := service.Session(token)
	if !found {
		return nil, errors.New("session not found")
	}

	discordUser, err := service.discordClient.FetchUser(token)
	if err != nil {
		return nil, err
	}

	session.Refresh(discordUser)
	if err = service.grantRoleToSession(session); err != nil {
		return nil, err
	}

	return session, nil
}

// DeleteSession revokes a user's access token at discordClient and deletes the session from cache.
// An error is returned if the request fails or the response status is not http.StatusOK.
func (service *AuthService) DeleteSession(token string) error {
	body := fmt.Sprintf(
		"client_id=%s&client_secret=%s&token=%s",
		env.DISCORD_CLIENT_ID.Value(),
		env.DISCORD_CLIENT_SECRET.Value(),
		token,
	)

	req, err := http.NewRequest("POST", service.discordOAuthConfig.Endpoint.TokenURL+"/revoke", bytes.NewBuffer([]byte(body)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := service.httpClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("discord responded with status %s after trying to revoke a token", res.Status)
	}

	delete(service.sessions, token)
	return nil
}

func (service *AuthService) AuthCodeURL(state string) string {
	return service.discordOAuthConfig.AuthCodeURL(state)
}

func (service *AuthService) ExchangeCode(code string) (*oauth2.Token, error) {
	return service.discordOAuthConfig.Exchange(context.Background(), code)
}

func (service *AuthService) grantRoleToSession(session *types.Session) error {
	session.AuthRole = types.AuthRoleUser

	members, err := service.membersRepo.MembersByDiscordID(session.DiscordUser.ID)
	if err != nil {
		return err
	}

	for _, member := range members {
		if member.IsAdmin {
			session.AuthRole = types.AuthRoleAdmin
			break
		}

		if session.AuthRole < types.AuthRoleMember && (member.ClanRole.IsElder() || member.ClanRole.IsMember()) {
			session.AuthRole = types.AuthRoleMember
		}

		if member.ClanRole.IsLeader() || member.ClanRole.IsCoLeader() {
			session.AuthRole = types.AuthRoleLeader
		}
	}

	return nil
}

// handleUnusedSessions deletes all sessions from cache that have not been used for longer than services.MaxSessionIdleTime.
func (service *AuthService) handleUnusedSessions() {
	for range time.Tick(time.Hour) {
		for token, session := range service.sessions {
			if session.LastUsed() > service.config.MaxSessionIdleTime {
				delete(service.sessions, token)
			}
		}
	}
}
