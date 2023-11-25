package repos

import (
	"errors"

	"gorm.io/gorm"

	"bot/store/postgres/models"
)

type IUsersRepo interface {
	User(discordID string) (*models.User, error)
	CreateOrUpdateUser(user *models.User) error
	UserIsAdmin(discordID string) (bool, error)
}

type UsersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) IUsersRepo {
	return &UsersRepo{db: db}
}

func (repo *UsersRepo) User(discordID string) (*models.User, error) {
	var user *models.User
	err := repo.db.First(&user, "discord_id = ?", discordID).Error
	return user, err
}

func (repo *UsersRepo) CreateOrUpdateUser(user *models.User) error {
	user.IsAdmin = false
	existingUser, err := repo.User(user.DiscordID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	if err == nil {
		user.IsAdmin = existingUser.IsAdmin
	}
	if user == existingUser {
		return nil
	}

	return repo.db.Save(user).Error
}

func (repo *UsersRepo) UserIsAdmin(discordID string) (bool, error) {
	user, err := repo.User(discordID)
	if err != nil {
		return false, err
	}

	return user.IsAdmin, nil
}
