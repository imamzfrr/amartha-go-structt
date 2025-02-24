package repository

import (
	"github.com/imamzfrr/amartha-go-struct/entity"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Create(customer *entity.Customer) error
	GetAll() ([]entity.Customer, error)
	GetByID(id string) (*entity.Customer, error)
	Update(id string, product *entity.Customer) error
	Delete(id string) error
}

type customerRepositoryImpl struct {
	db *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepositoryImpl{db}
}

func (r *customerRepositoryImpl) Create(customer *entity.Customer) error {
	return r.db.Create(customer).Error
}

func (r *customerRepositoryImpl) GetAll() ([]entity.Customer, error) {
	var customer []entity.Customer
	err := r.db.Find(&customer).Error
	return customer, err
}

func (r *customerRepositoryImpl) GetByID(id string) (*entity.Customer, error) {
	var product entity.Customer
	err := r.db.First(&product, "product_id = ?", id).Error
	return &product, err
}

func (r *customerRepositoryImpl) Update(id string, product *entity.Customer) error {
	return r.db.Model(&entity.Customer{}).Where("product_id = ?", id).Updates(product).Error
}

func (r *customerRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&entity.Customer{}, "product_id = ?", id).Error
}
