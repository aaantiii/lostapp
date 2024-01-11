package models

type MemberState struct {
	PlayerTag     string `gorm:"primaryKey"`
	ClanTag       string `gorm:"primaryKey"`
	KickpointLock bool   `gorm:"default:false;not null"`

	Member *ClanMember `gorm:"foreignKey:PlayerTag,ClanTag"`
}
