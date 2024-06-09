package mappers

import (
	"go-rest-employee/models"
	repository "go-rest-employee/pkg/repository/models"
)

func MapFromEmployeeResponse(repositoryEmployees []repository.EmployeeResponse) []models.Employee {
	serviceEmployees := make([]models.Employee, len(repositoryEmployees))
	for i, repositoryEmployee := range repositoryEmployees {
		serviceEmployees[i] = models.Employee{
			Id:        repositoryEmployee.Id,
			Name:      repositoryEmployee.Name,
			Surname:   repositoryEmployee.Surname,
			Phone:     repositoryEmployee.Phone,
			CompanyId: repositoryEmployee.CompanyId,
			Passport: models.Passport{
				Id:     repositoryEmployee.Passport.Id,
				Type:   repositoryEmployee.Passport.Type,
				Number: repositoryEmployee.Passport.Number,
			},
			Department: models.Department{
				Id:    repositoryEmployee.Department.Id,
				Name:  repositoryEmployee.Department.Name,
				Phone: repositoryEmployee.Department.Phone,
			},
		}
	}

	return serviceEmployees
}
