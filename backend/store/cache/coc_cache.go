package cache

import (
	"errors"
	"log"
	"sync"
	"time"

	"github.com/amaanq/coc.go"
	"gorm.io/gorm"

	"backend/api/types"
	"backend/client"
	"backend/store/postgres/models"
)

// CocCache stores current data from the COC-API.
type CocCache struct {
	sync.Mutex
	client        *client.CocClient
	db            *gorm.DB
	isMaintenance bool // true if the COC-API is in maintenance mode

	Players            []*types.Player
	PlayerByTag        SyncMap[*types.Player]
	PlayersByClanTag   SyncMap[[]*types.Player]
	PlayersByDiscordID SyncMap[[]*types.Player]
	Clans              []*types.Clan
	ClanByTag          SyncMap[*types.Clan]
}

const MemberOrder = "clan_role like 'member', clan_role like 'admin', clan_role like 'coLeader', clan_role like 'leader'"

func NewCocCache(db *gorm.DB, refreshInterval time.Duration) (*CocCache, error) {
	start := time.Now()
	log.Print("Initializing Clash of Clans cache...")

	cocClient, err := client.NewCocClient()
	if err != nil {
		return nil, err
	}

	cache := &CocCache{db: db, client: cocClient}
	if err = cache.refresh(); err != nil {
		return nil, err
	}

	go cache.startRefresh(refreshInterval)

	log.Printf("Initialized Clash of Clans cache in %s.", time.Since(start).Round(time.Millisecond).String())
	return cache, nil
}

func (cache *CocCache) IsMaintenance() bool {
	return cache.isMaintenance
}

func (cache *CocCache) startRefresh(interval time.Duration) {
	for range time.Tick(interval) {
		if err := cache.refresh(); err != nil {
			log.Printf("Error while refreshing COC cache: %v", err)
		}
	}
}

func (cache *CocCache) refresh() error {
	var wg sync.WaitGroup
	var errs [2]error

	wg.Add(2)
	var lostClans models.LostClans
	go func() {
		defer wg.Done()
		errs[0] = cache.db.Order("id").Find(&lostClans).Error
	}()

	var lostMembers models.Members
	go func() {
		defer wg.Done()
		errs[1] = cache.db.Preload("DiscordLink").Order(MemberOrder).Find(&lostMembers).Error
	}()
	wg.Wait()

	if err := errors.Join(errs[0], errs[1]); err != nil {
		return err
	}

	var clans []*coc.Clan
	var players []*coc.Player

	wg.Add(2)
	start := time.Now()
	go func() {
		defer wg.Done()
		clans, errs[0] = cache.fetchClans(lostClans.Tags())
	}()
	go func() {
		defer wg.Done()
		players, errs[1] = cache.fetchPlayers(lostMembers.TagsDistinct())
	}()
	wg.Wait()

	err := errors.Join(errs[0], errs[1])
	if cache.client.IsMaintenanceErr(err) {
		cache.isMaintenance = true
	} else {
		cache.isMaintenance = false
	}
	if err != nil {
		log.Print("Error while fetching data from COC-API.")
		return err
	}

	log.Printf("Successfully fetched %d clans and %d players from COC-API in %s.", len(clans), len(players), time.Since(start).Round(time.Millisecond).String())

	playersByTag := make(map[string]*coc.Player)
	for _, player := range players {
		playersByTag[player.Tag] = player
	}

	clansByTag := make(map[string]*coc.Clan)
	for _, clan := range clans {
		clansByTag[clan.Tag] = clan
	}

	// set clans first because players depends on clans
	cache.setClans(clans, playersByTag, lostMembers)
	cache.setPlayers(players, lostMembers)

	return nil
}

func (cache *CocCache) fetchClans(tags []string) ([]*coc.Clan, error) {
	if cache.isMaintenance {
		return nil, errors.New("COC-API is in maintenance mode")
	}

	clans, err := cache.client.GetClans(tags)
	if err != nil {
		return nil, err
	}

	return clans, nil
}

