package models

// Player links a Clash of Clans player tag with a Discord ID.
type Player struct {
	CocTag    string `gorm:"not null;primaryKey"`
	Name      string `gorm:"not null"`
	DiscordID string

	Members ClanMembers `gorm:"foreignKey:PlayerTag;references:CocTag"`
}

type Players []*Player
