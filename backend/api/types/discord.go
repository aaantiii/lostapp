package types

import "fmt"

// DiscordUserResponse is the response body of the Discord API when requesting a user.
type DiscordUserResponse struct {
	ID         string `json:"id" binding:"required,len=18,alphanum"`
	GlobalName string `json:"global_name" binding:"required"`
	Avatar     string `json:"avatar" binding:"required"`
}

// DiscordUser is used to store the user data from the Discord API and rename the JSON struct tags.
type DiscordUser struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatarUrl"`
}

func (res *DiscordUserResponse) DiscordUser() *DiscordUser {
	return &DiscordUser{
		ID:        res.ID,
		Username:  res.GlobalName,
		AvatarURL: fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s", res.ID, res.Avatar),
	}
}
