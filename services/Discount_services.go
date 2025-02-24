package services

import (
	"github.com/imamzfrr/amartha-go-struct/entity"
	"github.com/imamzfrr/amartha-go-struct/repository"
)

type DiscountService interface {
	Create(discount *entity.Discount) error
	GetAll() ([]entity.Discount, error)
	GetById(id string) (*entity.Discount, error)
	Update(id string, discount *entity.Discount) error
	Delete(id string) error
}

type discountServiceImpl struct {
	discountRepo repository.DiscountRepository
}

func NewDiscountService(repo repository.DiscountRepository) DiscountService {
	return &discountServiceImpl{repo}
}

func (s *discountServiceImpl) Create(discount *entity.Discount) error {
	return s.discountRepo.Create(discount)
}

func (s *discountServiceImpl) GetAll() ([]entity.Discount, error) {
	return s.discountRepo.GetAll()
}

func (s *discountServiceImpl) GetById(id string) (*entity.Discount, error) {
	return s.discountRepo.GetById(id)
}

func (s *discountServiceImpl) Update(id string, discount *entity.Discount) error {
	return s.discountRepo.Update(id, discount)
}

func (s *discountServiceImpl) Delete(id string) error {
	return s.discountRepo.Delete(id)
}
