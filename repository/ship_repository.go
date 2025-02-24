package repository

import (
	"github.com/imamzfrr/amartha-go-struct/entity"
	"gorm.io/gorm"
)

type ShipRepository interface {
	Create(ship *entity.ShipEntity) error
	GetAll() ([]entity.ShipEntity, error)
	GetById(id uint64) (*entity.ShipEntity, error)
	Update(id uint64, ship *entity.ShipEntity) error
	Delete(id uint64) error
}

type shipRepositoryImpl struct {
	db *gorm.DB
}

func NewShipRepository(db *gorm.DB) ShipRepository {
	return &shipRepositoryImpl{db}
}

func (r *shipRepositoryImpl) Create(ship *entity.ShipEntity) error {
	return r.db.Create(ship).Error
}

func (r *shipRepositoryImpl) GetAll() ([]entity.ShipEntity, error) {
	var ships []entity.ShipEntity
	err := r.db.Find(&ships).Error
	return ships, err
}

func (r *shipRepositoryImpl) GetById(id uint64) (*entity.ShipEntity, error) {
	var ship entity.ShipEntity
	err := r.db.First(&ship, "id = ?", id).Error
	return &ship, err
}

func (r *shipRepositoryImpl) Update(id uint64, ship *entity.ShipEntity) error {
	return r.db.Model(&entity.ShipEntity{}).Where("id = ?", id).Updates(ship).Error
}

func (r *shipRepositoryImpl) Delete(id uint64) error {
	return r.db.Delete(&entity.ShipEntity{}, "id = ?", id).Error
}
