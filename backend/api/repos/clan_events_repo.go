package repos

import (
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/aaantiii/lostapp/backend/api/types"
	"github.com/aaantiii/lostapp/backend/api/utils"
	"github.com/aaantiii/lostapp/backend/store/postgres"
	"github.com/aaantiii/lostapp/backend/store/postgres/models"
)

type IClanEventsRepo interface {
	// ClanEvents returns paginated clan events.
	ClanEvents(tag string, params types.PaginationParams) (*types.PaginatedResponse[*models.ClanEvent], error)
	// ClanEventByID returns the clan event with the given ID.
	ClanEventByID(id uint) (*models.ClanEvent, error)
	// CurrentClanEvent returns the currently active clan event for the given clan, or nil if there is none.
	CurrentClanEvent(clanTag string) (*models.ClanEvent, error)
	// AllActiveClanEvents returns all currently active clan events.
	AllActiveClanEvents() ([]*models.ClanEvent, error)
	// Count returns the number of clan events.
	Count(tag string) (int64, error)
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

func (r *ClanEventsRepo) ClanEvents(tag string, params types.PaginationParams) (*types.PaginatedResponse[*models.ClanEvent], error) {
	count, err := r.Count(tag)
	if err != nil {
		return nil, err
	}
	if err = utils.ValidatePagination(params, count); err != nil {
		return nil, err
	}

	var events []*models.ClanEvent
	if err = r.db.
		Scopes(postgres.WithPagination(params)).
		Preload(clause.Associations).
		Order("starts_at desc").
		Find(&events, "clan_tag = ?", tag).Error; err != nil {
		return nil, err
	}
	return types.NewPaginatedResponse(events, params, count), nil
}

func (r *ClanEventsRepo) Count(tag string) (int64, error) {
	var count int64
	err := r.db.
		Model(&models.ClanEvent{ClanTag: tag}).
		Count(&count).Error
	return count, err
}

func (r *ClanEventsRepo) ClanEventByID(id uint) (*models.ClanEvent, error) {
	var event *models.ClanEvent
	err := r.db.
		Preload(clause.Associations).
		First(&event, id).Error
	return event, err
}

func (r *ClanEventsRepo) CurrentClanEvent(clanTag string) (*models.ClanEvent, error) {
	var event *models.ClanEvent
	err := r.db.
		Order("timestamp desc").
		Preload("Clan").
		First(&event, "clan_tag = ? AND winner_player_tag IS NULL", clanTag).Error
	return event, err
}

func (r *ClanEventsRepo) AllActiveClanEvents() ([]*models.ClanEvent, error) {
	var events []*models.ClanEvent
	err := r.db.
		Order("ends_at").
		Preload("Clan").
		Find(&events, "winner_player_tag IS NULL").Error
	return events, err
}

func (r *ClanEventsRepo) ClanEventMembers(eventID uint, timestamp time.Time) ([]*models.ClanEventMember, error) {
	var members []*models.ClanEventMember
	err := r.db.Find(&members, "clan_event_id = ? AND timestamp = ?", eventID, timestamp).Error
	return members, err
}

func (r *ClanEventsRepo) CreateClanEvent(event *models.ClanEvent) (uint, error) {
	if err := r.db.Create(&event).Error; err != nil {
		return 0, err
	}
	return event.ID, nil
}

func (r *ClanEventsRepo) CreateClanEventMembers(members []*models.ClanEventMember) error {
	return r.db.Create(members).Error
}

func (r *ClanEventsRepo) UpdateClanEvent(event *models.ClanEvent) error {
	return r.db.Save(event).Error
}

func (r *ClanEventsRepo) DeleteClanEvent(id uint) error {
	return r.db.Delete(&models.ClanEvent{}, id).Error
}
