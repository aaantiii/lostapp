package models

// Player links a Clash of Clans player tag with a Discord ID.
type Player struct {
	CocTag    string `gorm:"not null;primaryKey" json:"cocTag"`
	Name      string `gorm:"not null" json:"name"`
	DiscordID string `json:"discordId"`

	Members ClanMembers `gorm:"foreignKey:PlayerTag;references:CocTag" json:"members,omitempty"`
}

type Players []*Player
