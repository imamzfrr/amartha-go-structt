package repository

import (
"github.com/imamzfrr/amartha-go-struct/entity"
"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(category *entity.CategoryEntity) error
	FindAll() ([]entity.CategoryEntity, error)
	FindById(id uint64) (*entity.CategoryEntity, error)
	Update(id uint64, category *entity.CategoryEntity) error
	Delete(id uint64) error
}

type categoryRepositoryImpl struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepositoryImpl{db}
}

func (r *categoryRepositoryImpl) Create(category *entity.CategoryEntity) error {
	return r.db.Create(category).Error
}

func (r *categoryRepositoryImpl) FindAll() ([]entity.CategoryEntity, error) {
	var categories []entity.CategoryEntity
	err := r.db.Find(&categories).Error
	return categories, err
}

func (r *categoryRepositoryImpl) FindById(id uint64) (*entity.CategoryEntity, error) {
	var category entity.CategoryEntity
	err := r.db.First(&category, "id = ?", id).Error
	return &category, err
}

func (r *categoryRepositoryImpl) Update(id uint64, category *entity.CategoryEntity) error {
	return r.db.Model(&entity.CategoryEntity{}).Where("id = ?", id).Updates(category).Error
}

func (r *categoryRepositoryImpl) Delete(id uint64) error {
	return r.db.Delete(&entity.CategoryEntity{}, "id = ?", id).Error
}

