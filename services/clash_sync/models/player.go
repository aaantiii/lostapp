package models

// Player links a Clash of Clans player tag with a Discord ID.
type Player struct {
	CocTag    string `gorm:"not null;primaryKey"`
	Name      string `gorm:"not null"`
	DiscordID string
}

type Players []*Player

func (players Players) Tags() []string {
	tags := make([]string, len(players))
	for i, player := range players {
		tags[i] = player.CocTag
	}

	return tags
}
