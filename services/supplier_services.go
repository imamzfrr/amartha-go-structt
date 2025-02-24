package services

import (
	"github.com/imamzfrr/amartha-go-struct/entity"
	"github.com/imamzfrr/amartha-go-struct/repository"
)

type SupplierService interface {
	Create(supplier *entity.SupplierEntity) error
	GetAll() ([]entity.SupplierEntity, error)
	GetById(id uint64) (*entity.SupplierEntity, error)
	Update(id uint64, supplier *entity.SupplierEntity) error
	Delete(id uint64) error
}

type supplierServiceImpl struct {
	supplierRepo repository.SupplierRepository
}

func NewSupplierService(repo repository.SupplierRepository) SupplierService {
	return &supplierServiceImpl{repo}
}

func (s *supplierServiceImpl) Create(supplier *entity.SupplierEntity) error {
	return s.supplierRepo.Create(supplier)
}

func (s *supplierServiceImpl) GetAll() ([]entity.SupplierEntity, error) {
	return s.supplierRepo.GetAll()
}

func (s *supplierServiceImpl) GetById(id uint64) (*entity.SupplierEntity, error) {
	return s.supplierRepo.GetById(id)
}

func (s *supplierServiceImpl) Update(id uint64, supplier *entity.SupplierEntity) error {
	return s.supplierRepo.Update(id, supplier)
}

func (s *supplierServiceImpl) Delete(id uint64) error {
	return s.supplierRepo.Delete(id)
}
