package repos

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/aaantiii/lostapp/backend/api/types"
	"github.com/aaantiii/lostapp/backend/api/utils"
	"github.com/aaantiii/lostapp/backend/store/postgres"
	"github.com/aaantiii/lostapp/backend/store/postgres/models"
)

type IMembersRepo interface {
	MemberByID(playerTag, clanTag string) (*models.ClanMember, error)
	Members(params types.MembersParams, preload ...postgres.Preloader) (models.ClanMembers, error)
	MembersPaginated(params types.MembersParams, preload ...postgres.Preloader) (*types.PaginatedResponse[*models.ClanMember], error)
	Count(params types.MembersParams) (int64, error)
	MemberTagsDistinct(params types.MembersParams) ([]string, error)
	CreateMember(member *models.ClanMember) error
	UpdateMemberRole(playerTag, clanTag string, role models.ClanRole) error
	DeleteMember(playerTag, clanTag string) error
}

type MembersRepo struct {
	db *gorm.DB
}

func NewMembersRepo(db *gorm.DB) IMembersRepo {
	return &MembersRepo{db: db}
}

func (r *MembersRepo) MemberByID(playerTag, clanTag string) (*models.ClanMember, error) {
	var members *models.ClanMember
	err := r.db.
		Preload(clause.Associations).
		First(&members, "player_tag = ? AND clan_tag = ?", playerTag, clanTag).Error
	return members, err
}

func (r *MembersRepo) Members(params types.MembersParams, preload ...postgres.Preloader) (models.ClanMembers, error) {
	var members models.ClanMembers
	err := r.db.
		Scopes(postgres.WithPreloading(preload...)).
		Where(params.Conds()).
		Find(&members).Error
	return members, err
}

func (r *MembersRepo) MembersPaginated(params types.MembersParams, preload ...postgres.Preloader) (*types.PaginatedResponse[*models.ClanMember], error) {
	count, err := r.Count(params)
	if err != nil {
		return nil, err
	}
	if err = utils.ValidatePagination(params.PaginationParams, count); err != nil {
		return nil, err
	}

	var members models.ClanMembers
	if err = r.db.
		Scopes(
			postgres.WithPagination(params.PaginationParams),
			postgres.WithPreloading(preload...),
		).
		Where(params.Conds()).
		Find(&members).Error; err != nil {
		return nil, err
	}
	return types.NewPaginatedResponse[*models.ClanMember](members, params.PaginationParams, count), nil
}

func (r *MembersRepo) Count(params types.MembersParams) (int64, error) {
	var count int64
	err := r.db.
		Model(&models.ClanMember{}).
		Where(params.Conds()).
		Count(&count).Error
	return count, err
}

func (r *MembersRepo) MembersByTags(clanTag string, playerTags ...string) (models.ClanMembers, error) {
	var members models.ClanMembers
	err := r.db.
		Preload(clause.Associations).
		Find(&members, "clan_tag = ? AND player_tag IN (?)", clanTag, playerTags).Error
	return members, err
}

func (r *MembersRepo) MemberTagsDistinct(params types.MembersParams) ([]string, error) {
	var tags []string
	err := r.db.
		Model(&models.ClanMember{}).
		Distinct("player_tag").
		Where(params.Conds()).
		Pluck("player_tag", &tags).Error
	return tags, err
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
