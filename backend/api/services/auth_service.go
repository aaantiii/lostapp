package services

import (
	"errors"
	"net/http"
	"time"

	cmap "github.com/orcaman/concurrent-map/v2"
	"golang.org/x/oauth2"

	"github.com/aaantiii/lostapp/backend/api/repos"
	"github.com/aaantiii/lostapp/backend/api/types"
	"github.com/aaantiii/lostapp/backend/client"
	"github.com/aaantiii/lostapp/backend/env"
	"github.com/aaantiii/lostapp/backend/store/postgres/models"
)

type IAuthService interface {
	Config() AuthServiceConfig
	Session(token string) (*types.Session, bool)
	CreateSession(token string) (*types.Session, error)
	RefreshSession(token string) (*types.Session, error)
	AuthCodeURL(state string) string
	ExchangeCode(code string) (*oauth2.Token, error)
}

type AuthService struct {
	guilds repos.IGuildsRepo
	users  repos.IUsersRepo

	config        AuthServiceConfig
	sessions      cmap.ConcurrentMap[string, types.Session]
	httpClient    *http.Client
	discordClient *client.DiscordClient
}

type AuthServiceConfig struct {
	MaxSessionIdleTime time.Duration // MaxSessionIdleTime defines how long a session will stay in cache without being used.
	MaxSessionDataAge  time.Duration // MaxSessionDataAge is the time after which a session is considered outdated. The data will be refetched when the session is used.
}

func NewAuthService(guildsRepo repos.IGuildsRepo, usersRepo repos.IUsersRepo) IAuthService {
	service := &AuthService{
		guilds: guildsRepo,
		users:  usersRepo,
		config: AuthServiceConfig{
			MaxSessionIdleTime: 4 * time.Hour,
			MaxSessionDataAge:  15 * time.Minute,
		},
		httpClient: client.NewHttpClient(),
		discordClient: client.NewDiscordClient(&oauth2.Config{
			ClientID:     env.DISCORD_CLIENT_ID.Value(),
			ClientSecret: env.DISCORD_CLIENT_SECRET.Value(),
			RedirectURL:  env.DISCORD_OAUTH_REDIRECT_URL.Value(),
			Scopes:       []string{"guilds.members.read"},
			Endpoint: oauth2.Endpoint{
				AuthURL:   client.DiscordAuthRoute.URL(),
				TokenURL:  client.DiscordTokenRoute.URL(),
				AuthStyle: oauth2.AuthStyleInParams,
			},
		}),
		sessions: cmap.New[types.Session](),
	}
	go service.handleUnusedSessions()

	return service
}

func (service *AuthService) Config() AuthServiceConfig {
	return service.config
}

// Session returns a types.Session associated with the provided token. An error is returned if the session does not exist.
func (service *AuthService) Session(token string) (*types.Session, bool) {
	session, ok := service.sessions.Get(token)
	if !ok {
		return nil, false
	}

	session.LastUsed = time.Now()
	service.sessions.Set(token, session)
	return &session, ok
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

	if err = service.users.CreateOrUpdateUser(&models.User{
		DiscordID: discordUser.ID,
		AvatarURL: discordUser.AvatarURL,
		Name:      discordUser.Name,
	}); err != nil {
		return nil, err
	}

	authUser, err := service.newAuthUser(discordUser)
	if err != nil {
		return nil, err
	}
	session := types.NewSession(authUser, token)

	service.sessions.Set(token, session)
	return &session, nil
}

// RefreshSession refetches the user from discord and updates the session in cache.
// An error is returned if the provided token does not exist or if the request to discord fails.
func (service *AuthService) RefreshSession(token string) (*types.Session, error) {
	if _, found := service.Session(token); !found {
		return nil, errors.New("tried to refresh non-existing session")
	}

	discordUser, err := service.discordClient.FetchDiscordUser(token)
	if err != nil {
		return nil, err
	}

	if err = service.users.CreateOrUpdateUser(&models.User{
		DiscordID: discordUser.ID,
		AvatarURL: discordUser.AvatarURL,
		Name:      discordUser.Name,
	}); err != nil {
		return nil, err
	}

	authUser, err := service.newAuthUser(discordUser)
	if err != nil {
		return nil, err
	}

	session := types.NewSession(authUser, token)
	service.sessions.Set(token, session)
	return &session, nil
}

// DeleteSession revokes a user's access token at discord and deletes the session from cache.
// An error is returned if the request fails or the response status is not http.StatusOK.
func (service *AuthService) DeleteSession(token string) error {

	service.sessions.Remove(token)
	return nil
}

func (service *AuthService) AuthCodeURL(state string) string {
	return service.discordClient.AuthCodeURL(state)
}

func (service *AuthService) ExchangeCode(code string) (*oauth2.Token, error) {
	return service.discordClient.ExchangeCode(code)
}

func (service *AuthService) newAuthUser(discordUser *types.DiscordUser) (*types.AuthUser, error) {
	guilds, err := service.guilds.GuildsByGuildID(env.DISCORD_GUILD_ID.Value())
	if err != nil {
		return nil, err
	}

	isAdmin, err := service.users.UserIsAdmin(discordUser.ID)
	if err != nil {
		return nil, err
	}

	user := &types.AuthUser{
		ID:        discordUser.ID,
		Name:      discordUser.Name,
		AvatarURL: discordUser.AvatarURL,
		IsAdmin:   isAdmin,
	}

	for _, guild := range guilds {
		if guild.IsLeader(discordUser.Roles) {
			user.LeaderOf = append(user.LeaderOf, guild.ClanTag)
		}
		if guild.IsCoLeader(discordUser.Roles) {
			user.CoLeaderOf = append(user.CoLeaderOf, guild.ClanTag)
		}
		if guild.IsMember(discordUser.Roles) {
			user.MemberOf = append(user.MemberOf, guild.ClanTag)
		}
	}

	if len(user.MemberOf) == 0 {
		return nil, types.ErrNotMember
	}

	return user, nil
}

// handleUnusedSessions deletes all sessions from cache that have not been used for longer than services.MaxSessionIdleTime.
func (service *AuthService) handleUnusedSessions() {
	for range time.Tick(time.Hour) {
		for token, session := range service.sessions.Items() {
			if time.Since(session.LastUsed) > service.config.MaxSessionIdleTime {
				service.sessions.Remove(token)
			}
		}
	}
}
