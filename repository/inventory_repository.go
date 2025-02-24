package repository

import (
	"github.com/imamzfrr/amartha-go-struct/entity"
	"gorm.io/gorm"
)

type InventoryRepository interface {
	Create(inventory *entity.Inventory) error
	GetAll() ([]entity.Inventory, error)
	GetById(id string) (*entity.Inventory, error)
	Update(id string, inventory *entity.Inventory) error
	Delete(id string) error
}

type inventoryRepositoryImpl struct {
	db *gorm.DB
}

func NewInventoryRepository(db *gorm.DB) InventoryRepository {
	return &inventoryRepositoryImpl{db}
}

func (r *inventoryRepositoryImpl) Create(inventory *entity.Inventory) error {
	return r.db.Create(inventory).Error
}

func (r *inventoryRepositoryImpl) GetAll() ([]entity.Inventory, error) {
	var inventories []entity.Inventory
	err := r.db.Find(&inventories).Error
	return inventories, err
}

func (r *inventoryRepositoryImpl) GetById(id string) (*entity.Inventory, error) {
	var inventory entity.Inventory
	err := r.db.First(&inventory, "product_id = ?", id).Error
	return &inventory, err
}

func (r *inventoryRepositoryImpl) Update(id string, inventory *entity.Inventory) error {
	return r.db.Model(&entity.Inventory{}).Where("product_id = ?", id).Updates(inventory).Error
}

func (r *inventoryRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&entity.Inventory{}, "product_id = ?", id).Error
}
