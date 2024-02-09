package models

type MemberState struct {
	PlayerTag     string `gorm:"primaryKey" json:"playerTag"`
	ClanTag       string `gorm:"primaryKey" json:"clanTag"`
	KickpointLock bool   `gorm:"default:false;not null" json:"kickpointLock"`

	Member *ClanMember `gorm:"foreignKey:PlayerTag,ClanTag" json:"member,omitempty"`
}
