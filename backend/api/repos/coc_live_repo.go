package repos

import (
	"github.com/amaanq/coc.go"
	"gorm.io/gorm"

	"backend/client"
	"backend/store/postgres/models"
)

type ICocLiveRepo interface {
	Save(clans []*coc.Clan, players []*coc.Player) error
}

type CocLiveRepo struct {
	db        *gorm.DB
	cocClient *client.CocClient
}

func NewCocLiveRepo(db *gorm.DB) ICocLiveRepo {
	return &CocLiveRepo{db: db}
}

func (repo *CocLiveRepo)

func (repo *CocLiveRepo) Save(cocClans []*coc.Clan, players []*coc.Player) error {
	if err := repo.saveClans(cocClans); err != nil {
		return err
	}
	if err := repo.savePlayers(players); err != nil {
		return err
	}
	if err := repo.savePlayerClans(cocClans); err != nil {
		return err
	}
	if err := repo.savePlayerStats(players); err != nil {
		return err
	}

	return nil
}

func (repo *CocLiveRepo) saveClans(cocClans []*coc.Clan) error {
	var members models.Members
	if err := repo.db.Find(&members).Error; err != nil {
		return err
	}

	membersByClan := make(map[string]models.Members)
	for _, member := range members {
		membersByClan[member.ClanTag] = append(membersByClan[member.ClanTag], member)
	}

	clans := make([]*models.Clan, len(cocClans))
	for i, cocClan := range cocClans {
		cocClan.Members = len(membersByClan[cocClan.Tag])
		clans[i] = models.NewClan(cocClan)
	}

	return repo.db.Save(clans).Error
}

// SavePlayers converts the given coc.Players to models.Player and saves them to the database.
func (repo *CocLiveRepo) savePlayers(cocPlayers []*coc.Player) error {
	var members models.Members
	if err := repo.db.Find(&members).Error; err != nil {
		return err
	}

	membersByTag := make(map[string]models.Members)
	for _, member := range members {
		membersByTag[member.PlayerTag] = append(membersByTag[member.PlayerTag], member)
	}

	players := make([]*models.Player, len(cocPlayers))
	for i, player := range cocPlayers {
		member, found := membersByTag[player.Tag]
		if found {
			continue
		}
		players[i] = models.NewPlayer(player, member[0].DiscordLink.DiscordID)
	}

	return repo.db.Save(players).Error
}

func (repo *CocLiveRepo) savePlayerClans(cocClans []*coc.Clan) error {
	var members models.Members
	if err := repo.db.Find(&members).Error; err != nil {
		return err
	}

	clanByTag := make(map[string]*coc.Clan, len(cocClans))
	for _, clan := range cocClans {
		clanByTag[clan.Tag] = clan
	}

	playerClans := make([]*models.PlayerClan, len(members))
	for i, member := range members {
		clan, found := clanByTag[member.ClanTag]
		if !found {
			continue
		}
		playerClans[i] = models.NewPlayerClan(member, clan.Name)
	}

	return repo.db.Save(playerClans).Error
}

func (repo *CocLiveRepo) savePlayerStats(cocPlayers []*coc.Player) error {
	playerStats := make([]*models.PlayerStats, len(cocPlayers))
	for i, player := range cocPlayers {
		playerStats[i] = models.NewPlayerStats(player)
	}

	return repo.db.Save(playerStats).Error
}
