package services

import (
	"github.com/imamzfrr/amartha-go-struct/entity"
	"github.com/imamzfrr/amartha-go-struct/repository"
)

type ReceiptService interface {
	Create(receipt *entity.Receipt) error
	GetAll() ([]entity.Receipt, error)
	GetById(id string) (*entity.Receipt, error)
	Update(id string, receipt *entity.Receipt) error
	Delete(id string) error
}

type receiptServiceImpl struct {
	receiptRepo repository.ReceiptRepository
}

func NewReceiptService(repo repository.ReceiptRepository) ReceiptService {
	return &receiptServiceImpl{repo}
}

func (s *receiptServiceImpl) Create(receipt *entity.Receipt) error {
	return s.receiptRepo.Create(receipt)
}

func (s *receiptServiceImpl) GetAll() ([]entity.Receipt, error) {
	return s.receiptRepo.GetAll()
}

func (s *receiptServiceImpl) GetById(id string) (*entity.Receipt, error) {
	return s.receiptRepo.GetById(id)
}

func (s *receiptServiceImpl) Update(id string, receipt *entity.Receipt) error {
	return s.receiptRepo.Update(id, receipt)
}

func (s *receiptServiceImpl) Delete(id string) error {
	return s.receiptRepo.Delete(id)
}
