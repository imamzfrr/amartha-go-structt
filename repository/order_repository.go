package repository

import (
	"github.com/imamzfrr/amartha-go-struct/entity"
	"gorm.io/gorm"
)

type OrderRepository interface {
	Create(order *entity.Order) error
	GetAll() ([]entity.Order, error)
	GetById(id string) (*entity.Order, error)
	Update(id string, order *entity.Order) error
	Delete(id string) error
}

type orderRepositoryImpl struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepositoryImpl{db}
}

func (r *orderRepositoryImpl) Create(order *entity.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepositoryImpl) GetAll() ([]entity.Order, error) {
	var orders []entity.Order
	err := r.db.Preload("OrderItems").Find(&orders).Error
	return orders, err
}

func (r *orderRepositoryImpl) GetById(id string) (*entity.Order, error) {
	var order entity.Order
	err := r.db.Preload("OrderItems").First(&order, "order_id = ?", id).Error
	return &order, err
}

func (r *orderRepositoryImpl) Update(id string, order *entity.Order) error {
	return r.db.Model(&entity.Order{}).Where("order_id = ?", id).Updates(order).Error
}

func (r *orderRepositoryImpl) Delete(id string) error {
	return r.db.Delete(&entity.Order{}, "order_id = ?", id).Error
}
