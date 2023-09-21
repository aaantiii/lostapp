package types

import (
	"fmt"
)

// DiscordGuildMember is the response body of the Discord API when requesting a user.
type DiscordGuildMember struct {
	User struct {
		ID     string `json:"id" binding:"required,len=18,alphanum"`
		Avatar string `json:"avatar" binding:"required"`
	} `json:"user"`
	Nick  string   `json:"nick" binding:"required"`
	Roles []string `json:"roles" binding:"required"`
}

// DiscordUser is used to store the user data from the Discord API and rename the JSON struct tags.
type DiscordUser struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	AvatarURL string   `json:"avatarUrl"`
	Roles     []string `json:"-"`
}

func (res *DiscordGuildMember) DiscordUser() *DiscordUser {
	if res == nil {
		return nil
	}

	return &DiscordUser{
		ID:        res.User.ID,
		Name:      res.Nick,
		AvatarURL: fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s", res.User.ID, res.User.Avatar),
		Roles:     res.Roles,
	}
}
