package employee

import (
	"go-rest-employee/models"
	"go-rest-employee/pkg/repository"
	mappers2 "go-rest-employee/pkg/service/employee/mappers"
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
	return s.repo.Create(mappers2.MapToCreateEmployee(employee))
}

func (s *ImplEmployee) GetByCompany(companyId int) ([]models.Employee, error) {
	employees, err := s.repo.GetByCompany(companyId)

	return mappers2.MapFromEmployeeResponse(employees), err
}

func (s *ImplEmployee) GetByDepartment(departmentName string) ([]models.Employee, error) {
	employees, err := s.repo.GetByDepartment(departmentName)

	return mappers2.MapFromEmployeeResponse(employees), err
}

func (s *ImplEmployee) Update(employeeId int, input models.Employee) error {
	update := mappers2.MapToUpdateEmployee(input)

	return s.repo.Update(employeeId, update)
}

func (s *ImplEmployee) Delete(employeeId int) error {
	return s.repo.Delete(employeeId)
}
