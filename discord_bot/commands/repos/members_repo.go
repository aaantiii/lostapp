package repos

import (
	"gorm.io/gorm"

	"bot/store/postgres/models"
)

type IMembersRepo interface {
	MembersByDiscordID(discordID string) (models.Members, error)
	MemberByID(playerTag, clanTag string) (*models.Member, error)
	MembersByTag(clanTag string, playerTags ...string) (models.Members, error)
	CreateMember(member *models.Member) error
	UpdateMemberRole(playerTag, clanTag string, role models.ClanRole) error
	DeleteMember(tag, clanTag string) error
}

type MembersRepo struct {
	db *gorm.DB
}

func NewMembersRepo(db *gorm.DB) IMembersRepo {
	return &MembersRepo{db: db}
}

func (repo *MembersRepo) MembersByDiscordID(discordID string) (models.Members, error) {
	var players []*models.Player
	err := repo.db.
		Preload("Members").
		Find(&players, "discord_id = ?", discordID).Error

	var members models.Members
	for _, player := range players {
		members = append(members, player.Members...)
	}

	return members, err
}

func (repo *MembersRepo) MemberByID(playerTag, clanTag string) (*models.Member, error) {
	var members *models.Member
	err := repo.db.
		Preload("Player").
		Preload("Clan").
		First(&members, "player_tag = ? AND clan_tag = ?", playerTag, clanTag).Error
	return members, err
}

func (repo *MembersRepo) MembersByTag(clanTag string, playerTags ...string) (models.Members, error) {
	var members models.Members
	err := repo.db.
		Preload("Player").
		Preload("Clan").
		Find(&members, "clan_tag = ? AND player_tag IN (?)", clanTag, playerTags).Error
	return members, err
}

func (repo *MembersRepo) MissingClanMembers(clanTag string, playerTags ...string) (models.Members, error) {
	var members models.Members
	err := repo.db.
		Preload("Player").
		Preload("Clan").
		Find(&members, "clan_tag = ? AND player_tag NOT IN (?)", clanTag, playerTags).Error
	return members, err
}

func (repo *MembersRepo) CreateMember(member *models.Member) error {
	return repo.db.Create(member).Error
}

func (repo *MembersRepo) UpdateMemberRole(playerTag, clanTag string, role models.ClanRole) error {
	return repo.db.
		Model(&models.Member{PlayerTag: playerTag, ClanTag: clanTag}).
		Update("clan_role", role).Error
}

func (repo *MembersRepo) DeleteMember(tag, clanTag string) error {
	return repo.db.Delete(&models.Member{}, "player_tag = ? AND clan_tag = ?", tag, clanTag).Error
}
