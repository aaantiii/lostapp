package repos

import (
	"errors"

	"gorm.io/gorm"

	"github.com/aaantiii/lostapp/backend/store/postgres/models"
)

type IUsersRepo interface {
	UserByDiscordID(discordID string) (*models.User, error)
	CreateOrUpdateUser(user *models.User) error
	UserIsAdmin(discordID string) (bool, error)
}

type UsersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) IUsersRepo {
	return &UsersRepo{db: db}
}

func (r *UsersRepo) UserByDiscordID(discordID string) (*models.User, error) {
	var user *models.User
	err := r.db.First(&user, "discord_id = ?", discordID).Error
	return user, err
}

func (r *UsersRepo) CreateOrUpdateUser(user *models.User) error {
	user.IsAdmin = false
	existingUser, err := r.UserByDiscordID(user.DiscordID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err == nil {
		user.IsAdmin = existingUser.IsAdmin
	}
	if user == existingUser {
		return nil
	}

	return r.db.Save(user).Error
}

func (r *UsersRepo) UserIsAdmin(discordID string) (bool, error) {
	user, err := r.UserByDiscordID(discordID)
	if err != nil {
		return false, err
	}

	return user.IsAdmin, nil
}
