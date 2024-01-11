package repos

import (
	"gorm.io/gorm"

	"bot/store/postgres/models"
)

type IMemberStatesRepo interface {
	IsKickpointLocked(playerTag, clanTag string) (bool, error)
	UpdateKickpointLockStatus(playerTag, clanTag string, signOff bool) error
}

type MemberStatesRepo struct {
	db *gorm.DB
}

func NewMemberStatesRepo(db *gorm.DB) IMemberStatesRepo {
	return &MemberStatesRepo{db: db}
}

func (repo *MemberStatesRepo) IsKickpointLocked(playerTag, clanTag string) (bool, error) {
	var state models.MemberState
	err := repo.db.
		Select("kickpoint_lock").
		First(&state, "player_tag = ? AND clan_tag = ?", playerTag, clanTag).Error
	return state.KickpointLock, err
}

func (repo *MemberStatesRepo) UpdateKickpointLockStatus(playerTag, clanTag string, signOff bool) error {
	return repo.db.Save(&models.MemberState{
		PlayerTag:     playerTag,
		ClanTag:       clanTag,
		KickpointLock: signOff,
	}).Error
}
