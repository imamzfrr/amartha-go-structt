package services

import (
	"github.com/imamzfrr/amartha-go-struct/entity"
	"github.com/imamzfrr/amartha-go-struct/repository"
)

type CategoryService interface {
	CreateCategory(category *entity.CategoryEntity) error
	GetAllCategories() ([]entity.CategoryEntity, error)
	GetCategoryByID(id uint64) (*entity.CategoryEntity, error)
	UpdateCategory(id uint64, category *entity.CategoryEntity) error
	DeleteCategory(id uint64) error
}

type categoryServiceImpl struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) CategoryService {
	return &categoryServiceImpl{repo}
}

func (s *categoryServiceImpl) CreateCategory(category *entity.CategoryEntity) error {
	return s.repo.Create(category)
}

func (s *categoryServiceImpl) GetAllCategories() ([]entity.CategoryEntity, error) {
	return s.repo.FindAll()
}

func (s *categoryServiceImpl) GetCategoryByID(id uint64) (*entity.CategoryEntity, error) {
	return s.repo.FindById(id)
}

func (s *categoryServiceImpl) UpdateCategory(id uint64, category *entity.CategoryEntity) error {
	return s.repo.Update(id, category)
}

func (s *categoryServiceImpl) DeleteCategory(id uint64) error {
	return s.repo.Delete(id)
}
