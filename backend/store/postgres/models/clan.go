package models

type Clan struct {
	Tag  string `gorm:"primaryKey;not null"`
	Name string

	Settings    *ClanSettings `gorm:"foreignKey:ClanTag;references:Tag"`
	ClanMembers ClanMembers   `gorm:"foreignKey:ClanTag;references:Tag"`
}

type Clans []Clan

func (clans Clans) Tags() []string {
	tags := make([]string, len(clans))
	for i, clan := range clans {
		tags[i] = clan.Tag
	}

	return tags
}
