package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/aaantiii/lostapp/backend/api/types"
	"github.com/aaantiii/lostapp/backend/env"
)

type DiscordClient struct {
	httpClient *http.Client
	config     *oauth2.Config
}

// discordGuildMemberResponse is the response body of the Discord API when requesting a user.
type discordGuildMemberResponse struct {
	User struct {
		ID     string `json:"id"`
		Avatar string `json:"avatar"`
	} `json:"user"`
	Nick  string   `json:"nick"`
	Roles []string `json:"roles"`
}

func (res *discordGuildMemberResponse) discordUser() *types.DiscordUser {
	if res == nil {
		return nil
	}
	return &types.DiscordUser{
		ID:        res.User.ID,
		Name:      res.Nick,
		AvatarURL: fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s", res.User.ID, res.User.Avatar),
		Roles:     res.Roles,
	}
}

type DiscordApiRoute string

const (
	DiscordBaseURL                          = "https://discord.com/api"
	DiscordGuildMemberRoute DiscordApiRoute = "/users/@me/guilds/733857906117574717/member"
	DiscordAuthRoute        DiscordApiRoute = "/oauth2/authorize"
	DiscordTokenRoute       DiscordApiRoute = "/oauth2/token"
)

func (route DiscordApiRoute) URL() string {
	return DiscordBaseURL + string(route)
}

func NewDiscordClient(config *oauth2.Config) *DiscordClient {
	return &DiscordClient{
		httpClient: NewHttpClient(),
		config:     config,
	}
}

func (client *DiscordClient) FetchDiscordUser(token string) (*types.DiscordUser, error) {
	var res *discordGuildMemberResponse
	if err := client.fetch(token, DiscordGuildMemberRoute, &res); err != nil {
		return nil, err
	}

	return res.discordUser(), nil
}

// fetch fetches data from the Discord-API and stores the response body in the variable pointed to by dest.
func (client *DiscordClient) fetch(token string, route DiscordApiRoute, dest interface{}) error {
	req, err := http.NewRequest(http.MethodGet, route.URL(), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Bearer "+token)

	res, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return errors.New("request to discord api failed")
	}
	return json.NewDecoder(res.Body).Decode(dest)
}

func (client *DiscordClient) RevokeToken(token string) error {
	bodyBuf := bytes.NewBuffer([]byte(fmt.Sprintf(
		"client_id=%s&client_secret=%s&token=%s",
		env.DISCORD_CLIENT_ID.Value(),
		env.DISCORD_CLIENT_SECRET.Value(),
		token,
	)))

	req, err := http.NewRequest(http.MethodPost, client.config.Endpoint.TokenURL+"/revoke", bodyBuf)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	res, err := client.httpClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("discord responded with status %s after trying to revoke a token", res.Status)
	}

	return nil
}

func (client *DiscordClient) AuthCodeURL(state string) string {
	return client.config.AuthCodeURL(state, oauth2.AccessTypeOnline)
}

func (client *DiscordClient) ExchangeCode(code string) (*oauth2.Token, error) {
	return client.config.Exchange(context.Background(), code)
}
