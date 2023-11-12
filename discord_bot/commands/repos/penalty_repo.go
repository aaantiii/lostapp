package repos

import (
	"gorm.io/gorm"

	"bot/store/postgres/models"
)

type IPenaltiesRepo interface {
	Create(*models.Penalty) error
	Delete(id uint) error
}

type PenaltiesRepo struct {
	db *gorm.DB
}

func NewPenaltiesRepo(db *gorm.DB) IPenaltiesRepo {
	return &PenaltiesRepo{db: db}
}

func (r *PenaltiesRepo) Create(penalty *models.Penalty) error {
	return r.db.Create(penalty).Error
}

func (r *PenaltiesRepo) Delete(id uint) error {
	return r.db.Delete(&models.Penalty{}, id).Error
}
