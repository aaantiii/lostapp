package repos

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/aaantiii/lostapp/backend/store/postgres/models"
)

type IClanEventsRepo interface {
	// ClanEventByID returns the clan event with the given ID.
	ClanEventByID(id uint) (*models.ClanEvent, error)
	// CurrentClanEvent returns the currently active clan event for the given clan, or nil if there is none.
	CurrentClanEvent(clanTag string) (*models.ClanEvent, error)
	// AllActiveClanEvents returns all currently active clan events.
	AllActiveClanEvents() ([]*models.ClanEvent, error)
	// ClanEventMembers returns all members of the clan event with the given ID and timestamp.
	ClanEventMembers(eventID uint, timestamp time.Time) ([]*models.ClanEventMember, error)
	// CreateClanEvent creates a new clan event.
	CreateClanEvent(event *models.ClanEvent) (uint, error)
	// CreateClanEventMembers creates the given clan event members.
	CreateClanEventMembers(members []*models.ClanEventMember) error
	// UpdateClanEvent updates the given clan event.
	UpdateClanEvent(event *models.ClanEvent) error
	// DeleteClanEvent deletes the clan event with the given ID.
	DeleteClanEvent(id uint) error
}

type ClanEventsRepo struct {
	db *gorm.DB
}

func NewClanEventsRepo(db *gorm.DB) IClanEventsRepo {
	return &ClanEventsRepo{
		db: db,
	}
}

func (repo *ClanEventsRepo) ClanEventByID(id uint) (*models.ClanEvent, error) {
	var event *models.ClanEvent
	err := repo.db.
		Preload(clause.Associations).
		First(&event, id).Error
	return event, err
}

func (repo *ClanEventsRepo) CurrentClanEvent(clanTag string) (*models.ClanEvent, error) {
	var event *models.ClanEvent
	err := repo.db.
		Order("timestamp desc").
		Preload("Clan").
		First(&event, "clan_tag = ? AND winner_player_tag IS NULL", clanTag).Error
	return event, err
}

func (repo *ClanEventsRepo) AllActiveClanEvents() ([]*models.ClanEvent, error) {
	var events []*models.ClanEvent
	err := repo.db.
		Order("ends_at").
		Preload("Clan").
		Find(&events, "winner_player_tag IS NULL").Error
	return events, err
}

func (repo *ClanEventsRepo) ClanEventMembers(eventID uint, timestamp time.Time) ([]*models.ClanEventMember, error) {
	var members []*models.ClanEventMember
	err := repo.db.Find(&members, "clan_event_id = ? AND timestamp = ?", eventID, timestamp).Error
	return members, err
}

func (repo *ClanEventsRepo) CreateClanEvent(event *models.ClanEvent) (uint, error) {
	if err := repo.db.Create(&event).Error; err != nil {
		return 0, err
	}
	return event.ID, nil
}

func (repo *ClanEventsRepo) CreateClanEventMembers(members []*models.ClanEventMember) error {
	return repo.db.Create(members).Error
}

func (repo *ClanEventsRepo) UpdateClanEvent(event *models.ClanEvent) error {
	return repo.db.Save(event).Error
}

func (repo *ClanEventsRepo) DeleteClanEvent(id uint) error {
	return repo.db.Delete(&models.ClanEvent{}, id).Error
}
