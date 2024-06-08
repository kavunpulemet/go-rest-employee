package mappers

import (
	"go-rest-employee/models"
	repository "go-rest-employee/pkg/repository/models"
)

func MapToCreateEmployee(input models.Employee) repository.CreateEmployee {
	return repository.CreateEmployee{
		Name:      input.Name,
		Surname:   input.Surname,
		Phone:     input.Phone,
		CompanyId: input.CompanyId,
		Passport: repository.Passport{
			Id:     input.Passport.Id,
			Type:   input.Passport.Type,
			Number: input.Passport.Number,
		},
		Department: repository.Department{
			Id:    input.Department.Id,
			Name:  input.Department.Name,
			Phone: input.Department.Phone,
		},
	}
}

func MapToUpdateEmployee(input models.Employee) repository.UpdateEmployee {
	return repository.UpdateEmployee{
		Name:         input.Name,
		Surname:      input.Surname,
		Phone:        input.Phone,
		CompanyId:    input.CompanyId,
		DepartmentId: input.Department.Id,
		Passport: repository.Passport{
			Type:   input.Passport.Type,
			Number: input.Passport.Number,
		},
		Department: repository.Department{
			Name:  input.Department.Name,
			Phone: input.Department.Phone,
		},
	}
}
