package repository

import (
	"github.com/imamzfrr/amartha-go-struct/entity"
	"gorm.io/gorm"
)

type SupplierRepository interface {
	Create(supplier *entity.SupplierEntity) error
	GetAll() ([]entity.SupplierEntity, error)
	GetById(id uint64) (*entity.SupplierEntity, error)
	Update(id uint64, supplier *entity.SupplierEntity) error
	Delete(id uint64) error
}

type supplierRepositoryImpl struct {
	db *gorm.DB
}

func NewSupplierRepository(db *gorm.DB) SupplierRepository {
	return &supplierRepositoryImpl{db}
}

func (r *supplierRepositoryImpl) Create(supplier *entity.SupplierEntity) error {
	return r.db.Create(supplier).Error
}

func (r *supplierRepositoryImpl) GetAll() ([]entity.SupplierEntity, error) {
	var suppliers []entity.SupplierEntity
	err := r.db.Find(&suppliers).Error
	return suppliers, err
}

func (r *supplierRepositoryImpl) GetById(id uint64) (*entity.SupplierEntity, error) {
	var supplier entity.SupplierEntity
	err := r.db.First(&supplier, "id = ?", id).Error
	return &supplier, err
}

func (r *supplierRepositoryImpl) Update(id uint64, supplier *entity.SupplierEntity) error {
	return r.db.Model(&entity.SupplierEntity{}).Where("id = ?", id).Updates(supplier).Error
}

func (r *supplierRepositoryImpl) Delete(id uint64) error {
	return r.db.Delete(&entity.SupplierEntity{}, "id = ?", id).Error
}
