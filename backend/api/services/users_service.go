package services

import (
	"github.com/aaantiii/lostapp/backend/api/repos"
	"github.com/aaantiii/lostapp/backend/store/postgres/models"
)

type IUsersService interface {
	UserByID(id string) (*models.User, error)
}

type UsersService struct {
	users repos.IUsersRepo
}

func NewUsersService(users repos.IUsersRepo) IUsersService {
	return &UsersService{users: users}
}

func (s *UsersService) UserByID(id string) (*models.User, error) {
	return s.users.UserByDiscordID(id)
}
