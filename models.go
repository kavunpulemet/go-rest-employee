package models

import "errors"

type Employee struct {
	Id           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"name"`
	Surname      string `json:"surname" db:"surname"`
	Phone        string `json:"phone" db:"phone"`
	CompanyId    int    `json:"company_id" db:"company_id"`
	PassportId   int    `json:"passport_id" db:"passport_id"`
	DepartmentId int    `json:"department_id" db:"department_id"`
}

type Passport struct {
	Id     int    `json:"passport_id" db:"passport_id"`
	Type   string `json:"passport_type" db:"passport_type"`
	Number string `json:"passport_number" db:"passport_number"`
}

type Department struct {
	Id    int    `json:"department_id" db:"department_id"`
	Name  string `json:"department_name" db:"department_name"`
	Phone string `json:"department_phone" db:"department_phone"`
}

type CreateEmployeeRequest struct {
	Name       string     `json:"name"`
	Surname    string     `json:"surname"`
	Phone      string     `json:"phone"`
	CompanyId  int        `json:"company_id"`
	Passport   Passport   `json:"passport"`
	Department Department `json:"department"`
}

type EmployeeResponse struct {
	Id         int    `json:"id" db:"id"`
	Name       string `json:"name" db:"name"`
	Surname    string `json:"surname" db:"surname"`
	Phone      string `json:"phone" db:"phone"`
	CompanyId  int    `json:"company_id" db:"company_id"`
	Passport   `json:"passport"`
	Department `json:"department"`
}

type UpdateEmployeeInput struct {
	Name         *string     `json:"name"`
	Surname      *string     `json:"surname"`
	Phone        *string     `json:"phone"`
	CompanyId    *int        `json:"company_id"`
	DepartmentId *int        `json:"department_id"`
	Passport     *Passport   `json:"passport"`
	Department   *Department `json:"department"`
}

func (i UpdateEmployeeInput) Validate() error {
	if i.Name == nil && i.Surname == nil && i.Phone == nil && i.CompanyId == nil && i.DepartmentId == nil && i.Passport == nil && i.Department == nil {
		return errors.New("update structure has no values")
	}

	return nil
}
