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
	RevokeSession(token string) error
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

func (s *AuthService) Config() AuthServiceConfig {
	return s.config
}

// Session returns a types.Session associated with the provided token. An error is returned if the session does not exist.
func (s *AuthService) Session(token string) (*types.Session, bool) {
	session, ok := s.sessions.Get(token)
	if !ok {
		return nil, false
	}

	session.LastUsed = time.Now()
	s.sessions.Set(token, session)
	return &session, ok
}

// CreateSession initializes a new types.Session with the provided token and gets the user and guilds for that token from the discord api.
// An error is returned if the provided token already exists or if it's not a valid token.
func (s *AuthService) CreateSession(token string) (*types.Session, error) {
	if _, found := s.Session(token); found {
		return nil, errors.New("session with provided token already exists")
	}

	discordUser, err := s.discordClient.FetchDiscordUser(token)
	if err != nil {
		return nil, err
	}

	if err = s.users.CreateOrUpdateUser(&models.User{
		DiscordID: discordUser.ID,
		AvatarURL: discordUser.AvatarURL,
		Name:      discordUser.Name,
	}); err != nil {
		return nil, err
	}

	authUser, err := s.newAuthUser(discordUser)
	if err != nil {
		return nil, err
	}
	session := types.NewSession(authUser, token)

	s.sessions.Set(token, session)
	return &session, nil
}

// RefreshSession refetches the user from discord and updates the session in cache.
// An error is returned if the provided token does not exist or if the request to discord fails.
func (s *AuthService) RefreshSession(token string) (*types.Session, error) {
	if _, found := s.Session(token); !found {
		return nil, errors.New("tried to refresh non-existing session")
	}

	discordUser, err := s.discordClient.FetchDiscordUser(token)
	if err != nil {
		return nil, err
	}

	if err = s.users.CreateOrUpdateUser(&models.User{
		DiscordID: discordUser.ID,
		AvatarURL: discordUser.AvatarURL,
		Name:      discordUser.Name,
	}); err != nil {
		return nil, err
	}

	authUser, err := s.newAuthUser(discordUser)
	if err != nil {
		return nil, err
	}

	session := types.NewSession(authUser, token)
	s.sessions.Set(token, session)
	return &session, nil
}

// RevokeSession revokes a user's access token at discord and deletes the session from cache.
// An error is returned if the request fails or the response status is not http.StatusOK.
func (s *AuthService) RevokeSession(token string) error {
	s.sessions.Remove(token)
	return s.discordClient.RevokeToken(token)
}

func (s *AuthService) AuthCodeURL(state string) string {
	return s.discordClient.AuthCodeURL(state)
}

func (s *AuthService) ExchangeCode(code string) (*oauth2.Token, error) {
	return s.discordClient.ExchangeCode(code)
}

func (s *AuthService) newAuthUser(discordUser *types.DiscordUser) (*types.AuthUser, error) {
	guilds, err := s.guilds.GuildsByGuildID(env.DISCORD_GUILD_ID.Value())
	if err != nil {
		return nil, err
	}

	isAdmin, err := s.users.UserIsAdmin(discordUser.ID)
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
			user.CoLeaderOf = append(user.CoLeaderOf, guild.ClanTag)
			user.MemberOf = append(user.MemberOf, guild.ClanTag)
			continue
		}
		if guild.IsCoLeader(discordUser.Roles) {
			user.CoLeaderOf = append(user.CoLeaderOf, guild.ClanTag)
			user.MemberOf = append(user.MemberOf, guild.ClanTag)
			continue
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

// handleUnusedSessions removes all sessions from cache that have not been used for longer than AuthServiceConfig.MaxSessionIdleTime.
func (s *AuthService) handleUnusedSessions() {
	for range time.Tick(time.Hour) {
		for token, session := range s.sessions.Items() {
			if time.Since(session.LastUsed) > s.config.MaxSessionIdleTime {
				s.sessions.Remove(token)
			}
		}
	}
}
