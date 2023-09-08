package models

type LostClan struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;not null"` // used to order the clans
	Tag  string `gorm:"unique;not null;size:10"`
	Name string `gorm:"not null;size:32"`

	Kickpoints []*Kickpoint      `json:"kickpoints" gorm:"foreignKey:ClanTag;references:Tag;onUpdate:CASCADE;onDelete:CASCADE"`
	Settings   *LostClanSettings `json:"settings" gorm:"foreignKey:ClanTag;references:Tag;onUpdate:CASCADE;onDelete:CASCADE"`
}

type LostClans []*LostClan

func (clans LostClans) Tags() []string {
	tags := make([]string, len(clans))
	for i, clan := range clans {
		tags[i] = clan.Tag
	}

	return tags
}
