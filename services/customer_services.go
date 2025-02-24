package services

import (
	"github.com/imamzfrr/amartha-go-struct/entity"
	"github.com/imamzfrr/amartha-go-struct/repository"
)

type CustomerService interface {
	Create(customer *entity.Customer) error
	GetAll() ([]entity.Customer, error)
	GetCustomerById(id string) (*entity.Customer, error)
	Update(id string, customer *entity.Customer) error
	Delete(id string) error
}

type customerServiceImpl struct {
	Repo repository.CustomerRepository
}

func NewCustomerService(repo repository.CustomerRepository) CustomerService {
	return &customerServiceImpl{repo}
}
func (s *customerServiceImpl) Create(customer *entity.Customer) error {
	return s.Repo.Create(customer)
}

func (s *customerServiceImpl) GetAll() ([]entity.Customer, error) {
	return s.Repo.GetAll()
}

func (s *customerServiceImpl) GetCustomerById(id string) (*entity.Customer, error) {
	return s.Repo.GetByID(id)
}

func (s *customerServiceImpl) Update(id string, customer *entity.Customer) error {
	return s.Repo.Update(id, customer)
}

func (s *customerServiceImpl) Delete(id string) error {
	return s.Repo.Delete(id)
}
