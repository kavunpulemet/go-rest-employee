package mappers

import (
	"github.com/magiconair/properties/assert"
	"go-rest-employee/models"
	repository "go-rest-employee/pkg/repository/models"
	"testing"
)

func TestMapToCreateEmployee(t *testing.T) {
	input := models.Employee{
		Id:        1,
		Name:      "Sergey",
		Surname:   "Dayneko",
		Phone:     "123456789",
		CompanyId: 123,
		Passport: models.Passport{
			Id:     321,
			Type:   "type",
			Number: "789",
		},
		Department: models.Department{
			Id:    456,
			Name:  "Department",
			Phone: "987654321",
		},
	}

	expected := repository.CreateEmployee{
		Name:      "Sergey",
		Surname:   "Dayneko",
		Phone:     "123456789",
		CompanyId: 123,
		Passport: repository.Passport{
			Id:     321,
			Type:   "type",
			Number: "789",
		},
		Department: repository.Department{
			Id:    456,
			Name:  "Department",
			Phone: "987654321",
		},
	}

	result := MapToCreateEmployee(input)

	assert.Equal(t, expected, result)
}

func TestMapToUpdateEmployee(t *testing.T) {
	input := models.Employee{
		Id:        1,
		Name:      "Sergey",
		Surname:   "Dayneko",
		Phone:     "123456789",
		CompanyId: 123,
		Passport: models.Passport{
			Id:     321,
			Type:   "type",
			Number: "789",
		},
		Department: models.Department{
			Id:    456,
			Name:  "Department",
			Phone: "987654321",
		},
	}

	expected := repository.UpdateEmployee{
		Name:         "Sergey",
		Surname:      "Dayneko",
		Phone:        "123456789",
		CompanyId:    123,
		DepartmentId: 456,
		Passport: repository.Passport{
			Type:   "type",
			Number: "789",
		},
		Department: repository.Department{
			Name:  "Department",
			Phone: "987654321",
		},
	}

	result := MapToUpdateEmployee(input)

	assert.Equal(t, expected, result)
}

func TestMapFromEmployeeResponse(t *testing.T) {
	input := []repository.EmployeeResponse{
		{
			Id:        1,
			Name:      "Sergey",
			Surname:   "Dayneko",
			Phone:     "123456789",
			CompanyId: 123,
			Passport: repository.Passport{
				Id:     321,
				Type:   "type",
				Number: "789",
			},
			Department: repository.Department{
				Id:    456,
				Name:  "Department",
				Phone: "987654321",
			},
		},
	}

	expected := []models.Employee{
		{
			Id:        1,
			Name:      "Sergey",
			Surname:   "Dayneko",
			Phone:     "123456789",
			CompanyId: 123,
			Passport: models.Passport{
				Id:     321,
				Type:   "type",
				Number: "789",
			},
			Department: models.Department{
				Id:    456,
				Name:  "Department",
				Phone: "987654321",
			},
		},
	}

	result := MapFromEmployeeResponse(input)

	assert.Equal(t, expected, result)
}
