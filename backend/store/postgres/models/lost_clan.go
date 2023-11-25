package models

import "gorm.io/gorm"

type LostClan struct {
	ID   uint   `gorm:"primaryKey;autoIncrement;not null"` // used for ordering
	Tag  string `gorm:"unique;not null;size:12"`
	Name string `gorm:"not null;size:32"`

	Kickpoints *[]Kickpoint      `json:"kickpoints" gorm:"foreignKey:ClanTag;references:Tag;onUpdate:CASCADE;onDelete:CASCADE"`
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

func SeedLostClans(db *gorm.DB) error {
	if db.Migrator().HasTable(&LostClan{}) {
		if err := db.First(&LostClan{}).Error; err == nil {
			return nil
		}
	}

	return db.Create(lostClansSeedData()).Error
}

func lostClansSeedData() []*LostClan {
	return []*LostClan{{
		Tag:  "#2820UPPQC",
		Name: "LOST F2P",
	}, {
		Tag:  "#2LG222Q0L",
		Name: "LOST F2P 2",
	}, {
		Tag:  "#2YUPV0UYC",
		Name: "LOST 3",
	}, {
		Tag:  "#2LU2V2LPU",
		Name: "LOST 4",
	}, {
		Tag:  "#2QC0QQPQ2",
		Name: "LOST 5",
	}, {
		Tag:  "#2YVPC20UY",
		Name: "LOST 6",
	}, {
		Tag:  "#2QQ29JCYV",
		Name: "LOST 7",
	}, {
		Tag:  "#2YVJV8VC0",
		Name: "LOST GP",
	}, {
		Tag:  "#202CVQ0GQ",
		Name: "Anthrazit",
	}}
}
