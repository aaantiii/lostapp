package models

type Clan struct {
	Tag  string `gorm:"primaryKey;not null" json:"tag"`
	Name string `json:"name"`

	Settings    *ClanSettings `gorm:"foreignKey:ClanTag;references:Tag" json:"settings,omitempty"`
	ClanMembers ClanMembers   `gorm:"foreignKey:ClanTag;references:Tag" json:"clanMembers,omitempty"`
}

type Clans []*Clan

func (clans Clans) Tags() []string {
	tags := make([]string, len(clans))
	for i, clan := range clans {
		tags[i] = clan.Tag
	}

	return tags
}
