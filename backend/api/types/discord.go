package types

// DiscordUser is used to store the user data from the Discord API and rename the JSON struct tags.
type DiscordUser struct {
	ID        string   `json:"id"`
	Name      string   `json:"name"`
	AvatarURL string   `json:"avatarUrl"`
	Roles     []string `json:"-"`
}
