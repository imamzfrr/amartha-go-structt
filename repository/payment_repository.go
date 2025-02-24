package repository

import (
	"github.com/imamzfrr/amartha-go-struct/entity"
	"gorm.io/gorm"
)

type PaymentRepository interface {
	Create(payment *entity.Payment) error
	GetAll() ([]entity.Payment, error)
	GetById(id string) (*entity.Payment, error)
	Update(id string, payment *entity.Payment) error
	Delete(id string) error
}

type paymentRepositoryImpl struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepositoryImpl{db}
}

func (r *paymentRepositoryImpl) Create(payment *entity.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepositoryImpl) GetAll() ([]entity.Payment, error) {
	var payments []entity.Payment
	err := r.db.Find(&payments).Error
	return payments, err
}

func (r *paymentRepositoryImpl) GetById(id string) (*entity.Payment, error) {
	var payment entity.Payment
	err := r.db.First(&payment, "payment_id = ?", id).Error
	return &payment, err
}

func (r *paymentRepositoryImpl) Update(id string, payment *entity.Payment) error {
	return r.db.Model(&entity.Payment{}).Where("payment_id = ?", id).Updates(payment).Error
}

func (r *paymentRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&entity.Payment{}, "payment_id = ?", id).Error
}
