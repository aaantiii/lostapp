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
	"backend/store/postgres/models"
)

type IAuthService interface {
	Config() *AuthServiceConfig
	Session(token string) (*types.Session, bool)
	CreateSession(token string) (*types.Session, error)
	RefreshSession(token string) (*types.Session, error)
	DeleteSession(token string) error
	AuthCodeURL(state string) string
	ExchangeCode(code string) (*oauth2.Token, error)
	Guild(clanTag string) (*models.Guild, error)
}

type AuthService struct {
	guildsRepo repos.IGuildsRepo
	usersRepo  repos.IUsersRepo

	config             AuthServiceConfig
	discordOAuthConfig oauth2.Config
	sessions           types.Sessions
	httpClient         *http.Client
	discordClient      *client.DiscordClient
}

type AuthServiceConfig struct {
	MaxSessionIdleTime time.Duration // MaxSessionIdleTime defines how long a session will stay in cache without being used.
	MaxSessionDataAge  time.Duration // MaxSessionDataAge is the time after which a session is considered outdated. The data will be refetched when the session is used.
}

func NewAuthService(guildsRepo repos.IGuildsRepo, usersRepo repos.IUsersRepo) *AuthService {
	service := &AuthService{
		guildsRepo: guildsRepo,
		usersRepo:  usersRepo,
		config: AuthServiceConfig{
			MaxSessionIdleTime: 4 * time.Hour,
			MaxSessionDataAge:  15 * time.Minute,
		},
		discordOAuthConfig: oauth2.Config{
			ClientID:     env.DISCORD_CLIENT_ID.Value(),
			ClientSecret: env.DISCORD_CLIENT_SECRET.Value(),
			RedirectURL:  env.DISCORD_OAUTH_REDIRECT_URL.Value(),
			Scopes:       []string{"guilds.members.read"},
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

	discordUser, err := service.discordClient.FetchDiscordUser(token)
	if err != nil {
		return nil, err
	}

	user := &models.User{DiscordID: discordUser.ID, AvatarURL: discordUser.AvatarURL, Name: discordUser.Name}
	if err = service.usersRepo.CreateOrUpdateUser(user); err != nil {
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

	discordUser, err := service.discordClient.FetchDiscordUser(token)
	if err != nil {
		return nil, err
	}

	session.Refresh(discordUser)
	if err = service.grantRoleToSession(session); err != nil {
		return nil, err
	}

	return session, nil
}

// DeleteSession revokes a user's access token at discord and deletes the session from cache.
// An error is returned if the request fails or the response status is not http.StatusOK.
func (service *AuthService) DeleteSession(token string) error {
	bodyBuf := bytes.NewBuffer([]byte(fmt.Sprintf(
		"client_id=%s&client_secret=%s&token=%s",
		env.DISCORD_CLIENT_ID.Value(),
		env.DISCORD_CLIENT_SECRET.Value(),
		token,
	)))

	req, err := http.NewRequest("POST", service.discordOAuthConfig.Endpoint.TokenURL+"/revoke", bodyBuf)
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

	if isAdmin, err := service.usersRepo.UserIsAdmin(session.DiscordUser.ID); isAdmin && err == nil {
		session.AuthRole = types.AuthRoleAdmin
		return nil
	}

	guilds, err := service.guildsRepo.Guilds()
	if err != nil {
		return err
	}

	for _, guild := range guilds {
		if guild.IsLeader(session.DiscordUser.Roles) || guild.IsCoLeader(session.DiscordUser.Roles) {
			session.AuthRole = types.AuthRoleLeader
			return nil
		}

		if guild.IsElder(session.DiscordUser.Roles) || guild.IsMember(session.DiscordUser.Roles) {
			session.AuthRole = types.AuthRoleMember
		}
	}

	return nil
}

func (service *AuthService) Guild(clanTag string) (*models.Guild, error) {
	return service.guildsRepo.Guild(clanTag)
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
