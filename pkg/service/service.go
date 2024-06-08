package service

import (
	"go-rest-employee/models"
	"go-rest-employee/pkg/repository"
	"go-rest-employee/pkg/service/mappers"
)

type EmployeeService interface {
	Create(employee models.Employee) (int, error)
	GetByCompany(companyId int) ([]models.Employee, error)
	GetByDepartment(departmentName string) ([]models.Employee, error)
	Update(employeeId int, input models.Employee) error
	Delete(employeeId int) error
}

type ImplEmployee struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) *ImplEmployee {
	return &ImplEmployee{repo: repo}
}

func (s *ImplEmployee) Create(employee models.Employee) (int, error) {
	return s.repo.Create(mappers.MapToCreateEmployee(employee))
}

func (s *ImplEmployee) GetByCompany(companyId int) ([]models.Employee, error) {
	employees, err := s.repo.GetByCompany(companyId)

	return mappers.MapFromEmployeeResponse(employees), err
}

func (s *ImplEmployee) GetByDepartment(departmentName string) ([]models.Employee, error) {
	employees, err := s.repo.GetByDepartment(departmentName)

	return mappers.MapFromEmployeeResponse(employees), err
}

func (s *ImplEmployee) Update(employeeId int, input models.Employee) error {
	update := mappers.MapToUpdateEmployee(input)
	//if err := update.Validate(); err != nil {
	//	return err
	//}

	return s.repo.Update(employeeId, update)
}

func (s *ImplEmployee) Delete(employeeId int) error {
	return s.repo.Delete(employeeId)
}
