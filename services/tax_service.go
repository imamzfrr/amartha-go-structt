package services

import (
	"github.com/imamzfrr/amartha-go-struct/entity"
	"github.com/imamzfrr/amartha-go-struct/repository"
)

type TaxService interface {
	Create(tax *entity.Tax) error
	GetAll() ([]entity.Tax, error)
	GetById(id string) (*entity.Tax, error)
	Update(id string, tax *entity.Tax) error
	Delete(id string) error
}

type taxServiceImpl struct {
	taxRepo repository.TaxRepository
}

func NewTaxService(repo repository.TaxRepository) TaxService {
	return &taxServiceImpl{repo}
}

func (s *taxServiceImpl) Create(tax *entity.Tax) error {
	return s.taxRepo.Create(tax)
}

func (s *taxServiceImpl) GetAll() ([]entity.Tax, error) {
	return s.taxRepo.GetAll()
}

func (s *taxServiceImpl) GetById(id string) (*entity.Tax, error) {
	return s.taxRepo.GetById(id)
}

func (s *taxServiceImpl) Update(id string, tax *entity.Tax) error {
	return s.taxRepo.Update(id, tax)
}

func (s *taxServiceImpl) Delete(id string) error {
	return s.taxRepo.Delete(id)
}