// fetchPlayers fetches []*coc.Player from the COC-API and returns a slice of them.
func (cache *CocCache) fetchPlayers(tags []string) ([]*coc.Player, error) {
	players := cache.client.GetPlayers(tags)

	failedPlayerCounter := 0
	var successPlayers []*coc.Player
	for _, player := range players {
		if player == nil {
			failedPlayerCounter++
			continue
		}
		successPlayers = append(successPlayers, player)
	}

	if failedPlayerCounter > 0 {
		log.Printf("Failed to fetch %d players from COC-API.", failedPlayerCounter)
		return nil, errors.New("failed to fetch players from COC-API")
	}

	return successPlayers, nil
}

// setClans sets the clan members of each clan as they are in the store and overwrites them in cache.
func (cache *CocCache) setClans(cocClans []*coc.Clan, playersByTag map[string]*coc.Player, members models.Members) {
	membersByClan := make(map[string]models.Members)
	for _, member := range members {
		membersByClan[member.ClanTag] = append(membersByClan[member.ClanTag], member)
	}

	cache.Lock()
	cache.Clans = make([]*types.Clan, len(cocClans))
	cache.ClanByTag = SyncMap[*types.Clan]{}
	for i, cocClan := range cocClans {
		clanMembers := membersByClan[cocClan.Tag]
		memberList := make([]types.ClanMember, 0)
		if clanMembers != nil && len(clanMembers) > 0 {
			for _, member := range clanMembers {
				if player, found := playersByTag[member.PlayerTag]; found {
					memberList = append(memberList, types.NewClanMember(player, member.ClanRole))
				}
			}
		}

		clan := types.NewClan(cocClan, memberList)
		cache.Clans[i] = clan
		cache.ClanByTag.Set(clan.Tag, clan)
	}

	cache.Unlock()
}

// setPlayers sets the clan and role of each player as they are in the members table.
func (cache *CocCache) setPlayers(cocPlayers []*coc.Player, members models.Members) {
	cache.Lock()
	cache.Players = make([]*types.Player, len(cocPlayers))
	cache.PlayerByTag = SyncMap[*types.Player]{}

	for i, cocPlayer := range cocPlayers {
		player := types.NewPlayer(cocPlayer)
		player.ComparableStatsByName = cache.comparableStatsByName(cocPlayer)
		cache.Players[i] = player
		cache.PlayerByTag.Set(player.Tag, player)
	}

	cache.PlayersByClanTag = SyncMap[[]*types.Player]{}
	for _, member := range members {
		if player, playerFound := cache.PlayerByTag.Get(member.PlayerTag); playerFound {
			player.DiscordID = member.DiscordLink.DiscordID
			if clan, clanFound := cache.ClanByTag.Get(member.ClanTag); clanFound {
				player.Clans = append(player.Clans, types.PlayerClan{
					Name: clan.Name,
					Tag:  clan.Tag,
					Role: member.ClanRole,
				})
				clanPlayers, _ := cache.PlayersByClanTag.Get(member.ClanTag)
				cache.PlayersByClanTag.Set(member.ClanTag, append(clanPlayers, player))
			}
		}
	}

	cache.PlayersByDiscordID = SyncMap[[]*types.Player]{}
	for _, player := range cache.Players {
		current, _ := cache.PlayersByDiscordID.Get(player.DiscordID)
		cache.PlayersByDiscordID.Set(player.DiscordID, append(current, player))
	}

	cache.Unlock()
}
func (cache *CocCache) comparableStatsByName(player *coc.Player) map[string]int {
	res := make(map[string]int, len(player.Achievements))

	res[types.StatSeasonWins.Name] = player.AttackWins
	for _, achievement := range player.Achievements {
		res[achievement.Name] = achievement.Value
	}

	return res
}
