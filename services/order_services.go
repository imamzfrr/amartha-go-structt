package services

import (
	"github.com/imamzfrr/amartha-go-struct/entity"
	"github.com/imamzfrr/amartha-go-struct/repository"
)

type OrderService interface {
	Create(order *entity.Order) error
	GetAll() ([]entity.Order, error)
	GetById(id string) (*entity.Order, error)
	Update(id string, order *entity.Order) error
	Delete(id string) error
}

type orderServiceImpl struct {
	orderRepo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) OrderService {
	return &orderServiceImpl{repo}
}

func (s *orderServiceImpl) Create(order *entity.Order) error {
	return s.orderRepo.Create(order)
}

func (s *orderServiceImpl) GetAll() ([]entity.Order, error) {
	return s.orderRepo.GetAll()
}

func (s *orderServiceImpl) GetById(id string) (*entity.Order, error) {
	return s.orderRepo.GetById(id)
}

func (s *orderServiceImpl) Update(id string, order *entity.Order) error {
	return s.orderRepo.Update(id, order)
}

func (s *orderServiceImpl) Delete(id string) error {
	return s.orderRepo.Delete(id)
}
