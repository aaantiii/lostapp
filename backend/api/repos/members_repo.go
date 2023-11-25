package repos

import (
	"gorm.io/gorm"

	"backend/store/cache"
	"backend/store/postgres/models"
)

type IMembersRepo interface {
	MembersByID(playerTag, clanTag string) (models.Members, error)
	MembersByDiscordID(discordID string) (models.Members, error)
	CreateMember(member *models.Member) error
	UpdateMember(member *models.Member) error
	DeleteMember(tag, clanTag string) error
}

type MembersRepo struct {
	db *gorm.DB
}

func NewMembersRepo(db *gorm.DB) IMembersRepo {
	return &MembersRepo{db: db}
}

func (repo *MembersRepo) MembersByID(playerTag, clanTag string) (models.Members, error) {
	var members models.Members
	err := repo.db.
		Preload("DiscordLink").
		Find(&members, "player_tag = ? AND clan_tag = ?", playerTag, clanTag).Error
	return members, err
}

func (repo *MembersRepo) MembersByDiscordID(discordID string) (models.Members, error) {
	var members models.Members
	err := repo.db.
		Raw("SELECT * FROM clan_member WHERE player_tag IN (SELECT coc_tag FROM player WHERE discord_id LIKE ?) ORDER BY ?;", discordID, cache.MemberOrder).
		Scan(&members).Error
	return members, err
}

func (repo *MembersRepo) CreateMember(member *models.Member) error {
	return repo.db.Create(member).Error
}

func (repo *MembersRepo) UpdateMember(member *models.Member) error {
	return repo.db.Save(member).Error
}

func (repo *MembersRepo) DeleteMember(tag, clanTag string) error {
	return repo.db.Delete(&models.Member{}, "player_tag = ? AND clan_tag = ?", tag, clanTag).Error
}
