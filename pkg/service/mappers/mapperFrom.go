package mappers

import (
	"go-rest-employee/models"
	repository "go-rest-employee/pkg/repository/models"
)

func MapFromEmployeeResponse(input []repository.EmployeeResponse) []models.Employee {
	output := make([]models.Employee, len(input))
	for i, repoEmp := range input {
		output[i] = models.Employee{
			Id:        repoEmp.Id,
			Name:      repoEmp.Name,
			Surname:   repoEmp.Surname,
			Phone:     repoEmp.Phone,
			CompanyId: repoEmp.CompanyId,
			Passport: models.Passport{
				Id:     repoEmp.Passport.Id,
				Type:   repoEmp.Passport.Type,
				Number: repoEmp.Passport.Number,
			},
			Department: models.Department{
				Id:    repoEmp.Department.Id,
				Name:  repoEmp.Department.Name,
				Phone: repoEmp.Department.Phone,
			},
		}
	}

	return output
}
