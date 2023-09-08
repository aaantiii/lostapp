package repos

import (
	"fmt"

	"gorm.io/gorm"

	"backend/api/types"
	"backend/store/cache"
)

type IClansRepo interface {
	IsMaintenance() bool
	Clans() []*types.Clan
	ClanByTag(clanTag string) (*types.Clan, error)
}

type ClansRepo struct {
	db    *gorm.DB
	cache *cache.CocCache
}

func NewClansRepo(db *gorm.DB, cache *cache.CocCache) *ClansRepo {
	return &ClansRepo{db: db, cache: cache}
}

func (repo *ClansRepo) IsMaintenance() bool {
	return repo.cache.IsMaintenance()
}

func (repo *ClansRepo) Clans() []*types.Clan {
	return repo.cache.Clans
}

func (repo *ClansRepo) ClanByTag(tag string) (*types.Clan, error) {
	clan, found := repo.cache.ClanByTag[tag]
	if !found {
		return nil, fmt.Errorf("no clan found with tag=%s", tag)
	}

	return clan, nil
}
