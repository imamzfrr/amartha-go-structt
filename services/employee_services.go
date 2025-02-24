package services

import (
	"github.com/imamzfrr/amartha-go-struct/entity"
	"github.com/imamzfrr/amartha-go-struct/repository"
)

type EmployeeService interface {
	Create(employee *entity.Employee) error
	GetAll() ([]entity.Employee, error)
	GetById(id string) (*entity.Employee, error)
	Update(id string, employee *entity.Employee) error
	Delete(id string) error
}

type employeeServiceImpl struct {
	EmployeeRepo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) EmployeeService {
	return &employeeServiceImpl{repo}
}

func (s *employeeServiceImpl) Create(Employee *entity.Employee) error {
	return s.EmployeeRepo.Create(Employee)
}

func (s *employeeServiceImpl) GetAll() ([]entity.Employee, error) {
	return s.EmployeeRepo.GetAll()
}

func (s *employeeServiceImpl) GetById(id string) (*entity.Employee, error) {
	return s.EmployeeRepo.GetById(id)
}

func (s *employeeServiceImpl) Update(id string, Employee *entity.Employee) error {
	return s.EmployeeRepo.Update(id, Employee)
}

func (s *employeeServiceImpl) Delete(id string) error {
	return s.EmployeeRepo.Delete(id)
}
