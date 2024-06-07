package service

import (
	models "go-rest-employee"
	"go-rest-employee/pkg/repository"
)

type EmployeeService interface {
	Create(employee models.CreateEmployeeRequest) (int, error)
	GetByCompany(companyId int) ([]models.EmployeeResponse, error)
	GetByDepartment(departmentName string) ([]models.EmployeeResponse, error)
	Update(employeeId int, input models.UpdateEmployeeInput) error
	Delete(employeeId int) error
}

type ImplEmployee struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) *ImplEmployee {
	return &ImplEmployee{repo: repo}
}

func (s *ImplEmployee) Create(employee models.CreateEmployeeRequest) (int, error) {
	return s.repo.Create(employee)
}

func (s *ImplEmployee) GetByCompany(companyId int) ([]models.EmployeeResponse, error) {
	return s.repo.GetByCompany(companyId)
}

func (s *ImplEmployee) GetByDepartment(departmentName string) ([]models.EmployeeResponse, error) {
	return s.repo.GetByDepartment(departmentName)
}

func (s *ImplEmployee) Update(employeeId int, input models.UpdateEmployeeInput) error {
	if err := input.Validate(); err != nil {
		return err
	}

	return s.repo.Update(employeeId, input)
}

func (s *ImplEmployee) Delete(employeeId int) error {
	return s.repo.Delete(employeeId)
}
