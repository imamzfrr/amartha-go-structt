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
	_, err := s.repo.Create(*category)
	return err
}

func (s *categoryServiceImpl) GetAllCategories() ([]entity.CategoryEntity, error) {
	return s.repo.FindAll()
}

func (s *categoryServiceImpl) GetCategoryByID(id uint64) (*entity.CategoryEntity, error) {
	category, err := s.repo.FindById(id)
	return &category, err
}

func (s *categoryServiceImpl) UpdateCategory(id uint64, category *entity.CategoryEntity) error {
	_, err := s.repo.Update(*category)
	return err
}

func (s *categoryServiceImpl) DeleteCategory(id uint64) error {
	_, err := s.repo.DeleteById(id)
	return err
}
