package services

import (
	"github.com/imamzfrr/amartha-go-struct/entity"
	"github.com/imamzfrr/amartha-go-struct/repository"
)

type ShipService interface {
	Create(ship *entity.ShipEntity) error
	GetAll() ([]entity.ShipEntity, error)
	GetById(id uint64) (*entity.ShipEntity, error)
	Update(id uint64, ship *entity.ShipEntity) error
	Delete(id uint64) error
}

type shipServiceImpl struct {
	shipRepo repository.ShipRepository
}

func NewShipService(repo repository.ShipRepository) ShipService {
	return &shipServiceImpl{repo}
}

func (s *shipServiceImpl) Create(ship *entity.ShipEntity) error {
	return s.shipRepo.Create(ship)
}

func (s *shipServiceImpl) GetAll() ([]entity.ShipEntity, error) {
	return s.shipRepo.GetAll()
}

func (s *shipServiceImpl) GetById(id uint64) (*entity.ShipEntity, error) {
	return s.shipRepo.GetById(id)
}

func (s *shipServiceImpl) Update(id uint64, ship *entity.ShipEntity) error {
	return s.shipRepo.Update(id, ship)
}

func (s *shipServiceImpl) Delete(id uint64) error {
	return s.shipRepo.Delete(id)
}
