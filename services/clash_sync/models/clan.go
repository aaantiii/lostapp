package models

type Clan struct {
	Tag   string `gorm:"primaryKey;not null"`
	Name  string
	Index int
}

type Clans []Clan

func (clans Clans) Tags() []string {
	tags := make([]string, len(clans))
	for i, clan := range clans {
		tags[i] = clan.Tag
	}
	return tags
}
