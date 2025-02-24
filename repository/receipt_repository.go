package repository

import (
	"github.com/imamzfrr/amartha-go-struct/entity"
	"gorm.io/gorm"
)

type ReceiptRepository interface {
	Create(receipt *entity.Receipt) error
	GetAll() ([]entity.Receipt, error)
	GetById(id string) (*entity.Receipt, error)
	Update(id string, receipt *entity.Receipt) error
	Delete(id string) error
}

type receiptRepositoryImpl struct {
	db *gorm.DB
}

func NewReceiptRepository(db *gorm.DB) ReceiptRepository {
	return &receiptRepositoryImpl{db}
}

func (r *receiptRepositoryImpl) Create(receipt *entity.Receipt) error {
	return r.db.Create(receipt).Error
}

func (r *receiptRepositoryImpl) GetAll() ([]entity.Receipt, error) {
	var receipts []entity.Receipt
	err := r.db.Find(&receipts).Error
	return receipts, err
}

func (r *receiptRepositoryImpl) GetById(id string) (*entity.Receipt, error) {
	var receipt entity.Receipt
	err := r.db.First(&receipt, "receipt_id = ?", id).Error
	return &receipt, err
}

func (r *receiptRepositoryImpl) Update(id string, receipt *entity.Receipt) error {
	return r.db.Model(&entity.Receipt{}).Where("receipt_id = ?", id).Updates(receipt).Error
}

func (r *receiptRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&entity.Receipt{}, "receipt_id = ?", id).Error
}
