package mappers

import (
	"go-rest-employee/models"
	repository "go-rest-employee/pkg/repository/models"
)

func MapToCreateEmployee(serviceEmployee models.Employee) repository.CreateEmployee {
	return repository.CreateEmployee{
		Name:      serviceEmployee.Name,
		Surname:   serviceEmployee.Surname,
		Phone:     serviceEmployee.Phone,
		CompanyId: serviceEmployee.CompanyId,
		Passport: repository.Passport{
			Id:     serviceEmployee.Passport.Id,
			Type:   serviceEmployee.Passport.Type,
			Number: serviceEmployee.Passport.Number,
		},
		Department: repository.Department{
			Id:    serviceEmployee.Department.Id,
			Name:  serviceEmployee.Department.Name,
			Phone: serviceEmployee.Department.Phone,
		},
	}
}

func MapToUpdateEmployee(serviceEmployee models.Employee) repository.UpdateEmployee {
	return repository.UpdateEmployee{
		Name:         serviceEmployee.Name,
		Surname:      serviceEmployee.Surname,
		Phone:        serviceEmployee.Phone,
		CompanyId:    serviceEmployee.CompanyId,
		DepartmentId: serviceEmployee.Department.Id,
		Passport: repository.Passport{
			Type:   serviceEmployee.Passport.Type,
			Number: serviceEmployee.Passport.Number,
		},
		Department: repository.Department{
			Name:  serviceEmployee.Department.Name,
			Phone: serviceEmployee.Department.Phone,
		},
	}
}
