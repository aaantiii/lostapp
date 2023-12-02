package validation

import (
	"bot/store/postgres/models"
)

const (
	TagMinLength = 4
	TagMaxLength = 12
)

func ValidateClanRole(role models.ClanRole) bool {
	return role == models.RoleMember || role == models.RoleElder || role == models.RoleCoLeader || role == models.RoleLeader
}
