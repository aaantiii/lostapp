package repos

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/aaantiii/lostapp/backend/store/postgres/models"
)

type IMembersRepo interface {
	MembersByClanTag(clanTag string) (models.ClanMembers, error)
	MembersByDiscordID(discordID string) (models.ClanMembers, error)
	MemberByID(playerTag, clanTag string) (*models.ClanMember, error)
	MembersByTag(clanTag string, playerTags ...string) (models.ClanMembers, error)
	MembersByPlayerTag(playerTag string) (models.ClanMembers, error)
	AllMemberTagsDistinct() ([]string, error)
	CreateMember(member *models.ClanMember) error
	UpdateMemberRole(playerTag, clanTag string, role models.ClanRole) error
	DeleteMember(tag, clanTag string) error
}

type MembersRepo struct {
	db *gorm.DB
}

func NewMembersRepo(db *gorm.DB) IMembersRepo {
	return &MembersRepo{db: db}
}

func (r *MembersRepo) MembersByClanTag(clanTag string) (models.ClanMembers, error) {
	var members models.ClanMembers
	if err := r.db.
		Preload("Player").
		Find(&members, "clan_tag = ?", clanTag).Error; err != nil {
		return nil, err
	}

	return members, nil
}

func (r *MembersRepo) MembersByDiscordID(discordID string) (models.ClanMembers, error) {
	var players []*models.Player
	err := r.db.
		Preload("ClanMembers").
		Find(&players, "discord_id = ?", discordID).Error

	var members models.ClanMembers
	for _, player := range players {
		members = append(members, player.Members...)
	}

	return members, err
}

func (r *MembersRepo) MemberByID(playerTag, clanTag string) (*models.ClanMember, error) {
	var members *models.ClanMember
	err := r.db.
		Preload(clause.Associations).
		First(&members, "player_tag = ? AND clan_tag = ?", playerTag, clanTag).Error
	return members, err
}

func (r *MembersRepo) MembersByTag(clanTag string, playerTags ...string) (models.ClanMembers, error) {
	var members models.ClanMembers
	err := r.db.
		Preload(clause.Associations).
		Find(&members, "clan_tag = ? AND player_tag IN (?)", clanTag, playerTags).Error
	return members, err
}

func (r *MembersRepo) MembersByPlayerTag(playerTag string) (models.ClanMembers, error) {
	var members models.ClanMembers
	err := r.db.
		Preload(clause.Associations).
		Find(&members, "player_tag = ?", playerTag).Error
	return members, err
}

func (r *MembersRepo) AllMemberTagsDistinct() ([]string, error) {
	var tags []string
	err := r.db.
		Model(&models.ClanMember{}).
		Distinct("player_tag").
		Pluck("player_tag", &tags).Error
	return tags, err
}

func (r *MembersRepo) MissingClanMembers(clanTag string, playerTags ...string) (models.ClanMembers, error) {
	var members models.ClanMembers
	err := r.db.
		Preload(clause.Associations).
		Find(&members, "clan_tag = ? AND player_tag NOT IN (?)", clanTag, playerTags).Error
	return members, err
}

func (r *MembersRepo) CreateMember(member *models.ClanMember) error {
	return r.db.Create(member).Error
}

func (r *MembersRepo) UpdateMemberRole(playerTag, clanTag string, role models.ClanRole) error {
	return r.db.
		Model(&models.ClanMember{PlayerTag: playerTag, ClanTag: clanTag}).
		Update("clan_role", role).Error
}

func (r *MembersRepo) DeleteMember(tag, clanTag string) error {
	return r.db.Delete(&models.ClanMember{}, "player_tag = ? AND clan_tag = ?", tag, clanTag).Error
}
