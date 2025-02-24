package services

import (
	"github.com/imamzfrr/amartha-go-struct/entity"
	"github.com/imamzfrr/amartha-go-struct/repository"
)

type PaymentService interface {
	Create(payment *entity.Payment) error
	GetAll() ([]entity.Payment, error)
	GetById(id string) (*entity.Payment, error)
	Update(id string, payment *entity.Payment) error
	Delete(id string) error
}

type paymentServiceImpl struct {
	PaymentRepo repository.PaymentRepository
}

func NewPaymentService(repo repository.PaymentRepository) PaymentService {
	return &paymentServiceImpl{repo}
}

func (s *paymentServiceImpl) Create(payment *entity.Payment) error {
	return s.PaymentRepo.Create(payment)
}

func (s *paymentServiceImpl) GetAll() ([]entity.Payment, error) {
	return s.PaymentRepo.GetAll()
}

func (s *paymentServiceImpl) GetById(id string) (*entity.Payment, error) {
	return s.PaymentRepo.GetById(id)
}

func (s *paymentServiceImpl) Update(id string, payment *entity.Payment) error {
	return s.PaymentRepo.Update(id, payment)
}

func (s *paymentServiceImpl) Delete(id string) error {
	return s.PaymentRepo.Delete(id)
}
