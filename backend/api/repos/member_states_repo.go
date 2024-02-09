package repos

import (
	"gorm.io/gorm"

	"github.com/aaantiii/lostapp/backend/store/postgres/models"
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

func (r *MemberStatesRepo) IsKickpointLocked(playerTag, clanTag string) (bool, error) {
	var state models.MemberState
	err := r.db.
		Select("kickpoint_lock").
		First(&state, "player_tag = ? AND clan_tag = ?", playerTag, clanTag).Error
	return state.KickpointLock, err
}

func (r *MemberStatesRepo) UpdateKickpointLockStatus(playerTag, clanTag string, signOff bool) error {
	return r.db.Save(&models.MemberState{
		PlayerTag:     playerTag,
		ClanTag:       clanTag,
		KickpointLock: signOff,
	}).Error
}
