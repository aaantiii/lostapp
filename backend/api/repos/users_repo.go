package repos

import (
	"gorm.io/gorm"

	"backend/store/postgres/models"
)

type IUsersRepo interface {
	User(discordID string) (*models.User, error)
	CreateOrUpdateUser(user *models.User) error
	UserIsAdmin(discordID string) (bool, error)
}

type UsersRepo struct {
	db *gorm.DB
}

func NewUsersRepo(db *gorm.DB) *UsersRepo {
	return &UsersRepo{db: db}
}

func (repo *UsersRepo) User(discordID string) (*models.User, error) {
	var user *models.User
	err := repo.db.First(&user, discordID).Error
	return user, err
}

func (repo *UsersRepo) CreateOrUpdateUser(user *models.User) error {
	user.IsAdmin = false
	existingUser, err := repo.User(user.DiscordID)
	if err == nil {
		user.IsAdmin = existingUser.IsAdmin
	}

	return repo.db.Save(user).Error
}

func (repo *UsersRepo) UserIsAdmin(discordID string) (bool, error) {
	var user models.User
	if err := repo.db.First(&user, discordID).Error; err != nil {
		return false, err
	}

	return user.IsAdmin, nil
}
