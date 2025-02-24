package services

import (
	"github.com/imamzfrr/amartha-go-struct/entity"
	"github.com/imamzfrr/amartha-go-struct/repository"
)

type InventoryService interface {
	Create(inventory *entity.Inventory) error
	GetAll() ([]entity.Inventory, error)
	GetById(id string) (*entity.Inventory, error)
	Update(id string, inventory *entity.Inventory) error
	Delete(id string) error
}

type inventoryServiceImpl struct {
	InventoryRepo repository.InventoryRepository
}

func NewInventoryService(repo repository.InventoryRepository) InventoryService {
	return &inventoryServiceImpl{repo}
}

func (s *inventoryServiceImpl) Create(inventory *entity.Inventory) error {
	return s.InventoryRepo.Create(inventory)
}

func (s *inventoryServiceImpl) GetAll() ([]entity.Inventory, error) {
	return s.InventoryRepo.GetAll()
}

func (s *inventoryServiceImpl) GetById(id string) (*entity.Inventory, error) {
	return s.InventoryRepo.GetById(id)
}

func (s *inventoryServiceImpl) Update(id string, inventory *entity.Inventory) error {
	return s.InventoryRepo.Update(id, inventory)
}

func (s *inventoryServiceImpl) Delete(id string) error {
	return s.InventoryRepo.Delete(id)
}
